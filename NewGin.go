package go-new-gin-engine

import (
	"github.com/gin-gonic/gin"
	"arc-acl/src/com/yniii/config"
	"arc-acl/src/com/yniii/utils/Thread"
)

/**
* @Author: yNsLuHan
* @Description: 根据config获取gin
* @File: NewGin
* @Version: 1.0.0
* @Date: 2021/6/8 5:43 下午
 */
func NewGin() *gin.Engine {
	var engine *gin.Engine
	// 根据config判断是否为debug模式
	//var conf =
	if config.GetConf().GetBool("Server.debug") {
		engine = gin.Default()
	} else {
		// 生产模式，没有控制台日志
		gin.SetMode(gin.ReleaseMode)
		// 调整cpu个数
		Thread.ConfigRuntime()
		// 开启协程
		go Thread.StatsWorker()
		engine = gin.New()
	}
	return engine
}
