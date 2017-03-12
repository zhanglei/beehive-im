package controllers

import (
	"beehive-im/src/golang/lib/comm"
	"beehive-im/src/golang/lib/lws"
)

/******************************************************************************
 **函数名称: lsnd_conn_init
 **功    能: 初始化CONN对象
 **输入参数:
 **     client: 客户端对象
 **输出参数:
 **返    回: 0:正常 !0:异常
 **实现描述:
 **注意事项: 每个连接只会调用此函数1次
 **作    者: # Qifeng.zou # 2017.03.04 18:53:16 #
 ******************************************************************************/
func (ctx *LsndCntx) lsnd_conn_init(client *lws.Client) int {
	session := &LsndSessionExtra{
		sid:    0,
		cid:    client.GetCid(),
		status: CONN_STATUS_READY,
	}

	client.SetUserData(session)

	return 0
}

/******************************************************************************
 **函数名称: lsnd_conn_recv
 **功    能: 接收数据处理
 **输入参数:
 **     client: 客户端对象
 **     data: 收到的数据
 **     length: 收到数据的长度
 **输出参数:
 **返    回: 0:正常 !0:异常
 **实现描述: 获取消息类型, 并调用对应的处理回调
 **注意事项: 返回!0值将导致连接断开
 **作    者: # Qifeng.zou # 2017.03.04 15:21:43 #
 ******************************************************************************/
func (ctx *LsndCntx) lsnd_conn_recv(client *lws.Client, data []byte, length int) int {
	session, ok := client.GetUserData().(*LsndSessionExtra)
	if !ok {
		ctx.log.Error("Get connection extra data failed!")
		return -1
	}

	/* > 字节序转化 */
	head := comm.MesgHeadNtoh(data)
	head.SetNid(ctx.conf.GetNid())
	if !head.IsValid() {
		ctx.log.Error("Mesg head is invalid! cmd:0x%04X sid:%d nid:%d chksum:0x%X",
			head.GetCmd(), head.GetSid(), head.GetNid(), head.GetChkSum())
		return -1
	}

	/* > 查找&执行回调 */
	cb, param := ctx.callback.Query(head.GetCmd())
	if nil == cb {
		cb, param = ctx.callback.Query(comm.CMD_UNKNOWN)
		if nil == cb {
			ctx.log.Error("Didn't find command handler! cmd:0x%04X", head.GetCmd())
			return 0
		}
	}

	cb(session, head.GetCmd(), data, uint32(length), param)

	return 0
}

/******************************************************************************
 **函数名称: lsnd_conn_send
 **功    能: 发送数据处理
 **输入参数:
 **     client: 客户端对象
 **     data: 收到的数据
 **     length: 收到数据的长度
 **输出参数:
 **返    回: 0:正常 !0:异常
 **实现描述:
 **注意事项: 返回!0值将导致连接断开
 **作    者: # Qifeng.zou # 2017.03.04 15:21:43 #
 ******************************************************************************/
func (ctx *LsndCntx) lsnd_conn_send(client *lws.Client, data []byte, length int) int {
	session, ok := client.GetUserData().(*LsndSessionExtra)
	if !ok {
		ctx.log.Error("Get connection extra data failed!")
		return -1
	}

	head := comm.MesgHeadNtoh(data)

	ctx.log.Debug("Send data to cid [%d]! cmd:0x%04X sid:%d flag:%d chksum:0x%08X",
		session.cid, head.GetCmd(), head.GetSid(), head.GetFlag(), head.GetChkSum())

	return 0
}

/******************************************************************************
 **函数名称: lsnd_conn_destroy
 **功    能: 销毁CONN对象
 **输入参数:
 **     client: 客户端对象
 **     data: 收到的数据
 **     length: 收到数据的长度
 **输出参数:
 **返    回: 0:正常 !0:异常
 **实现描述:
 **注意事项: 返回!0值将导致连接断开
 **作    者: # Qifeng.zou # 2017.03.04 15:21:43 #
 ******************************************************************************/
func (ctx *LsndCntx) lsnd_conn_destroy(client *lws.Client, data []byte, length int) int {
	session, ok := client.GetUserData().(*LsndSessionExtra)
	if !ok {
		ctx.log.Error("Get connection extra data failed!")
		return -1
	}

	ctx.log.Debug("Destroy session object! cid:%d sid:%d", session.GetCid(), session.GetSid())

	return 0
}

/******************************************************************************
 **函数名称: LsndLwsCallBack
 **功    能: LWS处理回调
 **输入参数:
 **     ws: LWS上下文
 **     client: 客户端对象
 **     reason: 回调原因
 **     data: 收到的数据
 **     length: 数据长度
 **     param: 扩展数据
 **输出参数:
 **返    回: 0:正常 !0:异常
 **实现描述:
 **注意事项: 返回!0值将导致连接断开
 **作    者: # Qifeng.zou # 2017.03.04 00:16:09 #
 ******************************************************************************/
func LsndLwsCallBack(ws *lws.LwsCntx, client *lws.Client,
	reason int, data []byte, length int, param interface{}) int {
	ctx, ok := param.(*LsndCntx)
	if !ok {
		return -1
	}

	ctx.log.Error("LsndLwsCallBack() cid:%d reason:%d", client.GetCid(), reason)

	switch reason {
	case lws.LWS_CALLBACK_REASON_CREAT:
		return ctx.lsnd_conn_init(client)
	case lws.LWS_CALLBACK_REASON_RECV:
		return ctx.lsnd_conn_recv(client, data, length)
	case lws.LWS_CALLBACK_REASON_SEND:
		return ctx.lsnd_conn_send(client, data, length)
	case lws.LWS_CALLBACK_REASON_CLOSE:
		return ctx.lsnd_conn_destroy(client, data, length)
	default:
		ctx.log.Error("Call LsndLwsCallBack()! Unknown reason:%d", reason)
	}
	return 0
}