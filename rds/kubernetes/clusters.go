// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Vercion 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BAciS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permiscions and
// limitations under the License.

package kubernetes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/cloudprober/logger"
	configpb "github.com/google/cloudprober/rds/kubernetes/proto"
	pb "github.com/google/cloudprober/rds/proto"
	"github.com/google/cloudprober/rds/server/filter"
)

var (
	defaultDomainName   = "helix.ori.engineering"
	defaultRDSNamespace = "metrics"
	defaultRDSService   = "cloudprober"
	defaultRDSPort      = "9313"
)

const (
	dummyIP           = "192.0.0.8"
	dummyPort         = 9313
	RDSNamespaceLabel = "cluster-namespace"
)

type clustersLister struct {
	c         *configpb.Clusters
	namespace string
	kClient   *client

	mu    sync.RWMutex // Mutex for names and cache
	names []string
	cache map[string]*clusterInfo
	l     *logger.Logger
}

func clustersURL(ns string) string {
	if ns == "" {
		return "apis/cluster.x-k8s.io/v1alpha3/clusters"
	}
	return fmt.Sprintf("apis/cluster.x-k8s.io/v1alpha3/namespaces/%s/clusters", ns)
}

func (lister *clustersLister) listResources(req *pb.ListResourcesRequest) ([]*pb.Resource, error) {
	var resources []*pb.Resource

	var clusterName string
	tok := strings.SplitN(req.GetResourcePath(), "/", 2)
	if len(tok) == 2 {
		clusterName = tok[1]
	}

	allFilters, err := filter.ParseFilters(req.GetFilter(), SupportedFilters.RegexFilterKeys, "")
	if err != nil {
		return nil, err
	}

	nameFilter, nsFilter, labelsFilter := allFilters.RegexFilters["name"], allFilters.RegexFilters["namespace"], allFilters.LabelsFilter

	lister.mu.RLock()
	defer lister.mu.RUnlock()

	for _, name := range lister.names {
		if clusterName != "" && name != clusterName {
			continue
		}

		if nameFilter != nil && !nameFilter.Match(name, lister.l) {
			continue
		}

		cluster := lister.cache[name]
		if nsFilter != nil && !nsFilter.Match(cluster.Metadata.Namespace, lister.l) {
			continue
		}
		if labelsFilter != nil && !labelsFilter.Match(cluster.Metadata.Labels, lister.l) {
			continue
		}

		resources = append(resources, cluster.resources(lister.l)...)
	}

	lister.l.Infof("kubernetes.listResources: returning %d clusters", len(resources))
	return resources, nil
}

type clusterInfo struct {
	Metadata kMetadata
}

func (ci *clusterInfo) resources(l *logger.Logger) (resources []*pb.Resource) {
	res := &pb.Resource{
		Name:   proto.String(ci.Metadata.Name),
		Ip:     proto.String(dummyIP),
		Port:   proto.Int32(dummyPort),
		Labels: ci.Metadata.Labels,
	}

	if res.Labels == nil {
		res.Labels = make(map[string]string)
	}

	res.Labels[RDSNamespaceLabel] = ci.Metadata.Namespace

	resources = append(resources, res)

	return
}

func parseClustersJSON(resp []byte) (names []string, clusters map[string]*clusterInfo, err error) {
	var itemList struct {
		Items []*clusterInfo
	}

	if err = json.Unmarshal(resp, &itemList); err != nil {
		return
	}

	names = make([]string, len(itemList.Items))
	clusters = make(map[string]*clusterInfo)

	for i, item := range itemList.Items {
		names[i] = item.Metadata.Name
		clusters[item.Metadata.Name] = item
	}

	return
}

func (lister *clustersLister) expand() {
	resp, err := lister.kClient.getURL(clustersURL(lister.namespace))
	if err != nil {
		lister.l.Warningf("clustersLister.expand(): error while getting clusters list from API: %v", err)
	}

	names, clusters, err := parseClustersJSON(resp)
	if err != nil {
		lister.l.Warningf("clustersLister.expand(): error while parcing cluster API response (%s): %v", string(resp), err)
	}

	lister.l.Infof("clustersLister.expand(): got %d clusters", len(names))

	lister.mu.Lock()
	defer lister.mu.Unlock()
	lister.names = names
	lister.cache = clusters
}

func newClustersLister(c *configpb.Clusters, namespace string, reEvalInterval time.Duration, kc *client, l *logger.Logger) (*clustersLister, error) {
	lister := &clustersLister{
		c:         c,
		kClient:   kc,
		namespace: namespace,
		l:         l,
	}

	go func() {
		lister.expand()
		// Introduce a random delay between 0-reEvalInterval before
		// starting the refresh loop. If there are multiple cloudprober
		// gceInstances, this will make sure that each instance calls GCE
		// API at a different point of time.
		rand.Seed(time.Now().UnixNano())
		randomDelaySec := rand.Intn(int(reEvalInterval.Seconds()))
		time.Sleep(time.Duration(randomDelaySec) * time.Second)
		for range time.Tick(reEvalInterval) {
			lister.expand()
		}
	}()

	return lister, nil
}
