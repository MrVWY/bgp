### BGP Route Policy Structure
[policy-component.png](https://github.com/osrg/gobgp/blob/master/docs/sources/policy-component.png)

### 接口  (22个) 后期只限post
+ basic class
    + StartBGP()  
    
    + StopBgp()
    
    + GetBgp()
    
+ Policy class
    + CreatePolicy()  
        + json parameter:  
        {  
           PolicyName: 策略名  
           StatementsName: 策略内容  
           PrefixSetName: 路由前缀  
           NeighborSetName: 对等体  
        }  
    
    + DeletePolicy()
        + json parameter:  
        {  
          PolicyName: 策略名  
        }  
        
    + ListPolicy()
        + json parameter:  
        {  
          PolicyName: 策略名  
        }  
        
    + AddStatementToPolicy()  
        + json parameter:  
        {  
           PolicyName: 策略名  
           StatementsName: 策略内容   
        }  

+ Peer class
    + CreatePeer()  
    + DeletePeer()
    + ListPeer()
    + UpdatePeer()
    + ResetPeer()
    + ShutdownPeer()
    + EnablePeer()
    + DisablePeer()
    + MonitorPeer()
    
+ Defined class
    + CreatePrefixSet()  
        + json parameter:  
        {  
            Type: 策略内容  
            SetName: 名称  
            ipPrefix: 路由前缀  
            MaskMin: 最小掩码长度  
            MaskMax: 最大掩码长度  
        }  
        
+ Statement class  
    + CreateStatement()  
        + json parameter:  
        {  
            StatementsName: 策略内容  
            PrefixSetName: 路由前缀  
            NeighborSetName: 对等体  
        }  
    + DeleteStatement()
    
    + ListStatement()