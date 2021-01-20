package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"

	"time"
)

type customerRepo struct {
	db *sqlx.DB
}

type CustomerExample struct {
	Criteria []struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}
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
	query, args, err := sqlx.In("select id, access_key, create_time, update_time, delete_status, phone, username, password, email, nickname, status from customer "+condition, params...)
	if err != nil {
		return list, err
	}
	return list, repo.db.Select(&list, query, args...)
}

func (ex *CustomerExample) AndIdIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndIdIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndIdEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndIdNotEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndIdGreaterThan(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndIdGreaterThanOrEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndIdLessThan(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndIdLessThanOrEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndIdIn(param []int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndIdNotIn(param []int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and id not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndAccessKeyLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndAccessKeyNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and access_key not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeGreaterThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeGreaterThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndCreateTimeLessThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeLessThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndCreateTimeNotIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and create_time not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeGreaterThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeGreaterThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndUpdateTimeLessThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeLessThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndUpdateTimeNotIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and update_time not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusGreaterThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusGreaterThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndDeleteStatusLessThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusLessThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndDeleteStatusNotIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and delete_status not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndPhoneIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPhoneNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPhoneGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPhoneGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndPhoneLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPhoneLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPhoneIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndPhoneNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and phone not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndUsernameIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUsernameNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUsernameGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUsernameGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndUsernameLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUsernameLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndUsernameIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndUsernameNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and username not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndPasswordIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPasswordNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPasswordGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPasswordGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndPasswordLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPasswordLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndPasswordIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndPasswordNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and password not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndEmailIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndEmailNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndEmailGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndEmailGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndEmailLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndEmailLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndEmailIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndEmailNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and email not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndNicknameIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndNicknameNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndNicknameGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndNicknameGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndNicknameLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndNicknameLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndNicknameIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndNicknameNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and nickname not in (?)", Param: param})
	return ex
}

func (ex *CustomerExample) AndStatusIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status = ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndStatusNotEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status <> ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndStatusGreaterThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status > ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndStatusGreaterThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status >= ?", Param: param})
	return ex
}
func (ex *CustomerExample) AndStatusLessThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status < ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndStatusLessThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status <= ?", Param: param})
	return ex
}

func (ex *CustomerExample) AndStatusIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status in (?)", Param: param})
	return ex
}
func (ex *CustomerExample) AndStatusNotIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment string
		Param    interface{}
		noValue  bool
	}{Fragment: "and status not in (?)", Param: param})
	return ex
}
