package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"beehive-im/src/golang/exec/listend-ws/controllers"
	"beehive-im/src/golang/exec/listend-ws/controllers/conf"
)

func main() {
	var conf conf.LsndConf

	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	/* > 加载LSND-WS配置 */
	if err := conf.LoadConf(); nil != err {
		fmt.Printf("Load configuration failed! errmsg:%s\n", err.Error())
		return
	}

	/* > 初始化LSND-WS环境 */
	ctx, err := controllers.LsndInit(&conf)
	if nil != err {
		fmt.Printf("Initialize context failed! errmsg:%s\n", err.Error())
		return
	}

	/* > 注册回调函数 */
	ctx.Register()

	/* > 启动侦听服务 */
	ctx.Launch()

	/* > 捕捉中断信号 */
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch

	return
}
