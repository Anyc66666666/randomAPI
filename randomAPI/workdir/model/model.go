package model

import "gorm.io/gorm"

type Sentences struct {
	gorm.Model
	Sentence string `form:"sentence" json:"sentence"`
}
type ParamS struct {
	Sentence string `form:"sentence" json:"sentence"`
}
type ParamI struct {
	ID uint `form:"id" json:"id"`
}
type ParamSI struct {
	Sentence string `form:"sentence" json:"sentence"`
	ID uint `form:"id" json:"id"`
}
type Config struct {
	DSN string `yaml:"dsn"`
	HOST string `yaml:"host"`
	PORT string `yaml:"port"`
}
