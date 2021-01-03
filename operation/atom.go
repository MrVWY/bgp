package operation

import gobgpapi "github.com/osrg/gobgp/api"

//MatchSet
func newMatchSet(name string) *gobgpapi.MatchSet {
	if name == "" {
		return nil
	}
	return &gobgpapi.MatchSet{MatchType: gobgpapi.MatchType_ANY, Name: name}
}

//CommunityAction
func newCommunityAction(CommunityAction string, Community []string) *gobgpapi.CommunityAction {
	var communityAction gobgpapi.CommunityActionType
	if CommunityAction == "" {
		return nil
	}else {
		switch CommunityAction {
		case "ADD":
			communityAction = gobgpapi.CommunityActionType_COMMUNITY_ADD
		case "REMOVE":
			communityAction = gobgpapi.CommunityActionType_COMMUNITY_REMOVE
		case "REPLACE":
			communityAction = gobgpapi.CommunityActionType_COMMUNITY_REPLACE
		}
	}
	return &gobgpapi.CommunityAction{ActionType:  communityAction, Communities: Community}
}

//NextHopAction  注意EBGP IBGP nexthop
func newNextHopAction(NextHopAddress string) *gobgpapi.NexthopAction {
	if NextHopAddress == "" {
		return nil
	}
	return &gobgpapi.NexthopAction{ Address: NextHopAddress, Self: true }
}

//selectRouteAction
func selectRouteAction(RouteAction string) gobgpapi.RouteAction {
	var Action gobgpapi.RouteAction
	switch RouteAction {
	case "NONE":
		Action = gobgpapi.RouteAction_NONE
	case "ACCEPT":
		Action = gobgpapi.RouteAction_ACCEPT
	case "REJECT":
		Action = gobgpapi.RouteAction_REJECT
	}
	return Action
}