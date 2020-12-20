package operation

import (
	gobgpapi "github.com/osrg/gobgp/api"
)

//Policy
func newAddPolicyRequest(policyName, StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string, Community []string, NextHopAddress string) *gobgpapi.AddPolicyRequest {
	return &gobgpapi.AddPolicyRequest{
		Policy:                  newPolicy(policyName, StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action, Community, NextHopAddress),
		ReferExistingStatements: false,
	}
}

func newDelPolicyRequest(policy *gobgpapi.Policy) *gobgpapi.DeletePolicyRequest {
	return &gobgpapi.DeletePolicyRequest{
		Policy:             policy,
		PreserveStatements: true,
		All:                false,
	}
}

func newPolicy(policyName, StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string, Community []string, NextHopAddress string) *gobgpapi.Policy {
	return &gobgpapi.Policy{
		Name:       policyName,
		Statements: []*gobgpapi.Statement{newStatements(StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action, Community, NextHopAddress)},
	}
}

//Statements
func newStatements(StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string, Community []string, NextHopAddress string) *gobgpapi.Statement {
	return &gobgpapi.Statement{
		Name: StatementsName,
		Conditions: &gobgpapi.Conditions{
			PrefixSet:    newMatchSet(PrefixSetName),
			NeighborSet:  newMatchSet(NeighborSetName),
			CommunitySet: newMatchSet(CommunitySetName),
		},
		Actions: &gobgpapi.Actions{
			RouteAction: selectRouteAction(action),
			Community:   newCommunityAction(CommunityAction, Community),
			Nexthop:     newNextHopAction(NextHopAddress),
		},
	}
}

func newDelStatements(Statements *gobgpapi.Statement) *gobgpapi.DeleteStatementRequest {
	return &gobgpapi.DeleteStatementRequest{
		Statement: Statements,
		All:       false,
	}
}

func newAddStatementRequest(StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string, Community []string, NextHopAddress string) *gobgpapi.AddStatementRequest {
	return &gobgpapi.AddStatementRequest{
		Statement: newStatements(StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action, Community, NextHopAddress),
	}
}

//Peer
func newAddPeerRequest(Description, NeighborAddress string, PeerAs, SendCommunity int) *gobgpapi.AddPeerRequest {
	return &gobgpapi.AddPeerRequest{Peer: newPeer(Description, NeighborAddress, PeerAs, SendCommunity)}
}

func newUpdatePeerRequest(newPeer *gobgpapi.Peer, DoSoftResetIn bool) *gobgpapi.UpdatePeerRequest {
	return &gobgpapi.UpdatePeerRequest{
		Peer:          newPeer,
		DoSoftResetIn: DoSoftResetIn, //软复位 可以在不中断BGP连接的情况下使BGP立即应用新的出口策略
	}
}

func newPeer(Description, NeighborAddress string, PeerAs, SendCommunity int) *gobgpapi.Peer {
	return &gobgpapi.Peer{
		ApplyPolicy:  nil,
		Conf:         newPeerConf(Description, NeighborAddress, PeerAs, SendCommunity),
		EbgpMultihop: &gobgpapi.EbgpMultihop{Enabled: true, MultihopTtl: 100},
		//State:          &gobgpapi.PeerState{},
		//Transport:      newTransport(),
		//RouteReflector: nil,
	}
}

//prepare to Loopback
func newTransport(LocalAddress, RemoteAddress string, LocalPort, RemotePort int) *gobgpapi.Transport {
	return &gobgpapi.Transport{
		LocalAddress: LocalAddress,
		LocalPort:    uint32(LocalPort),
		MtuDiscovery: false, //链路MTU最大值检测
		//neighbor passive
		//The neighbor passive command sets the TCP connection for the specified BGP neighbor or peer group to passive mode. When the peer’s transport connection mode is set to passive, it accepts TCP connections for BGP but does not initiate them.
		//The no neighbor passive command sets the specified BGP neighbor or peer group to active connection mode. BGP peers in active mode can both accept and initiate TCP connections for BGP. This is the default behavior.
		//The default neighbor passive command restores the default connection mode. The default mode is “active” for individual BGP peers, or the mode inherited from the peer group for peer group members.
		//neighbour ×.×.×.× transport connection-mode active/passive
		PassiveMode:   false,
		RemoteAddress: RemoteAddress,
		RemotePort:    uint32(RemotePort),
		//TcpMss: 0, //tcp_mss 最大分段长度
	}
}

//http://blog.sina.com.cn/s/blog_5ec3537101019suy.html
//no-prepend： Do not prepend local-as to updates from ebgp peers
func newPeerConf(Description, NeighborAddress string, PeerAs, SendCommunity int) *gobgpapi.PeerConf {
	return &gobgpapi.PeerConf{
		Description: Description,
		//local_as
		//Take a look at the topology below. If the ISP who owns AS 100 brought the ISP that owns AS 200,
		//then you’re gonna wanna eventually get R3 using AS 100. However, if on R3, we replace the BGP AS 200 process with 100,
		//then every neighbor will go down until they re-peer to the new AS number on R3. You’re extremely unlikely, in the real world,
		//t o get all your peers to adjust their config at the same time. So a good migration strategy would be to keep R3 using AS 200,
		//and when a peer is ready to change their config to peer with the new AS, we change R3’s AS number for that peer only.
		//This means all sessions with the other peers will be maintained because R3 is still peering with everyone else using AS 200. Let’s take a look at an example of how the local-as feature can be used to accomplish this.
		//LocalAs:         uint32(LocalAs),
		NeighborAddress:  NeighborAddress,
		PeerAs:           uint32(PeerAs),
		SendCommunity:    uint32(SendCommunity),
		RouteFlapDamping: false, //路由抖动 reducing the number of update messages sent between BGP peers
		AllowOwnAs:       1,
		ReplacePeerAs:    false, //前提是出现了secondary AS, 否则默认false
	}
}

func newPeerSate(Description, NeighborAddress, RouterId string, LocalAs, PeerAs int) *gobgpapi.PeerState {
	return &gobgpapi.PeerState{
		Description:      Description,
		LocalAs:          uint32(LocalAs),
		NeighborAddress:  NeighborAddress,
		PeerAs:           uint32(PeerAs),
		RouteFlapDamping: false,
		RouterId:         RouterId,
	}
}