#if !defined(__CHAT_H__)
#define __CHAT_H__

#include "log.h"
#include "comm.h"
#include "list.h"
#include "rb_tree.h"
#include "avl_tree.h"
#include "hash_tab.h"

/* 订阅项 */
typedef struct
{
    uint16_t cmd;               /* 命令类型 */
} chat_sub_item_t;

typedef struct
{
    uint64_t rid;               /* 聊天室ID */
    uint32_t gid;               /* 分组ID */
} chat_session_room_info_t;

/* 聊天室会话信息 */
typedef struct
{
    uint64_t sid;               /* 会话ID(主键) */

    uint64_t rid;               /* 聊天室ID */
    uint32_t gid;               /* 分组ID */

    hash_tab_t *sub;            /* 订阅消息列表 */
} chat_session_t;

/* 聊天室分组信息 */
typedef struct
{
    uint16_t gid;               /* 聊天室分组ID */

    uint64_t sid_num;           /* SID总数 */
    hash_tab_t *sid_set;        /* 聊天室分组中SID列表
                                   (以SID为主键, 存储的也是SID值) */
} chat_group_t;

/* 聊天室信息 */
typedef struct
{
    uint64_t rid;               /* ROOMID(主键) */

    uint64_t sid_num;           /* SID总数 */
    uint32_t grp_num;           /* 分组总数 */

    time_t create_tm;           /* 创建时间 */

    hash_tab_t *group_tab;      /* 聊天室分组管理表(以gid为组建 存储chat_group_t数据) */
} chat_room_t;

/* 全局信息 */
typedef struct
{
    log_cycle_t *log;           /* 日志对象 */

    hash_tab_t *room_tab;       /* 聊天室列表(注: ROOMID为主键 存储数据chat_room_t) */
    hash_tab_t *session_tab;    /* SESSION信息(注: SID为主键存储数据chat_session_t)
                                   注意: 此处的存储数据对象在group->sid_list被引用,
                                   释放时千万不能释放多次 */
} chat_tab_t;

chat_tab_t *chat_tab_init(int len, log_cycle_t *log); // OK

uint32_t chat_room_add_session(chat_tab_t *chat, uint64_t rid, uint32_t gid, uint64_t sid); // OK
int chat_del_session(chat_tab_t *chat, uint64_t sid); // OK
int chat_timeout_hdl(chat_tab_t *chat);

int chat_add_sub(chat_tab_t *chat, uint64_t sid, uint16_t cmd); // OK
int chat_del_sub(chat_tab_t *chat, uint64_t sid, uint16_t cmd); // OK
bool chat_has_sub(chat_tab_t *chat, uint64_t sid, uint16_t cmd); // OK

int chat_room_trav(chat_tab_t *chat, uint64_t rid, uint16_t gid, trav_cb_t proc, void *args); // OK
int chat_timeout_clean_hdl(chat_tab_t *chat);

#endif /*__CHAT_H__*/
