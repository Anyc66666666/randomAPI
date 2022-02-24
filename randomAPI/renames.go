package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
//批量修改文件名 用于批量上传文件or整理不规范文件名
func main(){
	
RenameFiles("./music","mp3")


}
func RenameFiles(Path string,FileFormat string){
	i:=1
	files,err:=ioutil.ReadDir(Path)
	if err!=nil{fmt.Printf("readDir failed ,err:%v\n",err)
		return
	}
	for _,f:=range files{
		fmt.Println(f.Name())
		newName:=fmt.Sprintf("%d."+FileFormat,i)
		i++
		err:=os.Rename(Path+"/"+f.Name(),Path+"/"+newName)
		if err!=nil{fmt.Printf("rename failed,err:%v\n",err)
			return}


	}
}