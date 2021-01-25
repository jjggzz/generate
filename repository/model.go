package repository

import (
	"time"
)

// 表名:customer
// comment:用户表
type Customer struct {
	Id           int64     `db:"id" json:"id,omitempty"`                      // 字段: [id] 类型:[bigint(20)] 注释: [主键ID]
	AccessKey    string    `db:"access_key" json:"accessKey,omitempty"`       // 字段: [access_key] 类型:[varchar(64)] 注释: [业务key]
	CreateTime   time.Time `db:"create_time" json:"createTime,omitempty"`     // 字段: [create_time] 类型:[datetime] 注释: [创建时间]
	UpdateTime   time.Time `db:"update_time" json:"updateTime,omitempty"`     // 字段: [update_time] 类型:[datetime] 注释: [修改时间]
	DeleteStatus int32     `db:"delete_status" json:"deleteStatus,omitempty"` // 字段: [delete_status] 类型:[int(11)] 注释: [状态 0:未删除 1:已删除]
	Phone        string    `db:"phone" json:"phone,omitempty"`                // 字段: [phone] 类型:[char(11)] 注释: [电话号码]
	Username     string    `db:"username" json:"username,omitempty"`          // 字段: [username] 类型:[varchar(32)] 注释: [用户名]
	Password     string    `db:"password" json:"password,omitempty"`          // 字段: [password] 类型:[varchar(64)] 注释: [密码]
	Email        string    `db:"email" json:"email,omitempty"`                // 字段: [email] 类型:[varchar(64)] 注释: [邮箱]
	Nickname     string    `db:"nickname" json:"nickname,omitempty"`          // 字段: [nickname] 类型:[varchar(32)] 注释: [昵称]
	Status       int32     `db:"status" json:"status,omitempty"`              // 字段: [status] 类型:[int(11)] 注释: [账号状态 0:冻结 1:启用]

}
