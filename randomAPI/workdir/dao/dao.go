package dao

import (
	"a/workdir/config"
	"a/workdir/model"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var D *sql.DB

func DbInit()error{
	var err error
	DB,err=gorm.Open(mysql.Open(config.CfgDsn()),&gorm.Config{})
	if err!=nil{return err}
	D,err=DB.DB()
	if err!=nil{
		return err}
	err= D.Ping()
	if err!=nil{return err}
	return DB.AutoMigrate(&model.Sentences{})

}
