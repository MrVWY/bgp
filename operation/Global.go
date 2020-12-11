package operation

import (
	"context"
	gobgpapi "github.com/osrg/gobgp/api"
)

func EditGlobalParameter(ctx context.Context, policyName string) (string, error) {
	Global := GetBgp(ctx)
	Global.ApplyPolicy.ImportPolicy.Name = policyName
	Global.ApplyPolicy.ImportPolicy.DefaultAction = gobgpapi.RouteAction_REJECT
	return "Successful", nil
}