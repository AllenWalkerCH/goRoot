package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller

}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.TplName = "index.tpl"
}
func (c *MainController) Post()  {

	}
func (c *MainController) Put()  {

	}
func (c *MainController) Delete()  {

	}


//给当前的控制器添加了一个新的方法（或action）Get，通过将main *MainController指定为这个方法的接收器
func (main *MainController) HelloSitepoint() {

	//初始化了三个模板变量，Website, Email 和 EmailName，把它们保存在Controller的一个叫Data的字段里面，该字段是一个map用来表示模板变量和值
	main.Data["Website"] = "AllenWalker's Website"
	main.Data["Email"] = "zp335497099@sina.com"
	main.Data["EmailName"] = "Allen Walker"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	//将this.TplName的值设置为default/hello-sitepoint.tpl来指定模板文件的名称。默认情况下，Beego会在views文件夹下查找指定的文件
	//模板文件的名字是要求一定准确的
	main.TplName = "default/hello-sitepoint.tpl"

}


//Beego的模板层继承自 Go 的 html/template package。想了解更多在模板中使用变量的例子，请参考text/template包中的例子。

//html/template的功能很多，所以我在这里介绍几个相关的功能。所有的模板变量都属于全局的上下文，可以通过点操作符来访问，通过{{}}语法来嵌入。

//想要访问我们在之前在controller的action中定义的那山歌模板变量，可以使用{{.Website}},{{.Email}}, 和 {{.EmailName}}


/*

1、首先我们引入了包github.com/astaxie/beego,我们知道Go语言里面引入包会深度优先的去执行引入包的初始化(变量和init函数，更多)，
beego包中会初始化一个BeeAPP的应用，初始化一些参数。

2、定义Controller，这里我们定义了一个struct为MainController，充分利用了Go语言的组合的概念，匿名包含了beego.Controller，
这样我们的MainController就拥有了beego.Controller的所有方法。

3、定义RESTFul方法，通过匿名组合之后，其实目前的MainController已经拥有了Get、Post、Delete、Put等方法，这些方法是分别用来对
应用户请求的Method函数，如果用户发起的是POST请求，那么就执行Post函数。所以这里我们定义了MainController的Get方法用来重写继承的Get函数，
这样当用户GET请求的时候就会执行该函数。

4、定义main函数，所有的Go应用程序和C语言一样都是Main函数作为入口，所以我们这里定义了我们应用的入口。

5、Router注册路由，路由就是告诉beego，当用户来请求的时候，该如何去调用相应的Controller，这里我们注册了请求/的时候，请求到MainController。
这里我们需要知道，Router函数的两个参数函数，第一个是路径，第二个是Controller的指针。

6、Run应用，最后一步就是把在1中初始化的BeeApp开启起来，其实就是内部监听了8080端口:Go默认情况会监听你本机所有的IP上面的8080端口

停止服务的话，请按ctrl+c

 */





 //****************************************************以上解说完毕****************************************************

 /*

 路由的主要功能是实现从请求地址到实现方法，beego中封装了Controller，所以路由是从路径到ControllerInterface的过程，ControllerInterface的方法有如下：
type ControllerInterface interface {
    Init(ct *Context, cn string)
    Prepare()
    Get()
    Post()
    Delete()
    Put()
    Head()
    Patch()
    Options()
    Finish()
    Render() error
}
 这些方法beego.Controller都已经实现了，所以只要用户定义struct的时候匿名包含就可以了。当然更灵活的方法就是用户可以去自定义类似的方法，然后实现自己的逻辑。


 用户可以通过如下的方式进行路由设置：  通过路径到控制器    一个路径对应一个控制器处理     一个接口对应一段逻辑
beego.Router("/", &controllers.MainController{})
beego.Router("/admin", &admin.UserController{})
beego.Router("/admin/index", &admin.ArticleController{})
beego.Router("/admin/addpkg", &admin.AddController{})

 */

/*

beego.Router("/api/:id([0-9]+)", &controllers.RController{})
自定义正则匹配 //匹配 /api/123 :id= 123

beego.Router("/news/:all", &controllers.RController{})
全匹配方式 //匹配 /news/path/to/123.html :all= path/to/123.html

beego.Router("/user/:username([\w]+)", &controllers.RController{})
正则字符串匹配 //匹配 /user/astaxie :username = astaxie

beego.Router("/download/*.*", &controllers.RController{})
*匹配方式 //匹配 /download/file/api.xml :path= file/api :ext=xml

beego.Router("/download/ceshi/*", &controllers.RController{})
*全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json

beego.Router("/:id:int", &controllers.RController{})
int类型设置方式 //匹配 :id为int类型，框架帮你实现了正则([0-9]+)

beego.Router("/:hi:string", &controllers.RController{})
string类型设置方式 //匹配 :hi为string类型。框架帮你实现了正则([\w]+)

*/








