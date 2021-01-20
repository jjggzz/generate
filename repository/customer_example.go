package repository

import (
	"time"
)

type CustomerExample struct {
	Criteria []struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}
}

func (ex *CustomerExample) AndIdIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndIdIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndIdEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdNotEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdGreaterThan(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdGreaterThanOrEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndIdLessThan(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdLessThanOrEqualTo(param int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdIn(param []int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdNotIn(param []int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndIdBetween(param1 int64, param2 int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndIdNotBetween(param1 int64, param2 int64) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndAccessKeyLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and access_key not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeGreaterThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeGreaterThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndCreateTimeLessThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeLessThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and create_time not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeGreaterThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeGreaterThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndUpdateTimeLessThan(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeLessThanOrEqualTo(param time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotIn(param []time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and update_time not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusGreaterThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusGreaterThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndDeleteStatusLessThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusLessThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusBetween(param1 int32, param2 int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotBetween(param1 int32, param2 int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and delete_status not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndPhoneLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneNotBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and phone not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndUsernameLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameNotBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and username not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndPasswordLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordNotBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and password not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndEmailLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailNotBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and email not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameNotEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameGreaterThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndNicknameLessThan(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameLessThanOrEqualTo(param string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameNotIn(param []string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameNotBetween(param1 string, param2 string) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and nickname not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusIsNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusIsNotNull() *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status = ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusNotEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status <> ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusGreaterThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status > ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusGreaterThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status >= ?", Param1: param})
	return ex
}
func (ex *CustomerExample) AndStatusLessThan(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status < ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusLessThanOrEqualTo(param int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status <= ?", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusNotIn(param []int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status not in (?)", Param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusBetween(param1 int32, param2 int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusNotBetween(param1 int32, param2 int32) *CustomerExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and status not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}
