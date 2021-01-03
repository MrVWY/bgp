### BGP Route Policy Structure
[policy-component.png](https://github.com/osrg/gobgp/blob/master/docs/sources/policy-component.png)

# 问题peer 10.15.0.1 doesn't have per peer policy ->  Set " route-server-client = true "  

### 接口  (22个)
+ basic class
    + StartBGP()  
    
    + StopBgp()
    
    + GetBgp()
    
+ Policy class
    + CreatePolicy() ✔
        + json parameter:  
        {  
           PolicyName: 策略名  
           StatementsName: 策略内容  
           PrefixSetName: 路由前缀列表名称  
           NeighborSetName: 对等体列表名称  
           CommunitySetName: 团体属性列表名称  
           CommunityAction: 团体属性的动作  
           Community: 设置团体属性
           Action: 路由动作  
           NextHop: 下一跳  
        }
    + DeletePolicy() ✔
        + json parameter:  
        {  
          PolicyName: 策略名  
        }
    + ListPolicy() ✔
        + json parameter:  
        {  
          PolicyName: 策略名  
        }
    + AddStatementToPolicy() ✔  
        + json parameter:  
        {  
           PolicyName: 策略名  
           StatementsName: 策略内容   
        }  

+ Peer class
    + CreatePeer()✔  
      + json parameter:  
      {  
 
      }
    + DeletePeer()✔  
      + json parameter:  
      {  

      }
    + ListPeer()✔  
      + json parameter:  
      {  
  
      }
    + UpdatePeer()✔  
      + json parameter:  
      {  
 
      }
    + ResetPeer()
    + ShutdownPeer()
    + EnablePeer()
    + DisablePeer()
    + MonitorPeer()
    
+ Defined class
    + CreatePrefixSet() ✔  
        + json parameter:  
        {  
            Type: Defined类型    
            PrefixSetName: 名称  
            ipPrefix: 路由前缀  
            MaskMin: 最小掩码长度  
            MaskMax: 最大掩码长度  
        }  
    + CreateCommunitySet() ✔  
        + json parameter:
        {
            Type: Defined类型  
            CommunitySetName: 团体属性列表名称  
            list: 团体属性值
            
        }  
    + CreateNeighborSet() ✔  
        + json parameter:
        {
            Type: Defined类型  
            CommunitySetName: 对等体列表名称  
            list: 对等体值            
        }  
    + DeleteDefined() ✔  
        + json parameter:  
        {  
            DefinedSetName: Defined名称  
        }  
        
+ Statement class  
    + CreateStatement() ✔  
        + json parameter:  
        {  
            StatementsName: 策略内容名称  
            PrefixSetName: 路由前缀列表名称  
            NeighborSetName: 对等体列表名称
            CommunitySetName: 团体属性列表名称  
            CommunityAction: 团体属性的动作  
            Community: 设置团体属性
            action: 路由动作               
        }  
    + DeleteStatement() ✔
        + json parameter:  
        {  
            StatementsName: 策略内容名称              
        }      
    + ListStatement() ✔  
        + json parameter:  
        {  
            StatementsName: 策略内容名称              
        }         