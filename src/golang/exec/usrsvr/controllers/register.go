package controllers

import (
	"errors"

	"beehive-im/src/golang/lib/comm"
	"beehive-im/src/golang/lib/mesg/seqsvr"
)

/* 注册处理 */
type UsrSvrRegisterCtrl struct {
	BaseController
}

/* 注册参数 */
type UsrSvrRegisterParam struct {
	uid    uint64 // 用户ID
	nation uint64 // 国际ID
	city   uint64 // 地市ID
	town   uint64 // 县城ID
}

func (this *UsrSvrRegisterCtrl) Register() {
	ctx := GetUsrSvrCtx()

	/* > 提取参数 */
	param, err := this.register_parse_param(ctx)
	if nil != err {
		ctx.log.Error("Parse register failed! uid:%d nation:%d city:%d town:%d",
			param.uid, param.nation, param.city, param.town)
		this.Error(comm.ERR_SVR_PARSE_PARAM, err.Error())
		return
	}

	ctx.log.Debug("Register param list. uid:%d nation:%d city:%d town:%d",
		param.uid, param.nation, param.city, param.town)

	this.register_handler(param)

	return
}

/******************************************************************************
 **函数名称: register_parse_param
 **功    能: 解析参数
 **输入参数:
 **     ctx: 全局对象
 **输出参数: NONE
 **返    回:
 **     param: 注册参数
 **     err: 错误描述
 **实现描述: 从URL中抽取参数字段
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.25 10:30:09 #
 ******************************************************************************/
func (this *UsrSvrRegisterCtrl) register_parse_param(ctx *UsrSvrCntx) (*UsrSvrRegisterParam, error) {
	param := &UsrSvrRegisterParam{}

	/* > 提取注册参数 */
	id, _ := this.GetInt64("uid")
	param.uid = uint64(id)

	id, _ = this.GetInt64("nation")
	param.nation = uint64(id)

	id, _ = this.GetInt64("city")
	param.city = uint64(id)

	id, _ = this.GetInt64("town")
	param.town = uint64(id)

	/* > 校验参数合法性 */
	if 0 == param.uid || 0 == param.nation {
		ctx.log.Error("Register param invalid! uid:%d nation:%d", param.uid, param.nation)
		return param, errors.New("Register param invalid!")
	}

	return param, nil
}

/* 注册应答 */
type UsrSvrRegisterRsp struct {
	Uid    uint64 `json:"uid"`    // 用户ID
	Sid    uint64 `json:"sid"`    // 会话ID
	Nation uint64 `json:"nation"` // 国家ID(国)
	City   uint64 `json:"city"`   // 城市ID(市)
	Town   uint64 `json:"town"`   // 城镇ID(县)
	Code   int    `json:"code"`   // 错误码
	ErrMsg string `json:"errmsg"` // 错误描述
}

/******************************************************************************
 **函数名称: register_handler
 **功    能: 注册处理
 **输入参数:
 **     param: 注册参数
 **输出参数: NONE
 **返    回: VOID
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.24 17:34:27 #
 ******************************************************************************/
func (this *UsrSvrRegisterCtrl) register_handler(param *UsrSvrRegisterParam) {
	ctx := GetUsrSvrCtx()

	/* > 取SEQSVR连接 */
	conn, err := ctx.seqsvr_pool.Get()
	if nil != err {
		ctx.log.Error("Get seqsvr connection pool failed! errmsg:%s", err.Error())
		this.Error(comm.ERR_SYS_RPC, err.Error())
		return
	}
	client := conn.(*seqsvr.SeqSvrThriftClient)
	defer ctx.seqsvr_pool.Put(client, false)

	/* > 申请会话ID */
	sid, err := client.AllocSid()
	if nil != err {
		ctx.log.Error("Alloc sid failed! errmsg:%s", err.Error())
		this.Error(comm.ERR_SYS_RPC, err.Error())
		return
	}

	ctx.log.Debug("Alloc sid success! uid:%d sid:%d", param.uid, sid)

	this.success(param, uint64(sid))

	return
}

/******************************************************************************
 **函数名称: success
 **功    能: 应答处理成功
 **输入参数:
 **     param: 注册参数
 **     sid: 会话SID
 **输出参数:
 **返    回: NONE
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.24 19:13:22 #
 ******************************************************************************/
func (this *UsrSvrRegisterCtrl) success(param *UsrSvrRegisterParam, sid uint64) {
	var resp UsrSvrRegisterRsp

	resp.Uid = param.uid
	resp.Sid = sid
	resp.Nation = param.nation
	resp.City = param.city
	resp.Town = param.town
	resp.Code = 0
	resp.ErrMsg = "OK"

	this.Data["json"] = &resp
	this.ServeJSON()
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
