package operation

import (
	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
	"log"
)

func AddExecPrefix(ctx context.Context, prefixSetName, ipPrefix string, MaskMin, MaskMax uint32)  {
	ListDefinedSetClient, err := Client.ListDefinedSet(ctx, &gobgpapi.ListDefinedSetRequest{Name: prefixSetName})
	if err != nil {
		log.Fatalf("addExecPrefix happen a err, err is %s", err)
	}
	DefinedSetResponse, err := ListDefinedSetClient.Recv()
	if err != nil {
		log.Fatalf("addExecPrefix.Recv happen a err, err is %s", err)
	}
	DefinedSets := DefinedSetResponse.GetDefinedSet()
	DefinedSets.Prefixes = append(DefinedSets.Prefixes, &gobgpapi.Prefix{IpPrefix: ipPrefix, MaskLengthMin: MaskMin, MaskLengthMax: MaskMax})
}

