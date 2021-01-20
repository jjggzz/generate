package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type demoRepo struct {
	db *sqlx.DB
}

type DemoExample struct {
	Criteria []struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}
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

// execute fail return empty slice and err
// execute success but no record return empty slice and nil
// execute success has record return data and nil
func (repo *demoRepo) SelectByExample(example *DemoExample) ([]*Demo, error) {
	var list []*Demo
	criteria := example.Criteria
	var params []interface{}
	var condition = ""
	if len(criteria) > 0 {
		var fragments []string
		for _, e := range criteria {
			fragments = append(fragments, e.Fragment)
			if !e.noValue {
				params = append(params, e.Param)
			}
		}
		condition += "where " + strings.TrimLeft(strings.Join(fragments, " "), "and")
	}
	return list, repo.db.Select(&list, "select id, bitType, longblobType from demo "+condition, params...)
}

func (ex *DemoExample) AndIdEqualTo(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id = ?", Param: param})
	return ex
}

func (ex *DemoExample) AndIdNotEqualTo(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id <> ?", Param: param})
	return ex
}
func (ex *DemoExample) AndIdIsNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id is null", noValue: true})
	return ex
}
func (ex *DemoExample) AndIdIsNotNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeEqualTo(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and bitType = ?", Param: param})
	return ex
}

func (ex *DemoExample) AndBitTypeNotEqualTo(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and bitType <> ?", Param: param})
	return ex
}
func (ex *DemoExample) AndBitTypeIsNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and bitType is null", noValue: true})
	return ex
}
func (ex *DemoExample) AndBitTypeIsNotNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and bitType is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeEqualTo(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and longblobType = ?", Param: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotEqualTo(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and longblobType <> ?", Param: param})
	return ex
}
func (ex *DemoExample) AndLongblobTypeIsNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and longblobType is null", noValue: true})
	return ex
}
func (ex *DemoExample) AndLongblobTypeIsNotNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and longblobType is not null", noValue: true})
	return ex
}
