package main

import (
	"auroralab/router"
	"auroralab/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMysql()
	appRouter := router.SetupRouter()
	appRouter.Run(":8079") // 监听并在 0.0.0.0:8080 上启动服务

}
