package operation

import (
	"context"
	gobgpapi "github.com/osrg/gobgp/api"
	"log"
)

func AddPeers() {

}

func DeletePeers() {

}

func ListPeers(ctx context.Context, address string) (*gobgpapi.Peer, error) {
	ListPeerClient, err := Client.ListPeer(ctx, &gobgpapi.ListPeerRequest{Address: address})
	if err != nil {
		log.Fatalf("ListPeers is happend a err, err is %s", err)
	}
	response, err := ListPeerClient.Recv()
	if err != nil {
		log.Fatalf("ListPeers.Client.Recv is happend a err, err is %s", err)
	}
	return response.GetPeer(), nil
}


//UpdatePeer()
//ResetPeer()
//ShutdownPeer()
//EnablePeer()
//DisablePeer()
//MonitorPeer()