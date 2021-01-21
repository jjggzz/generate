// author: JGZ
// time:   2021-01-19 10:17
package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func Test_get(t *testing.T) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "localhost:3306", "customer")
	db, err := sqlx.Open("mysql", dns)
	if err != nil {
		panic(err)
		return
	}

	repo := NewRepo(db)
	ex := new(CustomerExample).AndAccessKeyLike("ehuDIF")
	list, err := repo.CustomerRepo.SelectByExample(ex)
	if err != nil {
		panic(err)
		return
	}
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println("-------------------------")
	ex.Clear()
	ex.AndIdEqualTo(1640)
	list, err = repo.CustomerRepo.SelectByExample(ex)
	if err != nil {
		panic(err)
		return
	}
	for _, e := range list {
		fmt.Println(e)
	}
}
