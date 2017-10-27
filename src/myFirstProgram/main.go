package main

import (
	_ "myFirstProgram/routers"
	"github.com/astaxie/beego"
)

func main() {

	//Go语言内部其实已经提供了http.ServeFile，通过这个函数可以实现静态文件的服务。beego针对这个功能进行了一层封装，通过下面的方式进行静态文件注册
	beego.SetStaticPath("/static","public")

	//beego支持多个目录的静态文件注册，用户可以注册如下的静态文件目录：
	beego.SetStaticPath("/images","images")
	beego.SetStaticPath("/css","css")
	beego.SetStaticPath("/js","js")

	beego.Run()
}

