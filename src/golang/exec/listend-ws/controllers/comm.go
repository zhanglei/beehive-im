package controllers

import (
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 处理回调的管理

/******************************************************************************
 **函数名称: Register
 **功    能: 注册处理回调
 **输入参数:
 **     cmd: 消息类型
 **     cb: 处理回调
 **     param: 附加数据
 **输出参数: NONE
 **返    回: 0:成功 !0:失败
 **实现描述:
 **注意事项: 使用读锁
 **作    者: # Qifeng.zou # 2017.02.09 23:50:28 #
 ******************************************************************************/
func (tab *MesgCallBackTab) Register(cmd uint32, cb MesgCallBack, param interface{}) int {
	item := &MesgCallBackItem{
		cmd:   cmd,
		proc:  cb,
		param: param,
	}

	tab.Lock()
	tab.list[cmd] = item
	tab.Unlock()

	return 0
}

/******************************************************************************
 **函数名称: Query
 **功    能: 查找处理回调
 **输入参数:
 **     cmd: 消息类型
 **输出参数: NONE
 **返    回:
 **     cb: 回调函数
 **     param: 附加数据
 **实现描述:
 **注意事项: 使用读锁
 **作    者: # Qifeng.zou # 2017.02.09 23:50:28 #
 ******************************************************************************/
func (tab *MesgCallBackTab) Query(cmd uint32) (cb MesgCallBack, param interface{}) {
	tab.RLock()
	item, ok := tab.list[cmd]
	if !ok {
		tab.RUnlock()
		return nil, nil
	}
	cb = item.proc
	param = item.param
	tab.RUnlock()

	return cb, param
}

////////////////////////////////////////////////////////////////////////////////
// 会话SID <-> 连接CID 映射操作

/* 添加会话SID->连接CID映射 */
func (ctx *LsndCntx) sid_to_cid_init() int {
	for idx := 0; idx < LSND_SID2CID_LEN; idx += 1 {
		tab := &ctx.sid2cid.tab[idx%LSND_SID2CID_LEN]
		tab.list = make(map[uint64]uint64)
	}
	return 0
}

/* 添加会话SID->连接CID映射 */
func (ctx *LsndCntx) sid_to_cid_add(sid uint64, cid uint64) int {
	tab := &ctx.sid2cid.tab[sid%LSND_SID2CID_LEN]
	tab.Lock()
	tab.list[sid] = cid
	tab.Unlock()
	return 0
}

/* 通过会话SID查找连接CID */
func (ctx *LsndCntx) find_cid_by_sid(sid uint64) uint64 {
	tab := &ctx.sid2cid.tab[sid%LSND_SID2CID_LEN]
	tab.RLock()
	cid, _ := tab.list[sid]
	tab.RUnlock()
	return cid
}

////////////////////////////////////////////////////////////////////////////////
// 连接扩展数据操作

/* 设置会话SID */
func (session *LsndSessionExtra) SetSid(sid uint64) {
	session.Lock()
	defer session.Unlock()

	session.sid = sid
}

/* 获取会话SID */
func (session *LsndSessionExtra) GetSid() uint64 {
	session.RLock()
	defer session.RUnlock()

	return session.sid
}

/* 获取连接CID */
func (session *LsndSessionExtra) SetCid(cid uint64) {
	session.Lock()
	defer session.Unlock()

	session.cid = cid
}

/* 获取连接CID */
func (session *LsndSessionExtra) GetCid() uint64 {
	session.RLock()
	defer session.RUnlock()

	return session.cid
}

/* 设置连接状态 */
func (session *LsndSessionExtra) SetStatus(status int) {
	session.Lock()
	defer session.Unlock()

	session.status = status
}

/* 获取连接状态 */
func (session *LsndSessionExtra) GetStatus() int {
	session.RLock()
	defer session.RUnlock()

	return session.status
}

/* 判断连接状态 */
func (session *LsndSessionExtra) IsStatus(status int) bool {
	session.RLock()
	defer session.RUnlock()

	if session.status == status {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////

/* 加入被踢列表 */
func (ctx *LsndCntx) kick_add(cid uint64) {
	item := &LsndKickItem{cid: cid, ttl: time.Now().Unix() + 5}

	ctx.kick_list <- item
}
