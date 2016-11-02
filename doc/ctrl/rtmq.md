# RTMQ-PROXY组件接口
## 初始化接口
---
>函数名称:RtmqProxyInit()<br>
>功能描述:初始化RTMQ-PROXY上下文<br>
>参数列表:<br>
>  conf: 配置信息<br>
>  log: 日志对象<br>
>注意事项: 暂无<br>

#### 配置信息

|**序号**|**变量名**|**数据类型**|**描述**|**备注**|
|:------:|:------|:-------|:-------|:---------|:-------|
| 01 | NodeId | uint32 | 结点ID | 暂无 |
| 02 | RemoteAddr | string | 远端服务地址 | IP+PORT |
| 03 | WorkerNum | uint32 | 工作协程数 | 暂无 |
| 04 | SendChanLen | uint32 | 发送队列长度 | 暂无 |
| 05 | RecvChanLen | uint32 | 接收队列长度 | 暂无 |

## 回调注册接口
---
>函数名称:Register()<br>
>功能描述:注册处理回调<br>
>参数列表:<br>
>  cmd: 消息类型<br>
>  proc: 处理回调<br>
>  param: 附加参数<br>
>注意事项: 暂无<br>

## 数据发送接口
---
>函数名称:Send()<br>
>功能描述:发送数据<br>
>参数列表:<br>
>  cmd: 消息类型<br>
>  data: 被发送的数据<br>
>  length: 数据长度<br>
>注意事项: 暂无<br>