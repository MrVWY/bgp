package operation

import (
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
)

//ApplyPolicy
func newApplyPolicy(PolicyAssignmentName, Direction, RouteAction, PolicyName, ImOrOut string) (*gobgpapi.ApplyPolicy, error) {
	if ImOrOut == "Import" {
		Import, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, PolicyName)
		if err != nil {
			return nil, err
		}
		return &gobgpapi.ApplyPolicy{
			ImportPolicy: Import,
		}, nil
	}else if ImOrOut == "Export" {
		Export, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, PolicyName)
		if err != nil {
			return nil, err
		}
		return &gobgpapi.ApplyPolicy{
			ExportPolicy: Export,
		}, nil
	}else if ImOrOut == "ImportAndExport" {
		ImportAndExport, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, PolicyName)
		if err != nil {
			return nil, err
		}
		return &gobgpapi.ApplyPolicy{
			ImportPolicy: ImportAndExport,
			ExportPolicy: ImportAndExport,
		}, nil
	}
	return nil, fmt.Errorf("")
}

func newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, PolicyName string) (*gobgpapi.PolicyAssignment, error) {
	policy, err := ListPolicies(context.Background(), PolicyName)
	if err != nil {
		return nil, fmt.Errorf("ListPolicies happen a err, err is %s", err)
	}
	Policies := make([]*gobgpapi.Policy, 0)
	Policies = append(Policies, policy)

	var direction gobgpapi.PolicyDirection
	switch Direction {
	case "Import":
		direction = gobgpapi.PolicyDirection_IMPORT
	case "Export":
		direction = gobgpapi.PolicyDirection_EXPORT
	case "Unknown":
		direction = gobgpapi.PolicyDirection_UNKNOWN
	}

	var routeAction gobgpapi.RouteAction
	switch RouteAction {
	case "Accept":
		routeAction = gobgpapi.RouteAction_ACCEPT
	case "Reject":
		routeAction = gobgpapi.RouteAction_REJECT
	case "None":
		routeAction = gobgpapi.RouteAction_NONE
	}

	return &gobgpapi.PolicyAssignment{
		Name:          PolicyAssignmentName,
		Direction:     direction,
		Policies:      Policies,
		DefaultAction: routeAction,
	}, nil
}