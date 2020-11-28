package handle

import (
	"bgp/operation"
	gobgpapi "github.com/osrg/gobgp/api"
	"google.golang.org/grpc"
	"log"
)

func Init(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	operation.Client = gobgpapi.NewGobgpApiClient(conn)

}