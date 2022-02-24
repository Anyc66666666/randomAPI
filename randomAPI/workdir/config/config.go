package config

import (
	"a/workdir/model"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func CfgDsn()string{
	var c model.Config
	config,err:=ioutil.ReadFile("./workdir/config/config.yaml")
	if err!=nil{fmt.Printf("read file failed ,err:%v\n",err)}
	err=yaml.Unmarshal(config,&c)
	if err!=nil{fmt.Printf("yaml failed ,err:%v\n",err)}
	return c.DSN

}
func CfgHost()string{
	var c model.Config
	config,err:=ioutil.ReadFile("./workdir/config/config.yaml")
	if err!=nil{fmt.Printf("read file failed ,err:%v\n",err)}
	err=yaml.Unmarshal(config,&c)
	if err!=nil{fmt.Printf("yaml failed ,err:%v\n",err)}
	return c.HOST

}
func CfgPort()string{
	var c model.Config
	config,err:=ioutil.ReadFile("./workdir/config/config.yaml")
	if err!=nil{fmt.Printf("read file failed ,err:%v\n",err)}
	err=yaml.Unmarshal(config,&c)
	if err!=nil{fmt.Printf("yaml failed ,err:%v\n",err)}
	return c.PORT

}