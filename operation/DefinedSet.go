package operation

import (
	"errors"
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
	"strconv"
)

//DefinedSet
func newAddDefinedSet(Type string, DFSetName string, prefix []*gobgpapi.Prefix, list []string) *gobgpapi.AddDefinedSetRequest {
	return &gobgpapi.AddDefinedSetRequest{
		DefinedSet: newDefinedSet(Type, DFSetName, prefix, list),
	}
}

func newDelDefinedSet(DefinedSet *gobgpapi.DefinedSet) *gobgpapi.DeleteDefinedSetRequest {
	return &gobgpapi.DeleteDefinedSetRequest{
		DefinedSet: DefinedSet,
	}
}

func newDefinedSet(Type string, DFSetName string, prefix []*gobgpapi.Prefix, List []string) *gobgpapi.DefinedSet {
	var DefinedType gobgpapi.DefinedType

	switch Type {
	case "PREFIX":
		DefinedType = gobgpapi.DefinedType_PREFIX
	case "NEIGHBOR":
		DefinedType = gobgpapi.DefinedType_NEIGHBOR
	case "COMMUNITY":
		DefinedType = gobgpapi.DefinedType_COMMUNITY
	}
	return &gobgpapi.DefinedSet{
		DefinedType: DefinedType,
		Name:        DFSetName,
		List:        List,
		Prefixes:    prefix,
	}
}

func AddDefinedSetPrefixSet(ctx context.Context,Type, PrefixSetName, ipPrefix, MaskMin, MaskMax string) (string, error) {
	var err error
	Defined, err := ListDefinedSets(ctx, PrefixSetName)
	if err != nil {
		return "false", fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	if Defined != nil {
		max, _ := strconv.ParseUint(MaskMax, 10, 32)
		min, _ := strconv.ParseUint(MaskMin, 10, 32)
		Defined.Prefixes = append(Defined.Prefixes, &gobgpapi.Prefix{IpPrefix:ipPrefix, MaskLengthMin: uint32(min), MaskLengthMax: uint32(max)})
		return "Successful", nil
	}
	Prefix := newPrefixSet(ipPrefix, MaskMin, MaskMax)
	res := newAddDefinedSet(Type, PrefixSetName, Prefix, nil)
	_, err = Client.AddDefinedSet(ctx, res)
	if err != nil {
		return "false", fmt.Errorf("AddDefinedSet happen a err, err is %s", err)
	}
	return "Successful", nil
}

func AddDefinedSetCommunitySet(ctx context.Context, CommunitySetName, Type string, list []string) (string, error) {
	var err error
	Defined, err := ListDefinedSets(ctx, CommunitySetName)
	if err != nil {
		return "false", fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	if Defined != nil {
		return "false", fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	res := newAddDefinedSet(Type, CommunitySetName, nil, list)
	_, err = Client.AddDefinedSet(ctx, res)
	if err != nil {
		return "false", fmt.Errorf("AddDefinedSet happen a err, err is %s", err)
	}
	return "Successful", nil
}

func AddDefinedSetNeighborSet(ctx context.Context, NeighborSetName, Type string, list []string) (string, error) {
	var err error
	Defined, err := ListDefinedSets(ctx, NeighborSetName)
	if err != nil {
		return "false", fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	if Defined != nil {
		return "false", fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	res := newAddDefinedSet(Type, NeighborSetName, nil, list)
	_, err = Client.AddDefinedSet(ctx, res)
	if err != nil {
		return "false", fmt.Errorf("AddDefinedSet happen a err, err is %s", err)
	}
	return "Successful", nil
}

func DeleteDefinedSets(ctx context.Context, DefinedSetName string) (string, error) {
	var err error
	Defined, err := ListDefinedSets(ctx, DefinedSetName)
	if err != nil {
		return "false", fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	if Defined != nil {
		return "false", errors.New("The DefinedSet is exist ")
	}
	Del := newDelDefinedSet(Defined)
	_, err = Client.DeleteDefinedSet(ctx, Del)
	if err != nil {
		return "false", fmt.Errorf("DeleteDefinedSet happen a err, err is %s", err)
	}
	return "Successful", nil
}

func ListDefinedSets(ctx context.Context, DefinedSetName string) (*gobgpapi.DefinedSet, error) {
	ListDefinedSetClient, err := Client.ListDefinedSet(ctx, &gobgpapi.ListDefinedSetRequest{Name: DefinedSetName})
	if err != nil {
		return nil, fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	DefinedSetResponse, err := ListDefinedSetClient.Recv()
	if err != nil {
		return nil, fmt.Errorf("ListDefinedSets happen a err, err is %s", err)
	}
	return DefinedSetResponse.GetDefinedSet(), nil
}