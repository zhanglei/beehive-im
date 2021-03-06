package chat

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"

	"beehive-im/src/golang/lib/comm"
)

/* 群组角色 */
const (
	GROUP_ROLE_OWNER   = 1 // 群组-所有者
	GROUP_ROLE_MANAGER = 2 // 群组-管理员
)

/* 聊天室状态 */
const (
	GROUP_STAT_OPEN  = 1 // 群组-开启
	GROUP_STAT_CLOSE = 0 // 群组-关闭
)

/******************************************************************************
 **函数名称: GroupGetGidToNidSet
 **功    能: 通过群GID获取对应的帧听层NID列表
 **输入参数:
 **     pool: REDIS连接池
 **     gid: 群ID
 **输出参数: NONE
 **返    回: 帧听层ID列表
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.07 23:02:51 #
 ******************************************************************************/
func GroupGetGidToNidSet(pool *redis.Pool, gid uint64) (list []uint32, err error) {
	rds := pool.Get()
	defer rds.Close()

	ctm := time.Now().Unix()

	key := fmt.Sprintf(comm.CHAT_KEY_GID_TO_NID_ZSET, gid)
	nid_list, err := redis.Strings(rds.Do("ZRANGEBYSCORE", key, ctm, "+inf"))
	if nil != err {
		return nil, err
	}

	num := len(nid_list)
	list = make([]uint32, 0)
	for idx := 0; idx < num; idx += 1 {
		nid, _ := strconv.ParseInt(nid_list[idx], 10, 32)
		list = append(list, uint32(nid))
	}

	return list, nil
}

/******************************************************************************
 **函数名称: GroupGetGidToNidMap
 **功    能: 通过群GID获取对应的帧听层NID映射表
 **输入参数:
 **     pool: REDIS连接池
 **输出参数: NONE
 **返    回: gid->nid映射表
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.11.07 23:15:00 #
 ******************************************************************************/
func GroupGetGidToNidMap(pool *redis.Pool) (m map[uint64][]uint32, err error) {
	rds := pool.Get()
	defer rds.Close()

	ctm := time.Now().Unix()
	m = make(map[uint64][]uint32)

	off := 0
	key := fmt.Sprintf(comm.CHAT_KEY_GID_ZSET)
	for {
		gid_list, err := redis.Strings(rds.Do("ZRANGEBYSCORE",
			key, ctm, "+inf", "LIMIT", off, comm.CHAT_BAT_NUM))
		if nil != err {
			return nil, err
		}

		num := len(gid_list)
		for idx := 0; idx < num; idx += 1 {
			gid, _ := strconv.ParseInt(gid_list[idx], 10, 64)

			m[uint64(gid)], err = GroupGetGidToNidSet(pool, uint64(gid))
			if nil != err {
				return nil, err
			}
		}

		if num < comm.CHAT_BAT_NUM {
			break
		}
		off += num
	}

	return m, nil
}
