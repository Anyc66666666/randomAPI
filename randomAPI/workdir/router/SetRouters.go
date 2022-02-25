package router

import (
	"a/workdir/config"
	"a/workdir/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func RoutersInit(){
	r:=gin.Default()
	r=SetRouters(r)
	err:=r.Run(":"+config.CfgPort())
	if err!=nil{
		fmt.Printf("run failed,err:%v\n",err)
		return
	}

}

func SetRouters(r *gin.Engine)*gin.Engine{
	r.GET("/video",controller.Video)
	r.GET("/music",controller.Music)
	r.GET("/picture",controller.Picture)
	r.GET("/sentence",controller.Sentence)

	//------------------------
	//sentences with mysql web html
	//加载静态文件
	//r.Static("/xxx","./statics")
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string)template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./workdir/admin.html")//解析模板
	//r.LoadHTMLGlob("./**/*")  //解析模板 // **目录 //*所有文件
	r.GET("/admin",controller.Admin)
	r.GET("/sentences/all",controller.GetAll)
	r.GET("/sentences",controller.GetRandom)
	r.POST("/sentences/add",controller.Add)
	r.POST("/sentences/update",controller.Update)
	r.POST("/sentences/delete",controller.Delete)
	r.POST("/file/upload",controller.FilesUpload)
	r.GET("/file/:name",controller.DownFile)

	return r

}