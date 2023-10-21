package main

import (
	"RealWorld/router"
)

func main() {
	//gorm链接数据库
	ginServer := router.Router()
	ginServer.Run(":8082")
}
