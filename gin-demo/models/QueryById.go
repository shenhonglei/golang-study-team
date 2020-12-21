package models

import (
	"gin-demo/initDB"
)

type SysLogs struct {
	Id       int64  `json:"id,string"` //先定义好json结构
	Username string `json:"username,string"`
	Uri      string `json:"uri,string"`
}

func (syslog SysLogs) QueryById() SysLogs {
	initDB.Db.Find(&syslog)
	return syslog
}
