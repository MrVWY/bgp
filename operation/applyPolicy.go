package operation

import (
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
	"golang.org/x/net/context"
)

//ApplyPolicy
func newApplyPolicy(PolicyAssignmentName, Direction, RouteAction, importPolicyName, exportPolicy, ImOrOut string) (*gobgpapi.ApplyPolicy, error) {
	if ImOrOut == "Import" {
		Import, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, importPolicyName)
		if err != nil {
			return nil, err
		}
		return &gobgpapi.ApplyPolicy{
			ImportPolicy: Import,
		}, nil
	}else if ImOrOut == "Export" {
		ExportPolicy, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, exportPolicy)
		if err != nil {
			return nil, err
		}
		return &gobgpapi.ApplyPolicy{
			ExportPolicy: ExportPolicy,
		}, nil
	}else if ImOrOut == "ImportAndExport" {
		ImportAndExport, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, importPolicyName)
		if err != nil {
			return nil, err
		}
		ExportPolicy, err := newPolicyAssignment(PolicyAssignmentName, Direction, RouteAction, exportPolicy)
		if err != nil {
			return nil, err
		}
		return &gobgpapi.ApplyPolicy{
			ImportPolicy: ImportAndExport,
			ExportPolicy: ExportPolicy,
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

	return &gobgpapi.PolicyAssignment{
		Name:          PolicyAssignmentName,
		Direction:     direction,
		Policies:      Policies,
		DefaultAction: selectRouteAction(RouteAction),
	}, nil
}