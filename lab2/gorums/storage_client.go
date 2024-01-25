package gorums

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "dat520/info/lab2/gorums/proto"

	"github.com/relab/gorums"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageClient struct {
	conf *pb.Configuration
}

// Creates a new StorageClient with the provided srvAddresses as the configuration.
func NewStorageClient(srvAddresses []string) *StorageClient {
	mgr := pb.NewManager(
		gorums.WithDialTimeout(500*time.Millisecond),
		gorums.WithGrpcDialOptions(
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		),
	)
	// Create a configuration including all nodes.
	allNodesConfig, err := mgr.NewConfiguration(
		&qspec{len(srvAddresses)},
		gorums.WithNodeList(srvAddresses),
	)
	if err != nil {
		log.Fatal("Failed to create configuration:", err)
	}
	return &StorageClient{conf: allNodesConfig}
}

// Writes the provided value to a random server.
func (sc *StorageClient) WriteValue(value string) error {
	allNodes := sc.conf.Nodes()
	node := allNodes[rand.Intn(len(allNodes))]
	_, err := node.Write(context.Background(), &pb.WriteRequest{Value: value})
	if err != nil {
		return err
	}
	return nil
}

// Returns a slice of values stored on all servers.
func (sc *StorageClient) ReadValues() ([]string, error) {
	resp, err := sc.conf.Read(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Failed to read values: %v", err)
		return nil, err
	}
	return resp.GetValues(), nil
}

type qspec struct {
	numServers int
}

// Implements the quorum function for the Read method as required by the pb.QuorumSpec interface.
// Requires that all servers reply with a ReadResponse; does not tolerate any failures.
func (q *qspec) ReadQF(_ *emptypb.Empty, replies map[uint32]*pb.ReadResponse) (*pb.ReadResponse, bool) {
	if len(replies) != q.numServers {
		return nil, false
	}
	values := make([]string, 0, q.numServers)
	for _, r := range replies {
		values = append(values, r.GetValues()...)
	}
	return &pb.ReadResponse{Values: values}, true
}
