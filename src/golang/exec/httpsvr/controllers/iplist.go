package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/astaxie/beego"

	"chat/src/golang/lib/comm"
	"chat/src/golang/lib/crypt"
)

/* 获取IP列表 */
type HttpSvrIpListCtrl struct {
	BaseController
}

/* 注册参数 */
type HttpSvrIpListParam struct {
	uid      uint64 // 用户ID
	sid      uint64 // 会话ID
	clientip string // 客户端IP
}

func (this *HttpSvrIpListCtrl) IpList() {
	ctx := GetHttpCtx()

	/* > 提取注册参数 */
	param, err := this.parse_param(ctx)
	if nil != err {
		ctx.log.Error("Parse param failed! uid:%d sid:%d clientip:%s",
			param.uid, param.sid, param.clientip)
		this.response_fail(param, comm.ERR_SVR_PARSE_PARAM, err.Error())
		return
	}

	ctx.log.Debug("Param list. uid:%d sid:%d clientip:%s", param.uid, param.sid, param.clientip)

	/* > 获取IP列表 */
	this.handler(ctx, param)

	return
}

/******************************************************************************
 **函数名称: parse_param
 **功    能: 解析参数
 **输入参数:
 **     ctx: 上下文
 **输出参数: NONE
 **返    回:
 **     param: 注册参数
 **     err: 错误描述
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.25 23:17:52 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) parse_param(ctx *HttpSvrCntx) (*HttpSvrIpListParam, error) {
	var param *HttpSvrIpListParam

	/* > 提取注册参数 */
	id, _ := this.GetInt64("uid")
	param.uid = uint64(id)

	id, _ = this.GetInt64("sid")
	param.sid = uint64(id)

	param.clientip = this.GetString("client")

	/* > 校验参数合法性 */
	if 0 == param.uid || 0 == param.sid || "" == param.clientip {
		ctx.log.Error("Paramter is invalid. uid:%d sid:%d clientip:%s",
			param.uid, param.sid, param.clientip)
		return nil, errors.New("Paramter is invalid!")
	}

	return param, nil
}

/* IP列表应答 */
type HttpSvrIpListRsp struct {
	Uid    uint64   `json:"uid"`    // 用户ID
	Sid    uint64   `json:"sid"`    // 会话ID
	Token  string   `json:"token"`  // 鉴权TOKEN
	Len    int      `json:"len"`    // 列表长度
	List   []string `json:"list"`   // IP列表
	Errno  int      `json:"errno"`  // 错误码
	ErrMsg string   `json:"errmsg"` // 错误描述
}

/******************************************************************************
 **函数名称: handler
 **功    能: 获取IP列表
 **输入参数:
 **     ctx: 上下文
 **     param: URL参数
 **输出参数:
 **返    回: NONE
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.24 17:00:07 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) handler(ctx *HttpSvrCntx, param *HttpSvrIpListParam) {
	iplist := this.get_iplist(ctx, param.clientip)
	if nil == iplist {
		ctx.log.Error("Get ip list failed!")
		this.response_fail(param, comm.ERR_SYS_SYSTEM, "Get ip list failed!")
		return
	}

	this.response_success(param, iplist)

	return
}

/******************************************************************************
 **函数名称: response_fail
 **功    能: 应答错误信息
 **输入参数:
 **     param: 注册参数
 **     errno: 错误码
 **     errmsg: 错误描述
 **输出参数:
 **返    回: NONE
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.25 23:13:09 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) response_fail(param *HttpSvrIpListParam, errno int, errmsg string) {
	var resp HttpSvrIpListRsp

	resp.Uid = param.uid
	resp.Sid = 0
	resp.Errno = errno
	resp.ErrMsg = errmsg

	this.Data["json"] = &resp
	this.ServeJSON()
}

/******************************************************************************
 **函数名称: response_success
 **功    能: 应答处理成功
 **输入参数:
 **     param: 注册参数
 **     iplist: IP列表
 **输出参数:
 **返    回: NONE
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.25 23:13:02 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) response_success(param *HttpSvrIpListParam, iplist []string) {
	var resp HttpSvrIpListRsp

	resp.Uid = param.uid
	resp.Sid = param.sid
	resp.Token = this.gen_token(param)
	resp.Len = len(iplist)
	resp.List = iplist
	resp.Errno = 0
	resp.ErrMsg = "OK"

	this.Data["json"] = &resp
	this.ServeJSON()
}

/******************************************************************************
 **函数名称: gen_token
 **功    能: 生成TOKEN字串
 **输入参数:
 **输出参数: NONE
 **返    回: 加密TOKEN
 **实现描述:
 **注意事项:
 **     TOKEN的格式"${uid}:${ttl}:${sid}"
 **     uid: 用户ID
 **     ttl: 该token的最大生命时间
 **     sid: 会话UID
 **作    者: # Qifeng.zou # 2016.11.25 23:54:27 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) gen_token(param *HttpSvrIpListParam) string {
	ctx := GetHttpCtx()
	ttl := time.Now().Unix() + 3*comm.TIME_DAY

	/* > 原始TOKEN */
	token := fmt.Sprintf("%d:%d:%d", param.uid, ttl, param.sid)

	/* > 加密TOKEN */
	cry := crypt.CreateEncodeCtx(ctx.conf.Cipher)

	return crypt.Encode(cry, token)
}

/******************************************************************************
 **函数名称: get_iplist
 **功    能: 获取IP列表
 **输入参数:
 **     ctx: 上下文
 **     clientip: 客户端IP
 **输出参数: NONE
 **返    回: IP列表
 **实现描述:
 **注意事项: 加读锁
 **作    者: # Qifeng.zou # 2016.11.27 07:42:54 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) get_iplist(ctx *HttpSvrCntx, clientip string) []string {
	item := ctx.ipdict.Query(clientip)
	if nil == item {
		ctx.lsnlist.RLock()
		defer ctx.lsnlist.RUnlock()
		return this.def_get_iplist(ctx)
	}

	ctx.lsnlist.RLock()
	defer ctx.lsnlist.RUnlock()

	/* > 获取国家/地区下辖的运营商列表 */
	operators, ok := ctx.lsnlist.list[item.GetNation()]
	if nil == operators || !ok {
		return this.def_get_iplist(ctx)
	}

	/* > 获取运营商下辖的侦听层列表 */
	lsn_list, ok := operators[item.GetOperator()]
	if nil == lsn_list || !ok {
		return this.def_get_iplist(ctx)
	}

	items := make([]string, 0)
	items = append(items, lsn_list[rand.Intn(len(lsn_list))])
	return items
}

/******************************************************************************
 **函数名称: def_get_iplist
 **功    能: 获取默认IP列表
 **输入参数:
 **     ctx: 上下文
 **     clientip: 客户端IP
 **输出参数: NONE
 **返    回: IP列表
 **实现描述:
 **注意事项: 外部已经加读锁
 **作    者: # Qifeng.zou # 2016.11.27 19:33:49 #
 ******************************************************************************/
func (this *HttpSvrIpListCtrl) def_get_iplist(ctx *HttpSvrCntx) []string {
	var ok bool

	/* > 获取"默认"国家/地区下辖的运营商列表 */
	operators, ok := ctx.lsnlist.list["DEF"]
	if nil == operators || !ok {
		return nil
	}

	/* > 获取"默认"运营商下辖的侦听层列表 */
	lsn_list, ok := operators["DEF"]
	if nil == lsn_list || 0 == len(lsn_list) || !ok {
		return nil
	}

	items := make([]string, 0)
	items = append(items, lsn_list[rand.Intn(len(lsn_list))])
	return items
}
