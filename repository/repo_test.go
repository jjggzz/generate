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
	ex := new(CustomerExample).AndIdEqualTo(1641)
	cus, err := repo.CustomerRepo.SelectByPrimaryKey(1641)
	cus.Phone = "11111111111"
	count, err := repo.CustomerRepo.UpdateByExample(ex, cus)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(count)
}
