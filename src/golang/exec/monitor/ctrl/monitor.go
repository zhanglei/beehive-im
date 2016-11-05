package ctrl

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"

	"chat/src/golang/lib/comm"
	"chat/src/golang/lib/log"
	"chat/src/golang/lib/rtmq"
)

/* OLS上下文 */
type MonCntx struct {
	conf  *MonConf            /* 配置信息 */
	log   *logs.BeeLogger     /* 日志对象 */
	proxy *rtmq.RtmqProxyCntx /* 代理对象 */
	redis *redis.Pool         /* REDIS连接池 */
}

/******************************************************************************
 **函数名称: MonInit
 **功    能: 初始化对象
 **输入参数:
 **     conf: 配置信息
 **输出参数: NONE
 **返    回:
 **     ctx: 上下文
 **     err: 错误信息
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.10.30 22:32:23 #
 ******************************************************************************/
func MonInit(conf *MonConf) (ctx *MonCntx, err error) {
	ctx = &MonCntx{}

	ctx.conf = conf

	/* > 初始化日志 */
	ctx.log = log.Init(conf.Log.Level, conf.Log.Path, "monitor.log")
	if nil == ctx.log {
		return nil, errors.New("Initialize log failed!")
	}

	/* > REDIS连接池 */
	ctx.redis = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.RedisAddr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	if nil == ctx.redis {
		ctx.log.Error("Create redis pool failed! addr:%s", conf.RedisAddr)
		return nil, errors.New("Create redis pool failed!")
	}

	/* > 初始化RTMQ-PROXY */
	ctx.proxy = rtmq.ProxyInit(&conf.proxy, ctx.log)
	if nil == ctx.proxy {
		return nil, err
	}

	return ctx, nil
}

/******************************************************************************
 **函数名称: Register
 **功    能: 注册处理回调
 **输入参数: NONE
 **输出参数: NONE
 **返    回: VOID
 **实现描述: 注册回调函数
 **注意事项: 请在调用Launch()前完成此函数调用
 **作    者: # Qifeng.zou # 2016.10.30 22:32:23 #
 ******************************************************************************/
func (ctx *MonCntx) Register() {
	/* > 运维消息 */
	ctx.proxy.Register(comm.CMD_LSN_RPT, MonLsnRptHandler, ctx)
	ctx.proxy.Register(comm.CMD_FRWD_LIST, MonFrwdRptHandler, ctx)
}

/******************************************************************************
 **函数名称: Launch
 **功    能: 启动OLSVR服务
 **输入参数: NONE
 **输出参数: NONE
 **返    回: VOID
 **实现描述:
 **注意事项:
 **作    者: # Qifeng.zou # 2016.10.30 22:32:23 #
 ******************************************************************************/
func (ctx *MonCntx) Launch() {
	ctx.proxy.Launch()
}
