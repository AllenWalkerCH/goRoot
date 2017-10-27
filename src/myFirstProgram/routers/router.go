package routers






import (
	//这里是配置路由相关
	"myFirstProgram/controllers"
	"github.com/astaxie/beego"
)

//func init() {
//    beego.Router("/", &controllers.MainController{})
//}

func init() {

	beego.Router("/", &controllers.MainController{})
	//beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")

	//注意最后一行，它所做的事是当我们尝试去访问/hello-world的时候，它会调用MainController中的HelloSitepoint action。保存并等一
	// 下以便完成重新编译，之后在浏览器中打开http://localhost:8080/hello-world


	//增加传入参数     这样一来访问这个接口就必须要传入参数了
	beego.Router("/hello-world/:id[0-9]+",&controllers.MainController{},"get:HelloSitepoint")


	//Beego简化了通过所有方法类型的action访问，或将某个action的访问限制到一个或几个方法类型。也是在router.go中，将/hello-world路由做如下改变：
	//将请求分发
	//beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get,post:Get")


	/*
	# PUT request
	curl -X PUT http://localhost:8080/hello-world/213

	# DELETE request
	curl -X DELETE http://localhost:8080/hello-world/213
	*/

	}


