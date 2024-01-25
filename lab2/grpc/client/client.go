package main

import (
	"context"
	"flag"
	"log"

	pb "dat520/info/lab2/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	endpoint := flag.String("endpoint", "localhost:12111", "Endpoint on which server runs")
	cmd := flag.String("cmd", "insert", "Command to run (insert, lookup, keys)")
	key := flag.String("key", "foo", "Key to insert or lookup")
	val := flag.String("val", "bar", "Value to insert")
	flag.Parse()

	conn, err := grpc.Dial(*endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial %v: %v", *endpoint, err)
	}
	defer conn.Close()

	log.Printf("Connected to %v\n", *endpoint)
	client := pb.NewKeyValueServiceClient(conn)
	switch *cmd {
	case "insert":
		resp, err := client.Insert(context.Background(), &pb.InsertRequest{Key: *key, Value: *val})
		if err != nil {
			log.Fatalf("Failed to insert: %v", err)
		}
		log.Printf("Insert response: %v\n", resp)

	case "lookup":
		resp, err := client.Lookup(context.Background(), &pb.LookupRequest{Key: *key})
		if err != nil {
			log.Fatalf("Failed to lookup: %v", err)
		}
		log.Printf("Lookup response: %v\n", resp)

	case "keys":
		resp, err := client.Keys(context.Background(), &pb.KeysRequest{})
		if err != nil {
			log.Fatalf("Failed to get keys: %v", err)
		}
		log.Printf("Keys response: %v\n", resp)

	default:
		log.Fatalf("Unknown command: %v", *cmd)
	}
}
