package comm

const (
	CHAT_SID_TTL            = 300   // 会话SID-TTL
	CHAT_OP_TTL             = 30    // 运营商ID-TTL
	CHAT_NID_TTL            = 30    // 结点ID-TTL
	CHAT_BAT_NUM            = 1000  // 批量操作个数
	CHAT_ROOM_GROUP_MAX_NUM = 10000 // 各组最大人数
)

/* 时间转换成秒 */
const (
	TIME_MIN  = 60             // 分
	TIME_HOUR = 3600           // 时
	TIME_DAY  = 86400          // 天
	TIME_WEEK = 7 * 86400      // 周
	TIME_YEAR = 365 * TIME_DAY // 年
)

/* 加锁方式 */
const (
	NOLOCK = 0 // 不加锁
	RDLOCK = 1 // 加读锁
	WRLOCK = 2 // 加写锁
)

/* 网络类型 */
const (
	LSND_NET_UNKNOWN = 0 // 未知网络
	LSND_NET_TCP     = 1 // 网络类型: TCP
	LSND_NET_WS      = 2 // 网络类型: WebSocket
)
