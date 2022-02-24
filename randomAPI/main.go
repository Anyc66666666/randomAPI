package main

import (
	"a/workdir/dao"
	"a/workdir/router"
	"fmt"
)

func main(){
	//数据库、模型初始化
	err:=dao.DbInit()
	if err!=nil{fmt.Printf("Dbinit failed,err:%v\n",err)
		return}


	//路由初始化
	router.RoutersInit()



}













