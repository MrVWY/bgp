package handle

import (
	"bgp/operation"
	"golang.org/x/net/context"
)

func StartBGP() {

}

func CreateGlobalPolicy()  {

}

func CreatePolicy(ctx context.Context, PolicyName, StatementsName, PrefixSetName, NeighborSetName string) {
	_, _ = operation.AddPolicies(ctx, PolicyName, StatementsName, PrefixSetName, NeighborSetName)
}

func CreateStatement(ctx context.Context, StatementsName, PrefixSetName, NeighborSetName string) {
	_, _ = operation.AddStatements(ctx, StatementsName, PrefixSetName, NeighborSetName)
}

func CreatePrefixSet(ctx context.Context, Type, SetName, ipPrefix, MaskMin, MaskMax string) {
	_, _ = operation.AddDefinedSets(ctx, Type, SetName, ipPrefix, MaskMin, MaskMax)
}

func AddStatementToPolicy(ctx context.Context, PolicyName, StatementName string) {
	_, _ = operation.AddStatementToPolicy(ctx, PolicyName, StatementName)
}