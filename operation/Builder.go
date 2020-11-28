package operation

import (
	gobgpapi "github.com/osrg/gobgp/api"
	"strconv"
)

//Policy
func newAddPolicyRequest(policyName, StatementsName, PrefixSetName, NeighborSetName string) *gobgpapi.AddPolicyRequest {
	return &gobgpapi.AddPolicyRequest{
		Policy: newPolicy(policyName, StatementsName, PrefixSetName, NeighborSetName),
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

func newPolicy(policyName, StatementsName, PrefixSetName, NeighborSetName string) *gobgpapi.Policy {
	return &gobgpapi.Policy{
		Name : policyName,
		Statements: []*gobgpapi.Statement{ newStatements(StatementsName, PrefixSetName, NeighborSetName) },
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
func newStatements(StatementsName ,PrefixSetName, NeighborSetName string) *gobgpapi.Statement {
	return &gobgpapi.Statement{
		Name: StatementsName,
		Conditions: &gobgpapi.Conditions{
			PrefixSet: &gobgpapi.MatchSet{MatchType: gobgpapi.MatchType_ANY, Name: PrefixSetName},
			NeighborSet:       &gobgpapi.MatchSet{MatchType: gobgpapi.MatchType_ANY, Name: NeighborSetName},
		},
		Actions: &gobgpapi.Actions{
			RouteAction: gobgpapi.RouteAction_REJECT,
		},
	}
}

func newDelStatements(Statements *gobgpapi.Statement) *gobgpapi.DeleteStatementRequest {
	return &gobgpapi.DeleteStatementRequest{
		Statement: Statements,
		All: false,
	}
}

func newAddStatementRequest(StatementsName, PrefixSetName, NeighborSetName string) *gobgpapi.AddStatementRequest{
	return &gobgpapi.AddStatementRequest{
		Statement: newStatements(StatementsName, PrefixSetName, NeighborSetName),
	}
}

//Global
func newGlobal(As int, RouterID string, ListenPort int, ListenAddresses []string, UseMultiplePaths bool) *gobgpapi.Global {
	return &gobgpapi.Global{
		As: 				uint32(As),
		RouterId: 			RouterID,
		ListenPort: 		int32(ListenPort),
		ListenAddresses: 	ListenAddresses,
		UseMultiplePaths:      UseMultiplePaths,
	}
}

//Peer
func newAddPeerRequest(Description, NeighborAddress string, LocalAs, PeerAs, SendCommunity int) *gobgpapi.AddPeerRequest {
	return &gobgpapi.AddPeerRequest{ Peer: newPeer(Description, NeighborAddress, LocalAs, PeerAs, SendCommunity) }
}

func newPeer(Description, NeighborAddress string, LocalAs, PeerAs, SendCommunity int) *gobgpapi.Peer {
	return &gobgpapi.Peer{
		Conf:            newPeerConf(Description, NeighborAddress, LocalAs, PeerAs, SendCommunity),
		EbgpMultihop:    &gobgpapi.EbgpMultihop{ Enabled: false,  MultihopTtl: 0 },
		//State:           nil,
		//Timers:          nil,
		//Transport:       nil, //loopback
	}
}

//prepare to Loopback
func newTransport() *gobgpapi.Transport {
	return &gobgpapi.Transport{
		LocalAddress:  "",
		LocalPort:     0,
		MtuDiscovery:  false,
		PassiveMode:   false,
		RemoteAddress: "",
		RemotePort:    0,
		TcpMss:        0,
	}
}

func newPeerConf(Description, NeighborAddress string, LocalAs, PeerAs, SendCommunity int) *gobgpapi.PeerConf {
	return &gobgpapi.PeerConf{
		Description:       Description,
		LocalAs:           uint32(LocalAs),
		NeighborAddress:   NeighborAddress,
		PeerAs:            uint32(PeerAs),
		SendCommunity:     uint32(SendCommunity),
		//NeighborInterface: "",
	}
}