package repository

import ()

// 表名:demo
// comment:
type Demo struct {
	Id           int32  `db:"id" json:"id"`                     // 字段: [id] 类型:[int(11)] 注释: []
	BitType      []byte `db:"bitType" json:"bitType"`           // 字段: [bitType] 类型:[bit(1)] 注释: []
	LongblobType string `db:"longblobType" json:"longblobType"` // 字段: [longblobType] 类型:[varchar(255)] 注释: []

}
