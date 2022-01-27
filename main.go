package main

import (
	_ "beego_project/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	beego.Run()
}