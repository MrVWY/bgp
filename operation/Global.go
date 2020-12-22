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

func EditGlobalParameter(ctx context.Context, importPolicyName, exportPolicy, InOrOut, action string) (string, error) {
	Global := GetBgp(ctx)
	if InOrOut == "Import" {
		Global.ApplyPolicy.ImportPolicy.Name = importPolicyName
		Global.ApplyPolicy.ImportPolicy.DefaultAction = selectRouteAction(action)
	}else if InOrOut == "Export" {
		Global.ApplyPolicy.ExportPolicy.Name = exportPolicy
		Global.ApplyPolicy.ExportPolicy.DefaultAction = selectRouteAction(action)
	}else if InOrOut == "ImportAndExport" {
		Global.ApplyPolicy.ImportPolicy.Name = importPolicyName
		Global.ApplyPolicy.ImportPolicy.DefaultAction = selectRouteAction(action)
		Global.ApplyPolicy.ExportPolicy.Name = exportPolicy
		Global.ApplyPolicy.ExportPolicy.DefaultAction = selectRouteAction(action)
	}
	return "Successful", nil
}