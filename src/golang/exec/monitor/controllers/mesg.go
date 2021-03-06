package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"

	"beehive-im/src/golang/lib/comm"
	"beehive-im/src/golang/lib/mesg"
)

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

/******************************************************************************
 **函数名称: lsnd_info_isvalid
 **功    能: 判断LSN-RPT是否合法
 **输入参数:
 **     req: HB请求
 **输出参数: NONE
 **返    回: true:合法 false:非法
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 08:38:48 #
 ******************************************************************************/
func (ctx *MonSvrCntx) lsnd_info_isvalid(req *mesg.MesgLsndInfo) bool {
	if 0 == req.GetNid() ||
		0 == req.GetPort() ||
		0 == len(req.GetNation()) ||
		0 == len(req.GetIp()) {
		return false
	}
	return true
}

/******************************************************************************
 **函数名称: lsnd_info_parse
 **功    能: 解析LSND-INFO请求
 **输入参数:
 **     data: 接收的数据
 **输出参数: NONE
 **返    回:
 **     head: 通用协议头
 **     req: 协议体内容
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 08:14:10 #
 ******************************************************************************/
func (ctx *MonSvrCntx) lsnd_info_parse(data []byte) (
	head *comm.MesgHeader, req *mesg.MesgLsndInfo) {
	/* > 字节序转换 */
	head = comm.MesgHeadNtoh(data)
	if !head.IsValid(0) {
		ctx.log.Error("Header is invalid! cmd:0x%04X nid:%d",
			head.GetCmd(), head.GetNid())
		return nil, nil
	}

	/* > 解析PB协议 */
	req = &mesg.MesgLsndInfo{}
	err := proto.Unmarshal(data[comm.MESG_HEAD_SIZE:], req)
	if nil != err {
		ctx.log.Error("Unmarshal listend information failed! errmsg:%s", err.Error())
		return nil, nil
	}

	/* > 校验协议合法性 */
	if !ctx.lsnd_info_isvalid(req) {
		return nil, nil
	}

	return head, req
}

/******************************************************************************
 **函数名称: lsnd_info_has_conflict
 **功    能: 判断数据是否冲突
 **输入参数:
 **     req: 帧听层上报消息
 **输出参数: NONE
 **返    回: true:存在冲突 false:不存在冲突
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 10:34:00 #
 ******************************************************************************/
func (ctx *MonSvrCntx) lsnd_info_has_conflict(req *mesg.MesgLsndInfo) (has bool, err error) {
	rds := ctx.redis.Get()
	defer rds.Close()

	addr := fmt.Sprintf(comm.IM_FMT_IP_PORT_STR, req.GetIp(), req.GetPort())
	ok, err := redis.Bool(rds.Do("HEXISTS", comm.IM_KEY_LSND_ADDR_TO_NID, addr))
	if nil != err {
		ctx.log.Error("Exec hexists failed! err:%s", err.Error())
		return false, err
	} else if true == ok {
		nid, err := redis.Int(rds.Do("HGET", comm.IM_KEY_LSND_ADDR_TO_NID, addr))
		if nil != err {
			ctx.log.Error("Exec hget failed! err:%s", err.Error())
			return false, err
		} else if uint32(nid) != req.GetNid() {
			ctx.log.Error("Node id conflict! nid:%d/%d", nid, req.GetNid())
			return true, nil
		}
	}

	key := fmt.Sprintf(comm.IM_KEY_LSND_ATTR, req.GetNid())
	ok, err = redis.Bool(rds.Do("HEXISTS", key, comm.IM_LSND_ATTR_ADDR))
	if nil != err {
		ctx.log.Error("Exec hexists failed! err:%s", err.Error())
		return
	} else if true == ok {
		_addr, err := redis.String(rds.Do("HGET", key, comm.IM_LSND_ATTR_ADDR))
		if nil != err {
			ctx.log.Error("Exec hget failed! err:%s", err.Error())
			return false, err
		} else if _addr != addr {
			ctx.log.Error("Node id conflict! addr:%s/%s", addr, _addr)
			return true, nil
		}
	}

	return false, nil
}

/******************************************************************************
 **函数名称: lsnd_info_handler
 **功    能: LSND-INFO处理
 **输入参数:
 **     head: 协议头
 **     req: 上线请求
 **输出参数: NONE
 **返    回: 异常信息
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 08:41:18 #
 ******************************************************************************/
func (ctx *MonSvrCntx) lsnd_info_handler(head *comm.MesgHeader, req *mesg.MesgLsndInfo) {
	pl := ctx.redis.Get()
	defer func() {
		pl.Do("")
		pl.Close()
	}()

	/* > 判断数据是否冲突 */
	has, err := ctx.lsnd_info_has_conflict(req)
	if nil != err {
		ctx.log.Error("Something was wrong! errmsg:%s!", err.Error())
		return
	} else if true == has {
		ctx.log.Error("Data has conflict! type:%d nid:%d nation:%s opid:%d ip:%s port:%d",
			req.GetType(), req.GetNid(), req.GetNation(), req.GetOpid(), req.GetIp(), req.GetPort())
		return
	}

	ttl := time.Now().Unix() + comm.CHAT_OP_TTL

	/* 存储基本信息 */
	key := fmt.Sprintf(comm.IM_KEY_LSND_ATTR, req.GetNid())
	addr := fmt.Sprintf(comm.IM_FMT_IP_PORT_STR, req.GetIp(), req.GetPort())

	pl.Send("HSETNX", key, comm.IM_LSND_ATTR_ADDR, addr)                     /* 记录NID->ADDR映射 */
	pl.Send("HSET", key, comm.IM_LSND_ATTR_TYPE, req.GetType())              /* 侦听层类型 */
	pl.Send("HSET", key, comm.IM_LSND_ATTR_CONNECTION, req.GetConnections()) /* 记录NID在线连接数 */

	pl.Send("HSETNX", comm.IM_KEY_LSND_ADDR_TO_NID, addr, req.GetNid()) /* 记录ADDR->NID映射 */

	/* 侦听层ID集合 */
	pl.Send("ZADD", comm.IM_KEY_LSND_NID_ZSET, ttl, req.GetNid())

	/* 侦听层类型集合 */
	pl.Send("ZADD", comm.IM_KEY_LSND_TYPE_ZSET, ttl, req.GetType())

	/* 国家集合 */
	key = fmt.Sprintf(comm.IM_KEY_LSND_NATION_ZSET, req.GetType())
	pl.Send("ZADD", key, ttl, req.GetNation())

	/* 国家 -> 运营商列表 */
	key = fmt.Sprintf(comm.IM_KEY_LSND_OP_ZSET, req.GetType(), req.GetNation())
	pl.Send("ZADD", key, ttl, req.GetOpid())

	/* 国家+运营商 -> 结点列表 */
	key = fmt.Sprintf(comm.IM_KEY_LSND_OP_TO_NID_ZSET, req.GetType(), req.GetNation(), req.GetOpid())
	pl.Send("ZADD", key, ttl, req.GetNid())

	/* 国家+运营商 -> 侦听层IP列表 */
	key = fmt.Sprintf(comm.IM_KEY_LSND_IP_ZSET, req.GetType(), req.GetNation(), req.GetOpid())
	val := fmt.Sprintf(comm.IM_FMT_IP_PORT_STR, req.GetIp(), req.GetPort())
	pl.Send("ZADD", key, ttl, val)

	ctx.log.Debug("Handle listend information! type:%d nid:%d nation:%s opid:%d ip:%s port:%d user-num:%d",
		req.GetType(), req.GetNid(), req.GetNation(), req.GetOpid(), req.GetIp(), req.GetPort(), req.GetConnections())

	return
}

/******************************************************************************
 **函数名称: MonLsndInfoHandler
 **功    能: 帧听层上报
 **输入参数:
 **     cmd: 消息类型
 **     nid: 结点ID
 **     data: 收到数据
 **     length: 数据长度
 **     param: 附加参数
 **输出参数: NONE
 **返    回: VOID
 **实现描述:
 **协议格式:
 **     {
 **        required uint64 nid = 1;    // M|结点ID|数字|<br>
 **        required string nation = 2; // M|所属国家|字串|<br>
 **        required string name = 3;   // M|运营商名称|字串|<br>
 **        required string ipaddr = 4; // M|IP地址|字串|<br>
 **        required uint32 port = 5;   // M|端口号|数字|<br>
 **     }
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 06:32:03 #
 ******************************************************************************/
func MonLsndInfoHandler(cmd uint32, nid uint32, data []byte, length uint32, param interface{}) int {
	ctx, ok := param.(*MonSvrCntx)
	if !ok {
		return -1
	}

	ctx.log.Debug("Recv listend information!")

	/* 1. > 解析LSN-RPT请求 */
	head, req := ctx.lsnd_info_parse(data)
	if nil == head || nil == req {
		ctx.log.Error("Parse listend information failed!")
		return -1
	}

	/* 2. > LSND-INFO请求处理 */
	ctx.lsnd_info_handler(head, req)

	return 0
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

/******************************************************************************
 **函数名称: frwd_info_isvalid
 **功    能: 判断LSN-RPT是否合法
 **输入参数:
 **     req: HB请求
 **输出参数: NONE
 **返    回: true:合法 false:非法
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 11:05:03 #
 ******************************************************************************/
func (ctx *MonSvrCntx) frwd_info_isvalid(req *mesg.MesgFrwdInfo) bool {
	if 0 == req.GetForwardPort() ||
		0 == req.GetBackendPort() ||
		0 == len(req.GetIp()) {
		return false
	}
	return true
}

/******************************************************************************
 **函数名称: frwd_info_parse
 **功    能: 解析LSN-PRT请求
 **输入参数:
 **     data: 接收的数据
 **输出参数: NONE
 **返    回:
 **     head: 通用协议头
 **     req: 协议体内容
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 11:04:57 #
 ******************************************************************************/
func (ctx *MonSvrCntx) frwd_info_parse(data []byte) (
	head *comm.MesgHeader, req *mesg.MesgFrwdInfo) {
	/* > 字节序转换 */
	head = comm.MesgHeadNtoh(data)
	if !head.IsValid(0) {
		ctx.log.Error("Header is invalid! cmd:0x%04X nid:%d",
			head.GetCmd(), head.GetNid())
		return nil, nil
	}

	/* > 解析PB协议 */
	req = &mesg.MesgFrwdInfo{}
	err := proto.Unmarshal(data[comm.MESG_HEAD_SIZE:], req)
	if nil != err {
		ctx.log.Error("Unmarshal lsn-rpt failed! errmsg:%s", err.Error())
		return nil, nil
	}

	/* > 校验协议合法性 */
	if !ctx.frwd_info_isvalid(req) {
		return nil, nil
	}

	return head, req
}

/******************************************************************************
 **函数名称: frwd_info_has_conflict
 **功    能: 判断数据是否冲突
 **输入参数:
 **     req: 帧听层上报消息
 **输出参数: NONE
 **返    回: true:存在冲突 false:不存在冲突
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 11:04:49 #
 ******************************************************************************/
func (ctx *MonSvrCntx) frwd_info_has_conflict(req *mesg.MesgFrwdInfo) (has bool, err error) {
	rds := ctx.redis.Get()
	defer rds.Close()

	key := fmt.Sprintf(comm.IM_KEY_FRWD_ATTR, req.GetNid())

	vals, err := redis.Strings(rds.Do("HMGET", key,
		comm.IM_FRWD_ATTR_ADDR, comm.IM_FRWD_ATTR_BC_PORT, comm.IM_FRWD_ATTR_FWD_PORT))
	if nil != err {
		return false, nil
	}

	addr := vals[0]
	if addr != req.GetIp() {
		return true, errors.New(fmt.Sprintf("Address is conflict! nid:%d", req.GetNid()))
	}

	port, err := strconv.ParseInt(vals[1], 10, 32)
	if uint32(port) != req.GetBackendPort() {
		return true, errors.New(fmt.Sprintf("Backend port is conflict! nid:%d port:%d/%d",
			req.GetNid(), port, req.GetBackendPort()))
	}

	port, err = strconv.ParseInt(vals[2], 10, 32)
	if uint32(port) != req.GetForwardPort() {
		return true, errors.New(fmt.Sprintf("Forward port is conflict! nid:%d port:%d/%d",
			req.GetNid(), port, req.GetForwardPort()))
	}

	return false, nil
}

/******************************************************************************
 **函数名称: frwd_info_handler
 **功    能: FRWD-RPT处理
 **输入参数:
 **     head: 协议头
 **     req: 上线请求
 **输出参数: NONE
 **返    回: 异常信息
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 11:04:42 #
 ******************************************************************************/
func (ctx *MonSvrCntx) frwd_info_handler(head *comm.MesgHeader, req *mesg.MesgFrwdInfo) {
	pl := ctx.redis.Get()
	defer func() {
		pl.Do("")
		pl.Close()
	}()

	/* > 判断数据是否冲突 */
	has, err := ctx.frwd_info_has_conflict(req)
	if nil != err {
		ctx.log.Error("Something was wrong! errmsg:%s", err.Error())
		return
	} else if true == has {
		ctx.log.Error("Data has conflict!")
		return
	}

	ttl := time.Now().Unix() + comm.CHAT_NID_TTL

	/* > 更新数据存储 */
	key := fmt.Sprintf(comm.IM_KEY_FRWD_ATTR, req.GetNid())

	pl.Send("HSETNX", key, comm.IM_FRWD_ATTR_ADDR, req.GetIp())
	pl.Send("HSETNX", key, comm.IM_FRWD_ATTR_BC_PORT, req.GetBackendPort())
	pl.Send("HSETNX", key, comm.IM_FRWD_ATTR_FWD_PORT, req.GetForwardPort())

	pl.Send("ZADD", comm.IM_KEY_FRWD_NID_ZSET, ttl, req.GetNid())

	return
}

/******************************************************************************
 **函数名称: MonFrwdInfoHandler
 **功    能: 转发层上报
 **输入参数:
 **     cmd: 消息类型
 **     nid: 结点ID
 **     data: 收到数据
 **     length: 数据长度
 **     param: 附加参数
 **输出参数: NONE
 **返    回: VOID
 **实现描述:
 **协议格式:
 **     {
 **         required uint64 nid = 1;        // M|结点ID|数字|<br>
 **         required string ipaddr = 2;     // M|IP地址|字串|<br>
 **         required uint32 forward_port = 3;    // M|前端口号|数字|<br>
 **         required uint32 backend_port = 4;    // M|后端口号|数字|<br>
 **     }
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.04 11:04:36 #
 ******************************************************************************/
func MonFrwdInfoHandler(cmd uint32, nid uint32, data []byte, length uint32, param interface{}) int {
	ctx, ok := param.(*MonSvrCntx)
	if !ok {
		return -1
	}

	ctx.log.Debug("Recv frwd-info request!")

	/* 1. > 解析FRWD-RPT请求 */
	head, req := ctx.frwd_info_parse(data)
	if nil == head || nil == req {
		ctx.log.Error("Parse frwd-info failed!")
		return -1
	}

	/* 2. > LSN-RPT请求处理 */
	ctx.frwd_info_handler(head, req)

	return 0
}
