package operation

import (
	"errors"
	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
	"log"
)

func AddDefinedSets(ctx context.Context,Type, DFSetName, ipPrefix, MaskMin, MaskMax string) (string, error) {
	var err error
	Defined, err := ListDefinedSets(ctx, DFSetName)
	if err != nil {
		return "false", err
	}
	if Defined != nil {
		return "false", errors.New("The DefinedSet is exist ")
	}
	Prefix := newPrefixSet(ipPrefix, MaskMin, MaskMax)
	res := newAddDefinedSet(Type, DFSetName, Prefix)
	response, err := Client.AddDefinedSet(ctx, res)
	if err != nil {
		return "false", err
	}
	return response.String(), nil
}

func DeleteDefinedSets(ctx context.Context, DFSetName string) (string, error) {
	var err error
	Defined, err := ListDefinedSets(ctx, DFSetName)
	if err != nil {
		return "false", err
	}
	if Defined != nil {
		return "false", errors.New("The DefinedSet is exist ")
	}
	Del := newDelDefinedSet(Defined)
	response, err := Client.DeleteDefinedSet(ctx, Del)
	if err != nil {
		return "false", err
	}
	return response.String(), nil
}

func ListDefinedSets(ctx context.Context, DefinedSetName string) (*gobgpapi.DefinedSet, error) {
	ListDefinedSetClient, err := Client.ListDefinedSet(ctx, &gobgpapi.ListDefinedSetRequest{Name: DefinedSetName})
	if err != nil {
		log.Fatalf("ListDefinedSets happen a err, err is %s", err)
	}
	DefinedSetResponse, err := ListDefinedSetClient.Recv()
	if err != nil {
		log.Fatalf("ListDefinedSets happen a err, err is %s", err)
	}
	return DefinedSetResponse.GetDefinedSet(), nil
}