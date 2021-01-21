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
	ex := new(CustomerExample).AndIdIn([]int64{1639, 1640})
	count, err := repo.CustomerRepo.DeleteByExample(ex)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(count)
}
