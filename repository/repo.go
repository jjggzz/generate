package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

type demoRepo struct {
	db *sqlx.DB
}

func (repo *customerRepo) Insert(customer *Customer) error {
	return nil
}
func (repo *customerRepo) Update(customer *Customer) error {
	return nil
}
func (repo *customerRepo) DeleteByPrimaryKey(primaryKey int64) error {
	return nil
}
func (repo *customerRepo) SelectByPrimaryKey(primaryKey int64) (*Customer, error) {
	return nil, nil
}

func (repo *demoRepo) Insert(demo *Demo) error {
	return nil
}
func (repo *demoRepo) Update(demo *Demo) error {
	return nil
}
func (repo *demoRepo) DeleteByPrimaryKey(primaryKey int32) error {
	return nil
}
func (repo *demoRepo) SelectByPrimaryKey(primaryKey int32) (*Demo, error) {
	return nil, nil
}
