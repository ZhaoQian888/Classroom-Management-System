package main

import (
	"Classroom-Management-System/config"
	"Classroom-Management-System/router"
)

func main() {
	config.Init()

	//初始化路由
	r := router.SetRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
