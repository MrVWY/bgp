package operation

import (
	"context"
	gobgpapi "github.com/osrg/gobgp/api"
)

//Global
func newGlobal(As int, RouterID string, ListenPort int, ListenAddresses []string, UseMultiplePaths bool) *gobgpapi.Global {
	return &gobgpapi.Global{
		As:               uint32(As),
		RouterId:         RouterID,
		ListenPort:       int32(ListenPort),
		ListenAddresses:  ListenAddresses,
		UseMultiplePaths: UseMultiplePaths,
	}
}

func EditGlobalParameter(ctx context.Context, policyName, InOrOut, action string) (string, error) {
	Global := GetBgp(ctx)
	Global.ApplyPolicy.ImportPolicy.Name = policyName
	Global.ApplyPolicy.ImportPolicy.DefaultAction = gobgpapi.RouteAction_REJECT

	if InOrOut == "Import" {

	}else if InOrOut == "Export" {

	}else if InOrOut == "ImportAndExport" {

	}

	return "Successful", nil
}