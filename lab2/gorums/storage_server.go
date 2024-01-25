package gorums

import (
	"log"
	"net"
	"sync"

	pb "dat520/info/lab2/gorums/proto"

	"github.com/relab/gorums"
	"google.golang.org/protobuf/types/known/emptypb"
)

// The storage server should implement the server interface defined in the proto file.
type StorageServer struct {
	sync.RWMutex
	data []string
}

// Creates a new StorageServer.
func NewStorageServer() *StorageServer {
	return &StorageServer{
		data: make([]string, 0),
	}
}

// Start the server listening on the provided address string.
// The function should be non-blocking.
// Returns the full listening address of the server as string.
// Hint: Use goroutine to start the server.
func (s *StorageServer) StartServer(addr string) string {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("error creating listener:", err)
		return ""
	}
	gorumsSrv := gorums.NewServer()
	pb.RegisterStorageServiceServer(gorumsSrv, s)
	go func() {
		if err := gorumsSrv.Serve(lis); err != nil {
			log.Printf("Unable to serve: %v", err)
		}
	}()
	return lis.Addr().String()
}

// Write the provided value to the server's data slice.
func (s *StorageServer) Write(ctx gorums.ServerCtx, request *pb.WriteRequest) (response *emptypb.Empty, err error) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, request.GetValue())
	return &emptypb.Empty{}, nil
}

// Read returns the server's data slice.
func (s *StorageServer) Read(ctx gorums.ServerCtx, request *emptypb.Empty) (response *pb.ReadResponse, err error) {
	s.RLock()
	defer s.RUnlock()
	return &pb.ReadResponse{Values: s.data}, nil
}

// Returns the data slice on this server
func (s *StorageServer) GetData() []string {
	s.RLock()
	defer s.RUnlock()
	return s.data
}

// Sets the data slice to a value
func (s *StorageServer) SetData(data []string) {
	s.Lock()
	defer s.Unlock()
	s.data = data
}
