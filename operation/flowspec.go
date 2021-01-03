package operation

import (
	"fmt"
	"os/exec"
)

//Remotely-Triggered Black Hole RTHB
//    <MATCH> : { destination <PREFIX> [<OFFSET>] |
//                protocol <PROTOCOLS>... |
//                tcp-flags <TCP_FLAGS>... |
//                port <ITEM>... |
//                destination-port <ITEM>... |
//                icmp-type <ITEM>... |
//                icmp-code <ITEM>... |

//TYPE	CODE	Description	Query	Error
//0	0	Echo Reply——回显应答（Ping应答）	x
//3	0	Network Unreachable——网络不可达	 	x
//3	1	Host Unreachable——主机不可达	 	x
//3	2	Protocol Unreachable——协议不可达	 	x
//3	3	Port Unreachable——端口不可达	 	x
//3	4	Fragmentation needed but no frag. bit set——需要进行分片但设置不分片比特	 	x
//3	5	Source routing failed——源站选路失败	 	x
//3	6	Destination network unknown——目的网络未知	 	x
//3	7	Destination host unknown——目的主机未知	 	x
//3	8	Source host isolated (obsolete)——源主机被隔离（作废不用）	 	x
//3	9	Destination network administratively prohibited——目的网络被强制禁止	 	x
//3	10	Destination host administratively prohibited——目的主机被强制禁止	 	x
//3	11	Network unreachable for TOS——由于服务类型TOS，网络不可达	 	x
//3	12	Host unreachable for TOS——由于服务类型TOS，主机不可达	 	x
//3	13	Communication administratively prohibited by filtering——由于过滤，通信被强制禁止	 	x
//3	14	Host precedence violation——主机越权	 	x
//3	15	Precedence cutoff in effect——优先中止生效	 	x
//4	0	Source quench——源端被关闭（基本流控制）
//5	0	Redirect for network——对网络重定向
//5	1	Redirect for host——对主机重定向
//5	2	Redirect for TOS and network——对服务类型和网络重定向
//5	3	Redirect for TOS and host——对服务类型和主机重定向
//8	0	Echo request——回显请求（Ping请求）	x
//9	0	Router advertisement——路由器通告
//10	0	Route solicitation——路由器请求
//11	0	TTL equals 0 during transit——传输期间生存时间为0	 	x
//11	1	TTL equals 0 during reassembly——在数据报组装期间生存时间为0	 	x
//12	0	IP header bad (catchall error)——坏的IP首部（包括各种差错）	 	x
//12	1	Required options missing——缺少必需的选项	 	x
//13	0	Timestamp request (obsolete)——时间戳请求（作废不用）	x
//14	 	Timestamp reply (obsolete)——时间戳应答（作废不用）	x
//15	0	Information request (obsolete)——信息请求（作废不用）	x
//16	0	Information reply (obsolete)——信息应答（作废不用）	x
//17	0	Address mask request——地址掩码请求	x
//18	0	Address mask reply——地址掩码应答

var (
	GoBgpPath = "./home/BGP/gobgp"
)

type match struct {
	destination     string
	protocol        string
	port            string
	destinationPort string
	icmpType        string
	icmpCode        string
	dscp            string
}

//gobgp global rib -a ipv4-flowspec add match destination 10.0.0.0/24 then accept
//gobgp global rib -a ipv4-flowspec

func AddFlowSpec(destination, protocol, destinationPort, action string) (string, error) {
	cmd := exec.Command(GoBgpPath, "global rib", "-a", "ipv4-flowspec", fmt.Sprintf("add match destination %s destinationPort %s protocol %s", destination, destinationPort, protocol),
		fmt.Sprintf("then %s", action))

	if err := cmd.Start(); err != nil {
		return "false", fmt.Errorf("cmd start happen a err, %s", err)
	}
	return "true", nil
}


