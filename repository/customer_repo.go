package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type customerRepo struct {
	db *sqlx.DB
}

// execute fail return -1 and err
// execute success return count rows and nil
func (repo *customerRepo) Count() (int64, error) {
	var count int64
	err := repo.db.Get(&count, "select count(*) from customer")
	if err != nil {
		return -1, err
	}
	return count, nil
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *customerRepo) Insert(customer *Customer) (int64, error) {
	result, err := repo.db.Exec("insert into customer(id, access_key, create_time, update_time, delete_status, phone, username, password, email, nickname, status) value (?,?,?,?,?,?,?,?,?,?,?)", customer.Id, customer.AccessKey, customer.CreateTime, customer.UpdateTime, customer.DeleteStatus, customer.Phone, customer.Username, customer.Password, customer.Email, customer.Nickname, customer.Status)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *customerRepo) UpdateByPrimaryKey(primaryKey int64, customer *Customer) (int64, error) {
	result, err := repo.db.Exec("update customer set id=?, access_key=?, create_time=?, update_time=?, delete_status=?, phone=?, username=?, password=?, email=?, nickname=?, status=? where id = ?", customer.Id, customer.AccessKey, customer.CreateTime, customer.UpdateTime, customer.DeleteStatus, customer.Phone, customer.Username, customer.Password, customer.Email, customer.Nickname, customer.Status, primaryKey)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return -1 and err
// execute success return affected rows and nil
func (repo *customerRepo) DeleteByPrimaryKey(primaryKey int64) (int64, error) {
	result, err := repo.db.Exec("delete from customer where id = ?", primaryKey)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// execute fail return nil and err
// execute success but no record return nil and nil
// execute success has record return data and nil
func (repo *customerRepo) SelectByPrimaryKey(primaryKey int64) (*Customer, error) {
	customer := new(Customer)
	err := repo.db.Get(customer, "select id, access_key, create_time, update_time, delete_status, phone, username, password, email, nickname, status from customer where id = ?", primaryKey)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// execute fail return empty slice and err
// execute success but no record return empty slice and nil
// execute success has record return data and nil
func (repo *customerRepo) SelectByExample(example *CustomerExample) ([]*Customer, error) {
	var list []*Customer
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
	query, args, err := sqlx.In("select id, access_key, create_time, update_time, delete_status, phone, username, password, email, nickname, status from customer "+condition, params...)
	if err != nil {
		return list, err
	}
	return list, repo.db.Select(&list, query, args...)
}
