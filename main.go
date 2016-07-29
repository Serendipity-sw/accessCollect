package main

import (
	"github.com/guotie/config"
	"fmt"
	"flag"
	"strings"
	"github.com/guotie/deferinit"
	"runtime"
	"github.com/smtc/glog"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

var (
	authorize string //授权码
	urls interface{} //需要请求的地址
	configFn                                    = flag.String("config", "./config.json", "config file path")
	debugFlag                                   = flag.Bool("d", false, "debug mode")
	rootPrefix                           string
	rt                                   *gin.Engine
)

/**
服务启动
创建人:邵炜
创建时间:2016年7月29日15:49:56
 */
func serverRun(cfn string, debug bool) {
	config.ReadCfg(cfn)
	logInit(debug)
	authorize=config.GetString("authorize")
	urls=config.Get("urls")
	rootPrefix=config.GetString("rootPrefix")
	port:=config.GetIntDefault("port",8000)

	if len(rootPrefix) != 0 {
		if !strings.HasPrefix(rootPrefix, "/") {
			rootPrefix = "/" + rootPrefix
		}
		if strings.HasSuffix(rootPrefix, "/") {
			rootPrefix = rootPrefix[0 : len(rootPrefix)-1]
		}
	}
	deferinit.InitAll()
	glog.Info("init all module successfully.\n")

	// 设置多cpu运行
	runtime.GOMAXPROCS(runtime.NumCPU())

	deferinit.RunRoutines()
	glog.Info("run routines successfully.\n")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	rt = gin.Default()
	router(rt)
	go rt.Run(fmt.Sprintf(":%d", port))
}

func serverExit() {
	// 结束所有go routine
	deferinit.StopRoutines()
	glog.Info("stop routine successfully.\n")

	deferinit.FiniAll()
	glog.Info("fini all modules successfully.\n")
}

func main() {
	if checkPid() {
		return
	}

	flag.Parse()

	serverRun(*configFn, *debugFlag)

	c := make(chan os.Signal, 1)
	writePid()
	// 信号处理
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	// 等待信号
	<-c
	serverExit()
	rmPidFile()
	glog.Close()
	os.Exit(0)
}

func router(r *gin.Engine) {
	g := &r.RouterGroup
	if rootPrefix != "" {
		g = r.Group(rootPrefix)
	}
	{
		g.GET("/selectAccessFun",selectAccessFun)
	}
}
