package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type demoRepo struct {
	db *sqlx.DB
}

// execute fail return -1 and err
// execute success return count rows and nil
func (repo *demoRepo) Count() (int64, error) {
	var count int64
	err := repo.db.Get(&count, "select count(*) from demo")
	if err != nil {
		return -1, err
	}
	return count, nil
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *demoRepo) Insert(demo *Demo) (int64, error) {
	result, err := repo.db.Exec("insert into demo(id, bitType, longblobType) value (?,?,?)", demo.Id, demo.BitType, demo.LongblobType)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *demoRepo) UpdateByPrimaryKey(primaryKey int32, demo *Demo) (int64, error) {
	result, err := repo.db.Exec("update demo set id=?, bitType=?, longblobType=? where id = ?", demo.Id, demo.BitType, demo.LongblobType, primaryKey)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *demoRepo) DeleteByPrimaryKey(primaryKey int32) (int64, error) {
	result, err := repo.db.Exec("delete from demo where id = ?", primaryKey)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return nil and err
// execute success but no record return nil and nil
// execute success has record return data and nil
func (repo *demoRepo) SelectByPrimaryKey(primaryKey int32) (*Demo, error) {
	demo := new(Demo)
	err := repo.db.Get(demo, "select id, bitType, longblobType from demo where id = ?", primaryKey)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return demo, nil
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *demoRepo) DeleteByExample(example *DemoExample) (int64, error) {
	criteria := example.criteria
	var params []interface{}
	var condition = ""
	if len(criteria) > 0 {
		var fragments []string
		for _, e := range criteria {
			fragments = append(fragments, e.fragment)
			if !e.noValue {
				params = append(params, e.param1)
			}
			if e.betweenValue {
				params = append(params, e.param2)
			}
		}
		condition += "where " + strings.TrimLeft(strings.Join(fragments, " "), "and")
	}
	query, args, err := sqlx.In("delete from demo "+condition, params...)
	if err != nil {
		return -1, err
	}
	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return empty slice and err
// execute success but no record return empty slice and nil
// execute success has record return data and nil
func (repo *demoRepo) SelectByExample(example *DemoExample) ([]*Demo, error) {
	var list []*Demo
	criteria := example.criteria
	var params []interface{}
	var condition = ""
	if len(criteria) > 0 {
		var fragments []string
		for _, e := range criteria {
			fragments = append(fragments, e.fragment)
			if !e.noValue {
				params = append(params, e.param1)
			}
			if e.betweenValue {
				params = append(params, e.param2)
			}
		}
		condition += "where " + strings.TrimLeft(strings.Join(fragments, " "), "and")
	}
	query, args, err := sqlx.In("select id, bitType, longblobType from demo "+condition, params...)
	if err != nil {
		return list, err
	}
	return list, repo.db.Select(&list, query, args...)
}
