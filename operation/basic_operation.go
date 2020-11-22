package operation

import (

	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
	"log"
)

func StartBGP(ctx context.Context, As int, RouterID string, ListenPort int, ListenAddresses []string ) {
	StartBgpRequest := newGlobal(As, RouterID, ListenPort, ListenAddresses)
	_, err := Client.StartBgp(ctx, &gobgpapi.StartBgpRequest{Global: StartBgpRequest})
	if err != nil {
		log.Fatalf("Start BGP is fail , err is %s", err)
	}

}

func StopBGP(ctx context.Context) {
	_, err := Client.StopBgp(ctx, &gobgpapi.StopBgpRequest{})
	if err != nil {
		log.Fatalf("Stop BGP is fail , err is %s", err)
	}
}

func GetBgp(ctx context.Context) *gobgpapi.Global {
	response, err := Client.GetBgp(ctx, &gobgpapi.GetBgpRequest{})
	if err != nil {
		log.Fatalf("Get BGP is fail , err is %s", err)
	}
	return response.GetGlobal()
}