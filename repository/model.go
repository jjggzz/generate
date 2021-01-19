package repository

import (
	"time"
)

// 表名:customer
// comment:用户表
type Customer struct {
	Id           int64     `db:"id" json:"id"`                      // 字段: [id] 类型:[bigint(20)] 注释: [主键ID]
	AccessKey    string    `db:"access_key" json:"accessKey"`       // 字段: [access_key] 类型:[varchar(64)] 注释: [业务key]
	CreateTime   time.Time `db:"create_time" json:"createTime"`     // 字段: [create_time] 类型:[datetime] 注释: [创建时间]
	UpdateTime   time.Time `db:"update_time" json:"updateTime"`     // 字段: [update_time] 类型:[datetime] 注释: [修改时间]
	DeleteStatus int32     `db:"delete_status" json:"deleteStatus"` // 字段: [delete_status] 类型:[int(11)] 注释: [状态 0:未删除 1:已删除]
	Phone        string    `db:"phone" json:"phone"`                // 字段: [phone] 类型:[char(11)] 注释: [电话号码]
	Username     string    `db:"username" json:"username"`          // 字段: [username] 类型:[varchar(32)] 注释: [用户名]
	Password     string    `db:"password" json:"password"`          // 字段: [password] 类型:[varchar(64)] 注释: [密码]
	Email        string    `db:"email" json:"email"`                // 字段: [email] 类型:[varchar(64)] 注释: [邮箱]
	Nickname     string    `db:"nickname" json:"nickname"`          // 字段: [nickname] 类型:[varchar(32)] 注释: [昵称]
	Status       int32     `db:"status" json:"status"`              // 字段: [status] 类型:[int(11)] 注释: [账号状态 0:冻结 1:启用]

}

// 表名:demo
// comment:
type Demo struct {
	Id           int32  `db:"id" json:"id"`                     // 字段: [id] 类型:[int(11)] 注释: []
	BitType      []byte `db:"bitType" json:"bitType"`           // 字段: [bitType] 类型:[bit(1)] 注释: []
	LongblobType string `db:"longblobType" json:"longblobType"` // 字段: [longblobType] 类型:[varchar(255)] 注释: []

}
