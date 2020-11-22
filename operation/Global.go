package operation

import (
	"context"
	gobgpapi "github.com/osrg/gobgp/api"
)

func EditGlobalParameter(ctx context.Context, policy *gobgpapi.Policy) {
	Global := GetBgp(ctx)
	Global.ApplyPolicy.ImportPolicy.Name = policy.Name
	Global.ApplyPolicy.ImportPolicy.DefaultAction = gobgpapi.RouteAction_REJECT
}