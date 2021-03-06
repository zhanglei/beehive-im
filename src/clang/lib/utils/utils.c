/******************************************************************************
 ** Coypright(C) 2014-2024 Qiware technology Co., Ltd
 **
 ** 文件名: utils.c
 ** 版本号: 1.0
 ** 描  述: 
 ** 作  者: # Qifeng.zou # 2016年06月15日 星期三 00时12分29秒 #
 ******************************************************************************/
#include "comm.h"

/******************************************************************************
 **函数名称: tlz_gen_serail
 **功    能: 生成系统流水号
 **输入参数:
 **     nid: 结点ID
 **     svrid: 服务ID
 **     seq: 序列号
 **输出参数:
 **返    回: 系统流水号
 **实现描述: 
 **注意事项: 
 **作    者: # Qifeng.zou # 2016.06.15 00:14:20 #
 ******************************************************************************/
uint64_t tlz_gen_serail(uint16_t nid, uint16_t svrid, uint32_t seq)
{
    serial_t s;

    s.nid = nid; // 结点ID
    s.svrid = svrid; // 服务索引(如: 第几个线程)
    s.seq = seq; // 序列号

    return s.serial;
}

/******************************************************************************
 **函数名称: tlz_gen_sid
 **功    能: 生成会话ID
 **输入参数:
 **     nid: 结点ID
 **     sid: 服务ID
 **     seq: 序列号
 **输出参数:
 **返    回: 系统流水号
 **实现描述: 
 **注意事项: 
 **作    者: # Qifeng.zou # 2016.06.15 00:14:20 #
 ******************************************************************************/
uint64_t tlz_gen_sid(uint16_t nid, uint16_t svrid, uint32_t seq)
{
    serial_t s;

    s.nid = nid; // 结点ID
    s.svrid = svrid; // 服务索引(如: 第几个线程)
    s.seq = seq; // 序列号

    return s.serial;
}
