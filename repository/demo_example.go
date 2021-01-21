package repository

import ()

type DemoExample struct {
	criteria []struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}
}

func (ex *DemoExample) AndIdIsNull() *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id is null", noValue: true})
	return ex
}

func (ex *DemoExample) AndIdIsNotNull() *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndIdEqualTo(param int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id = ?", param1: param})
	return ex
}

func (ex *DemoExample) AndIdNotEqualTo(param int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id <> ?", param1: param})
	return ex
}

func (ex *DemoExample) AndIdGreaterThan(param int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id > ?", param1: param})
	return ex
}

func (ex *DemoExample) AndIdGreaterThanOrEqualTo(param int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id >= ?", param1: param})
	return ex
}
func (ex *DemoExample) AndIdLessThan(param int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id < ?", param1: param})
	return ex
}

func (ex *DemoExample) AndIdLessThanOrEqualTo(param int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id <= ?", param1: param})
	return ex
}

func (ex *DemoExample) AndIdIn(param []int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id in (?)", param1: param})
	return ex
}

func (ex *DemoExample) AndIdNotIn(param []int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id not in (?)", param1: param})
	return ex
}

func (ex *DemoExample) AndIdBetween(param1 int32, param2 int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndIdNotBetween(param1 int32, param2 int32) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeIsNull() *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType is null", noValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeIsNotNull() *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeEqualTo(param []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType = ?", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeNotEqualTo(param []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType <> ?", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeGreaterThan(param []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType > ?", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeGreaterThanOrEqualTo(param []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType >= ?", param1: param})
	return ex
}
func (ex *DemoExample) AndBitTypeLessThan(param []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType < ?", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeLessThanOrEqualTo(param []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType <= ?", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeIn(param [][]byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType in (?)", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeNotIn(param [][]byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType not in (?)", param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeBetween(param1 []byte, param2 []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeNotBetween(param1 []byte, param2 []byte) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and bitType not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeIsNull() *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType is null", noValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeIsNotNull() *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeEqualTo(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType = ?", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotEqualTo(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType <> ?", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeGreaterThan(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType > ?", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeGreaterThanOrEqualTo(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType >= ?", param1: param})
	return ex
}
func (ex *DemoExample) AndLongblobTypeLessThan(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType < ?", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeLessThanOrEqualTo(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType <= ?", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeIn(param []string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType in (?)", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotIn(param []string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType not in (?)", param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeBetween(param1 string, param2 string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotBetween(param1 string, param2 string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeLike(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType like ?", param1: param})
	return ex
}
func (ex *DemoExample) AndLongblobTypeNotLike(param string) *DemoExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and longblobType not like ?", param1: param})
	return ex
}
