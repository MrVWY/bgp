package operation

import (
	gobgpapi "github.com/osrg/gobgp/api"
	"strconv"
)

//Policy
func newAddPolicyRequest(policyName, StatementsName, PrefixSetName string) *gobgpapi.AddPolicyRequest {
	return &gobgpapi.AddPolicyRequest{
		Policy: newPolicy(policyName, StatementsName, PrefixSetName),
		ReferExistingStatements: false,
	}
}

func newDelPolicyRequest(policy *gobgpapi.Policy) *gobgpapi.DeletePolicyRequest {
	return &gobgpapi.DeletePolicyRequest{
		Policy: policy,
		PreserveStatements: true,
		All: false,
	}
}

func newPolicy(policyName, StatementsName, PrefixSetName string) *gobgpapi.Policy {
	return &gobgpapi.Policy{
		Name : policyName,
		Statements: []*gobgpapi.Statement{ newStatements(StatementsName, PrefixSetName) },
	}
}

//DefinedSet
func newAddDefinedSet(Type string, DFSetName string, prefix []*gobgpapi.Prefix) *gobgpapi.AddDefinedSetRequest {
	return &gobgpapi.AddDefinedSetRequest{
		DefinedSet: newDefinedSet(Type, DFSetName, prefix),
	}
}

func newDelDefinedSet(DefinedSet *gobgpapi.DefinedSet) *gobgpapi.DeleteDefinedSetRequest {
	return &gobgpapi.DeleteDefinedSetRequest{
		DefinedSet: DefinedSet,
	}
}

func newDefinedSet(Type string, DFSetName string, prefix []*gobgpapi.Prefix) *gobgpapi.DefinedSet {
	var DefinedType gobgpapi.DefinedType

	switch Type {
		case "PREFIX": DefinedType = gobgpapi.DefinedType_PREFIX
		case "NEIGHBOR": DefinedType = gobgpapi.DefinedType_NEIGHBOR
		case "TAG": DefinedType = gobgpapi.DefinedType_TAG
		case "AS_PATH" : DefinedType = gobgpapi.DefinedType_AS_PATH
		case "COMMUNITY": DefinedType = gobgpapi.DefinedType_COMMUNITY
		case "EXT_COMMUNITY" : DefinedType = gobgpapi.DefinedType_EXT_COMMUNITY
		case "LARGE_COMMUNITY": DefinedType = gobgpapi.DefinedType_LARGE_COMMUNITY
		case "NEXT_HOP" : DefinedType = gobgpapi.DefinedType_NEXT_HOP
	}
	return &gobgpapi.DefinedSet{
		DefinedType: 	DefinedType,
		Name: 			DFSetName,
		List:        	nil,
		Prefixes: 		prefix,
	}
}

//PrefixSet
func newPrefixSet(ipPrefix, MaskMin, MaskMax string) []*gobgpapi.Prefix {
	max, _ := strconv.ParseUint(MaskMax, 10, 32)
	min, _ := strconv.ParseUint(MaskMin, 10, 32)
	return []*gobgpapi.Prefix{
		{ipPrefix, uint32(min), uint32(max)},
	}
}

//Statements
func newStatements(StatementsName ,PrefixSetName string) *gobgpapi.Statement {
	return &gobgpapi.Statement{
		Name: StatementsName,
		Conditions: &gobgpapi.Conditions{
			PrefixSet: &gobgpapi.MatchSet{MatchType: gobgpapi.MatchType_ANY, Name: PrefixSetName},
			NeighborSet:       nil,
			AsPathLength:      nil,
			AsPathSet:         nil,
			CommunitySet:      nil,
			ExtCommunitySet:   nil,
			RpkiResult:        0,
			RouteType:         0,
			LargeCommunitySet: nil,
			NextHopInList:     nil,
			AfiSafiIn:         nil,
		},
		Actions: &gobgpapi.Actions{
			RouteAction: gobgpapi.RouteAction_REJECT,
			Community:      nil,
			Med:            nil,
			AsPrepend:      nil,
			ExtCommunity:   nil,
			Nexthop:        nil,
			LocalPref:      nil,
			LargeCommunity: nil,
		},
	}
}

func newDelStatements(Statements *gobgpapi.Statement) *gobgpapi.DeleteStatementRequest {
	return &gobgpapi.DeleteStatementRequest{
		Statement: Statements,
		All: false,
	}
}

func newAddStatementRequest(StatementsName, PrefixSetName string) *gobgpapi.AddStatementRequest{
	return &gobgpapi.AddStatementRequest{
		Statement: newStatements(StatementsName, PrefixSetName),
	}
}

//MatchSet
func newMatchSet(Type, MatchName string) *gobgpapi.MatchSet {
	var MatchTypes gobgpapi.MatchType
	switch Type {
	case "ANY": MatchTypes = gobgpapi.MatchType_ANY
	case "ALL": MatchTypes = gobgpapi.MatchType_ALL
	case "INVERT": MatchTypes = gobgpapi.MatchType_INVERT
	}
	return &gobgpapi.MatchSet{
		MatchType: MatchTypes,
		Name: MatchName,
	}
}

//Global
func newGlobal(As int, RouterID string, ListenPort int, ListenAddresses []string) *gobgpapi.Global {
	return &gobgpapi.Global{
		As: 				uint32(As),
		RouterId: 			RouterID,
		ListenPort: 		int32(ListenPort),
		ListenAddresses: 	ListenAddresses,
		Families:              nil,
		UseMultiplePaths:      false,
		RouteSelectionOptions: nil,
		DefaultRouteDistance:  nil,
		Confederation:         nil,
		GracefulRestart:       nil,
		ApplyPolicy:           nil,
	}
}

//Peer
func newPeer() *gobgpapi.Peer {
	return &gobgpapi.Peer{
		Conf: newPeerConf(),
		Transport: &gobgpapi.Transport{
			LocalAddress:  "",
			LocalPort:     0,
			MtuDiscovery:  false,
			PassiveMode:   false,
			RemoteAddress: "",
			RemotePort:    0,
			TcpMss:        0,
		},
	}
}

func newPeerConf() *gobgpapi.PeerConf {
	return &gobgpapi.PeerConf{
		Description: "",
		NeighborAddress: "123",
		PeerAs: 12,
	}
}