// author: JGZ
// time:   2021-01-19 10:17
package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

func Test_get(t *testing.T) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "localhost:3306", "customer")
	db, err := sqlx.Open("mysql", dns)
	if err != nil {
		panic(err)
		return
	}

	repo := NewRepo(db)
	_ = Customer{
		AccessKey:    "12345",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		DeleteStatus: 0,
		Phone:        "18376301877",
		Username:     "18376301877",
		Password:     "",
		Email:        "123@qq.com",
		Nickname:     "jjggzz",
		Status:       0,
	}
	count, err := repo.CustomerRepo.SelectByPrimaryKey(1640)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(count)

}
