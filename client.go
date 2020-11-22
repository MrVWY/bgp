package main

import (
	"bgp/handle"
	"bgp/operation"
	"context"
	gobgpapi "github.com/osrg/gobgp/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	operation.Client = gobgpapi.NewGobgpApiClient(conn)


	handle.CreatePrefixSet(context.Background(), "PREFIX", "text", "10.1.1.0/24", "24", "32")
}
