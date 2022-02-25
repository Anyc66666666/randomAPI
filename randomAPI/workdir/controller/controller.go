package controller

import (
	"a/sentence"
	"a/workdir/config"
	"a/workdir/dao"
	"a/workdir/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
)

func Video(c *gin.Context) {
	//	b,_:= ioutil.ReadFile("./1.txt")
	//	x,_:=strconv.Atoi(string(b))
	x:=VNums()
	n:=rand.Intn(x)+1
	//	 fmt.Println(n)
	// 		 paths:=fmt.Sprintf("/file/%d.mp4",n)
	// 		 c.Request.URL.Path=paths  //路由重定向
	// 		 r.HandleContext(c)
	paths:=fmt.Sprintf(config.CfgHost()+":"+config.CfgPort()+"/file/%d.mp4",n)
	c.Redirect(http.StatusTemporaryRedirect,paths)

}
func Music(c *gin.Context) {
	// 		b,_:= ioutil.ReadFile("./2.txt")
	// 		x,_:=strconv.Atoi(string(b))
	x:=MNums()
	n:=rand.Intn(x)+1
	//	 fmt.Println(n)
	// 		 paths:=fmt.Sprintf("/file/%d.mp4",n)
	// 		 c.Request.URL.Path=paths  //路由重定向
	// 		 r.HandleContext(c)
	paths:=fmt.Sprintf(config.CfgHost()+":"+config.CfgPort()+"/file/%d.mp3",n)
	c.Redirect(http.StatusTemporaryRedirect,paths)

}
func Picture(c *gin.Context) {
	// 		b,_:= ioutil.ReadFile("./2.txt")
	// 		x,_:=strconv.Atoi(string(b))
	x:=PNums()
	n:=rand.Intn(x)+1
	//	 fmt.Println(n)
	// 		 paths:=fmt.Sprintf("/file/%d.mp4",n)
	// 		 c.Request.URL.Path=paths  //路由重定向
	// 		 r.HandleContext(c)
	paths:=fmt.Sprintf(config.CfgHost()+":"+config.CfgPort()+"/file/%d.png",n)
	c.Redirect(http.StatusTemporaryRedirect,paths)

}
func Sentence(c *gin.Context){
	c.JSON(200,sentence.ST())

}
func Admin(c *gin.Context) {
	c.HTML(200,"admin.html",gin.H{//模板渲染
		"video":VNums(),
		"music":MNums(),
		"picture":PNums(),
		"sentence":SNums(),     //   http://127.0.0.1:8468/sentences/add
		"SentencesAdd":config.CfgHost()+":"+config.CfgPort()+"/sentences/add",
		"SentencesUpdate":config.CfgHost()+":"+config.CfgPort()+"/sentences/update",
		"SentencesDelete":config.CfgHost()+":"+config.CfgPort()+"/sentences/delete",
		"SentencesAll":config.CfgHost()+":"+config.CfgPort()+"/sentences/all",
		"FileUpload":config.CfgHost()+":"+config.CfgPort()+"/file/upload",
	})
}
func DownFile(c *gin.Context){
	// 	for i:=0;i<108;i++{
	// 		filename:=fmt.Sprintf("./%d.mp3",i)
	// 		c.File(filename)//响应一个文件
	// }
	var filename string
	name:=c.Param("name")
	b:=strings.Split(name,".")
	if string(b[len(b)-1])=="mp3"{
		filename=path.Join("./music/",name)
	}
	if string(b[len(b)-1])=="mp4"{
		filename=path.Join("./video/",name)
	}
	if string(b[len(b)-1])=="jpg"||string(b[len(b)-1])=="png"||string(b[len(b)-1])=="jpeg"{
		filename=path.Join("./picture/",name)
	}

	c.File(filename)//响应一个文件

}
func FileUpload(c *gin.Context) {
	var NewName string

	//从请求中读取文件
	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	//将读取的文件保存在本地（服务端本地）
	//dst:=fmt.Sprintf("./%s",f.Filename)
	b:=strings.Split(f.Filename,".")
	if string(b[len(b)-1])=="mp3"{
		NewName=fmt.Sprintf("./music/%d.mp3",MNums()+1)
	}
	if string(b[len(b)-1])=="mp4"{
		NewName=fmt.Sprintf("./video/%d.mp4",VNums()+1)
	}
	if string(b[len(b)-1])=="jpg"||string(b[len(b)-1])=="jpeg"||string(b[len(b)-1])=="png"{
		NewName=fmt.Sprintf("./picture/%d.png",PNums()+1)
	}

	OldName := path.Join("./cache", f.Filename)
	//NewName:=fmt.Sprintf("./2/%d.mp4",VNums()+1)
	c.SaveUploadedFile(f, OldName)
	os.Rename(OldName,NewName)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg":b[len(b)-1],
	})

}
func FilesUpload(c *gin.Context) {
	var NewName string

	//从请求中读取文件
	form,err:=c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	files:=form.File["f1"]

	for _,f:=range files{
		//将读取的文件保存在本地（服务端本地）
		//dst:=fmt.Sprintf("./%s",f.Filename)
		b:=strings.Split(f.Filename,".")
		if string(b[len(b)-1])=="mp3"{
			NewName=fmt.Sprintf("./music/%d.mp3",MNums()+1)
		}
		if string(b[len(b)-1])=="mp4"{
			NewName=fmt.Sprintf("./video/%d.mp4",VNums()+1)
		}
		if string(b[len(b)-1])=="jpg"||string(b[len(b)-1])=="jpeg"||string(b[len(b)-1])=="png"{
			NewName=fmt.Sprintf("./picture/%d.png",PNums()+1)
		}

		OldName := path.Join("./cache", f.Filename)
		//NewName:=fmt.Sprintf("./2/%d.mp4",VNums()+1)
		c.SaveUploadedFile(f, OldName)
		os.Rename(OldName,NewName)
		c.JSON(200, gin.H{
			"status": "ok",
			"msg":b[len(b)-1],
		})
	}



}
func GetAll(c *gin.Context){
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)
		return}
	var s []model.Sentences
	dao.DB.Debug().Find(&s)
	c.JSON(200,s)}

func GetRandom(c *gin.Context){
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)
		return}
	var s []model.Sentences
	var ss model.Sentences
	t:=dao.DB.Debug().Find(&s)
	n:=rand.Intn(int(t.RowsAffected))+1
	dao.DB.Debug().Where("id=?",n).First(&ss)
	c.JSON(200,ss.Sentence)


}
func Add(c *gin.Context){
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)
		return}
	var p model.ParamS
	var s model.Sentences
	err=c.ShouldBind(&p)
	if err!=nil{
		fmt.Printf("bind failed,err:%v\n",err)
		c.JSON(200,err)
		return}
	if strings.TrimSpace(p.Sentence)==""{
		c.JSON(200,gin.H{
			"msg":"不能只输入空白哦！",
		})
		return}
	s.Sentence=p.Sentence
	dao.DB.Debug().Create(&s)
	c.JSON(200,gin.H{
		"msg":"添加成功哦！",
	})
}
func Update(c *gin.Context){
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)
		c.JSON(200,err)
		return}
	var s model.Sentences
	err=c.ShouldBind(&s)
	if err!=nil{
		c.JSON(200,err)
		fmt.Printf("bind failed,err:%v\n",err)
		return}
	if strings.TrimSpace(s.Sentence)==""{
		return}
	dao.DB.Debug().Where("id=?",s.ID).Update("sentence",s.Sentence)
	c.JSON(200,gin.H{
		"msg":"修改成功哦！",
	})


}
func Delete(c *gin.Context){
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)
		c.JSON(200,err)
		return}
	var s model.Sentences
	var p model.ParamI
	err=c.ShouldBind(&p)
	if err!=nil{
		c.JSON(200,err)
		fmt.Printf("bind failed,err:%v\n",err)
		return}
	s.ID=p.ID
	dao.DB.Debug().Where("id=?",s.ID).Delete(&s)
	c.JSON(200,gin.H{
		"msg":"删除成功了呀!",
	})
}

func SNums()int{
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)}
	var s []model.Sentences
	t:=dao.DB.Debug().Find(&s)
	return int(t.RowsAffected)
}
func MNums()int{
	files,_:=ioutil.ReadDir("./music")
	return 	len(files)
}

func VNums()int{
	files,_:=ioutil.ReadDir("./video")
	return 	len(files)
}
func PNums()int{
	files,_:=ioutil.ReadDir("./picture")
	return 	len(files)
}
