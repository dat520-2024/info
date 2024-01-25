package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "dat520/info/lab2/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	endpoint := flag.String("endpoint", "localhost:12111", "Endpoint on which server runs or to which client connects")
	flag.Parse()

	listener, err := net.Listen("tcp", *endpoint)
	if err != nil {
		log.Fatalf("Failed to listen on %v: %v", *endpoint, err)
	}
	fmt.Printf("Listener started on %v\n", *endpoint)

	server := NewKeyValueServicesServer()
	grpcServer := grpc.NewServer()
	pb.RegisterKeyValueServiceServer(grpcServer, server)
	fmt.Printf("Preparing to serve incoming requests.\n")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type keyValueServicesServer struct {
	// lock for the kv map
	mu sync.Mutex
	kv map[string]string
	// this must be included in implementers of the pb.KeyValueServicesServer interface
	pb.UnimplementedKeyValueServiceServer
}

// NewKeyValueServicesServer returns an initialized KeyValueServicesServer
func NewKeyValueServicesServer() *keyValueServicesServer {
	return &keyValueServicesServer{
		kv: make(map[string]string),
	}
}

// Insert inserts a key-value pair from the request into the server's map, and
// returns a response to the client indicating whether or not the insert was successful.
func (s *keyValueServicesServer) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.kv[req.Key] = req.Value
	log.Printf("Inserted key-value pair: %v\n", req)
	return &pb.InsertResponse{Success: true}, nil
}

// Lookup returns a response to containing the value corresponding to the request's key.
// If the key is not found, the response's value is empty.
func (s *keyValueServicesServer) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.kv[req.Key]
	if ok {
		return &pb.LookupResponse{Value: val}, nil
	}
	return &pb.LookupResponse{Value: ""}, nil
}

// Keys returns a response to containing a slice of all the keys in the server's map.
func (s *keyValueServicesServer) Keys(ctx context.Context, req *pb.KeysRequest) (*pb.KeysResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]string, 0, len(s.kv))
	for k := range s.kv {
		keys = append(keys, k)
	}
	return &pb.KeysResponse{Keys: keys}, nil
}
