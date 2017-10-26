package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//给当前的控制器添加了一个新的方法（或action）Get，通过将main *MainController指定为这个方法的接收器
func (c *MainController) Get() {
	//初始化了三个模板变量，Website, Email 和 EmailName，把它们保存在Controller的一个叫Data的字段里面，该字段是一个map用来表示模板变量和值
	c.Data["Website"] = "My Website"
	c.Data["Email"] = "zp335497099@sina.com"
	c.Data["EmailName"] = "pengzhou"
	//将this.TplNames的值设置为default/hello-sitepoint.tpl来指定模板文件的名称。默认情况下，Beego会在views文件夹下查找指定的文件
	c.TplName = "default/hello-sitepoint.tpl"
}


//Beego的模板层继承自 Go 的 html/template package。想了解更多在模板中使用变量的例子，请参考text/template包中的例子。
//html/template的功能很多，所以我在这里介绍几个相关的功能。所有的模板变量都属于全局的上下文，可以通过点操作符来访问，通过{{}}语法来嵌入。
//想要访问我们在之前在controller的action中定义的那山歌模板变量，可以使用{{.Website}},{{.Email}}, 和 {{.EmailName}}


