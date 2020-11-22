package handle

import (
	"bgp/operation"
	"golang.org/x/net/context"
)

func StartBGP() {

}

func CreateGlobalPolicy()  {

}

func CreatePolicy() {

}

func CreateStatement(ctx context.Context, StatementsName, PrefixSetName string) {
	_, _ = operation.AddStatements(ctx, StatementsName, PrefixSetName)
}

func CreatePrefixSet(ctx context.Context, Type, SetName, ipPrefix, MaskMin, MaskMax string) {
	_, _ = operation.AddDefinedSets(ctx, Type, SetName, ipPrefix, MaskMin, MaskMax)

}

func AddStatementToPolicy(ctx context.Context, PolicyName, StatementName string) {
	_, _ = operation.AddStatementToPolicy(ctx, PolicyName, StatementName)
}


//[global.config]
//
//as = 65000
//
//router-id = "10.0.0.5"
//
//
//
//[[neighbors]]
//
//[neighbors.config]
//
//neighbor-address = "10.0.0.1"
//
//peer-as = 65000
//
//[[neighbors.afi-safis]]
//
//[neighbors.afi-safis.config]
//
//afi-safi-name = "l2vpn-evpn"
//
//[neighbors.transport.config]
//
//local-address = "10.15.0.100"