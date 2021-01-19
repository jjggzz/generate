package repository

import (
	"github.com/jmoiron/sqlx"
)

type demoRepo struct {
	db *sqlx.DB
}

const demoResultSql = "id, bitType, longblobType"

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
	demo := new(Demo)
	return demo, repo.db.Get(demo, "select "+demoResultSql+" from demo where id = ?", primaryKey)
}
