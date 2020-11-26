package operation

import (
	"context"
	"errors"
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
)

func AddPeers(ctx context.Context, Description, NeighborAddress string, LocalAs, PeerAs, SendCommunity int) (string, error) {
	var err error
	has, err := ListPeers(ctx, NeighborAddress)
	if err != nil {
		return "false",  fmt.Errorf("ListPeers is happend a err, err is %s", err)
	}
	if has != nil {
		return "false", errors.New("The Peer is exist ")
	}
	peer := newAddPeerRequest(Description, NeighborAddress, LocalAs, PeerAs, SendCommunity)
	_, err = Client.AddPeer(ctx, peer)
	if err != nil {
		return "false",  fmt.Errorf("AddPeers is happend a err, err is %s", err)
	}
	return "Successful", nil
}

func DeletePeers(ctx context.Context, NeighborAddress string) (string, error) {
	var err error
	has, err := ListPeers(ctx, NeighborAddress)
	if err != nil {
		return "false", fmt.Errorf("ListPeers is happend a err, err is %s", err)
	}
	if has != nil {
		return "false", errors.New("The Peer is exist ")
	}
	_, err = Client.DeletePeer(ctx, &gobgpapi.DeletePeerRequest{Address: NeighborAddress} )
	if err != nil {
		return "false", fmt.Errorf("DeletePeers is happend a err, err is %s", err)
	}
	return "Successful", nil
}

func ListPeers(ctx context.Context, address string) (*gobgpapi.Peer, error) {
	ListPeerClient, err := Client.ListPeer(ctx, &gobgpapi.ListPeerRequest{ Address: address, EnableAdvertised: false } )  //广播?
	if err != nil {
		return nil, fmt.Errorf("ListPeers is happend a err, err is %s", err)
	}
	response, err := ListPeerClient.Recv()
	if err != nil {
		return nil, fmt.Errorf("ListPeers.Recv is happend a err, err is %s", err)
	}
	return response.GetPeer(), nil
}



//UpdatePeer()
//ResetPeer()
//ShutdownPeer()
//EnablePeer()
//DisablePeer()
//MonitorPeer()