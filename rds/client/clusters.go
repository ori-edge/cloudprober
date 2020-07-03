package client

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/cloudprober/logger"
	configpb "github.com/google/cloudprober/rds/client/proto"
	k8s "github.com/google/cloudprober/rds/kubernetes"
	pb "github.com/google/cloudprober/rds/proto"
	spb "github.com/google/cloudprober/rds/proto"
	"github.com/google/cloudprober/targets/endpoint"

	"google.golang.org/protobuf/proto"
)

const (
	clustersResourceType = "clusters"
	hardCodedPrefix      = "cloudprober.metrics."
	hardCodedSuffx       = ".telefonica.helix.ori.engineering:9313"
)

func newRDSClient(req *spb.ListResourcesRequest, serverOpts *configpb.ClientConf_ServerOptions, listResources ListResourcesFunc, l *logger.Logger) (Targets, error) {
	c := &configpb.ClientConf{
		Request:       req,
		ReEvalSec:     proto.Int32(10), // Refresh client every 10s.
		ServerOptions: serverOpts,
	}

	client, err := New(c, listResources, l)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *MultiClient) rdsClustersRequest() *spb.ListResourcesRequest {
	return &spb.ListResourcesRequest{
		Provider:     proto.String("k8s"),
		ResourcePath: proto.String(fmt.Sprintf("%s/%s", client.resourceType, "")),
		Filter:       client.filters,
		IpConfig:     client.ipConfig,
	}
}

func (client *MultiClient) buildRemoteFQDN(cluster endpoint.Endpoint) *string {

	remoteSVC := client.serverOpts.GetMulticlusterOptions().GetService()
	remoteNS := client.serverOpts.GetMulticlusterOptions().GetNamespace()

	clusterName := cluster.Name
	localNS := cluster.Labels[k8s.RDSNamespaceLabel]

	subDomain := client.serverOpts.GetMulticlusterOptions().GetDomain()

	fqdn := strings.Join([]string{remoteSVC, remoteNS, clusterName, localNS, subDomain}, ".")
	url := fmt.Sprintf("%s:%d", fqdn, cluster.Port)
	fmt.Printf("Built remote RDS FQDN: %s", url)

	return proto.String(url)
}

func (client *MultiClient) initClients(clusters []endpoint.Endpoint) error {

	for _, cluster := range clusters {
		childServerOpts := client.serverOpts
		childServerOpts.ServerAddress = client.buildRemoteFQDN(cluster)
		childClient, err := newRDSClient(client.c.GetRequest(), childServerOpts, client.listResources, client.l)
		if err != nil {
			return err
		}
		client.clients[cluster.Name] = childClient
	}

	return nil
}

// ListEndpoints returns the list of GCE endpoints.
func (client *MultiClient) ListEndpoints() []endpoint.Endpoint {
	var ep []endpoint.Endpoint
	for _, client := range client.clients {
		ep = append(ep, client.ListEndpoints()...)
	}
	return ep
}

// Resolve resolves the name into an IP address. Unless explicitly configured
// to use DNS, we use the RDS client to determine the resource IPs.
func (client *MultiClient) Resolve(name string, ipVer int) (net.IP, error) {

	var ip net.IP
	var err error

	fmt.Printf("Dummy Resolve() in Multiclient")

	return ip, err
}

func (client *MultiClient) refreshState(time time.Duration) {
	fmt.Printf("Dummy refreshState() in Multiclient")
}

// MultiClient wrapper
type MultiClient struct {
	mu            sync.RWMutex
	c             *configpb.ClientConf
	serverOpts    *configpb.ClientConf_ServerOptions
	l             *logger.Logger
	clients       map[string]Targets
	resourceType  string
	ipConfig      *spb.IPConfig
	filters       []*spb.Filter
	listResources func(context.Context, *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error)
}

func (client *MultiClient) getClusters() ([]endpoint.Endpoint, error) {

	client.serverOpts.Multicluster = proto.Bool(false)
	cl, err := newRDSClient(client.rdsClustersRequest(), client.serverOpts, client.listResources, client.l)
	if err != nil {
		return nil, err
	}

	return cl.ListEndpoints(), nil
}

// NewMultiClient is a wrapper around multiple RDS client instances
func NewMultiClient(c *configpb.ClientConf, listResources ListResourcesFunc, l *logger.Logger) (Targets, error) {
	mc := &MultiClient{
		c:             c,
		clients:       make(map[string]Targets),
		serverOpts:    c.GetServerOptions(),
		l:             l,
		resourceType:  clustersResourceType,
		listResources: listResources,
	}

	clusters, err := mc.getClusters()
	if err != nil {
		return nil, err
	}

	mc.initClients(clusters)

	return mc, nil
}
