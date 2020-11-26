package operation

import (
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
)

func AddExecPrefix(ctx context.Context, prefixSetName, ipPrefix string, MaskMin, MaskMax uint32) (string, error) {
	ListDefinedSetClient, err := Client.ListDefinedSet(ctx, &gobgpapi.ListDefinedSetRequest{Name: prefixSetName})
	if err != nil {
		return "false", fmt.Errorf("addExecPrefix happen a err, err is %s", err)
	}
	DefinedSetResponse, err := ListDefinedSetClient.Recv()
	if err != nil {
		return "false", fmt.Errorf("addExecPrefix.Recv happen a err, err is %s", err)
	}
	DefinedSets := DefinedSetResponse.GetDefinedSet()
	DefinedSets.Prefixes = append(DefinedSets.Prefixes, &gobgpapi.Prefix{IpPrefix: ipPrefix, MaskLengthMin: MaskMin, MaskLengthMax: MaskMax})
	return "Successful", nil
}

