package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file config.yaml")
		return
	}

	// 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Println("load config file error:", err)
		return
	}

	// 加载日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Println("init logger error:", err)
		return
	}

	// MySQL
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql error:", err)
		return
	}
	defer mysql.Close()

	// redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Println("init redis error:", err)
		return
	}
	defer redis.Close()

	// snowflake
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Println("init snowflake error:", err)
		return
	}

	// gin框架翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Println("init trans error", err)
		return
	}

	// 注册路由
	r := router.InitRouter(setting.Conf.Mode)
	port := fmt.Sprintf(":%d", setting.Conf.Port)
	err := r.Run(port)
	if err != nil {
		fmt.Println("run server error", err)
	}
}