package demo

import "time"

type Customer struct {
	Id           int64
	AccessKey    string
	CreateTime   time.Time
	UpdateTime   time.Time
	DeleteStatus int32
	Phone        string
	Username     string
	Password     string
	Email        string
	Nickname     string
	Status       int32
}
type Demo struct {
	Id           int32
	BitType      bool
	LongblobType string
}
