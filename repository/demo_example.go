package repository

import ()

type DemoExample struct {
	Criteria []struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}
}

func (ex *DemoExample) AndIdIsNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id is null", noValue: true})
	return ex
}

func (ex *DemoExample) AndIdIsNotNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndIdEqualTo(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id = ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdNotEqualTo(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id <> ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdGreaterThan(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id > ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdGreaterThanOrEqualTo(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id >= ?", Param1: param})
	return ex
}
func (ex *DemoExample) AndIdLessThan(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id < ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdLessThanOrEqualTo(param int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id <= ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdIn(param []int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id in (?)", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdNotIn(param []int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id not in (?)", Param1: param})
	return ex
}

func (ex *DemoExample) AndIdBetween(param1 int32, param2 int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndIdNotBetween(param1 int32, param2 int32) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and id not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeIsNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType is null", noValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeIsNotNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeEqualTo(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType = ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeNotEqualTo(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType <> ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeGreaterThan(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType > ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeGreaterThanOrEqualTo(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType >= ?", Param1: param})
	return ex
}
func (ex *DemoExample) AndBitTypeLessThan(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType < ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeLessThanOrEqualTo(param []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType <= ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeIn(param [][]byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType in (?)", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeNotIn(param [][]byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType not in (?)", Param1: param})
	return ex
}

func (ex *DemoExample) AndBitTypeBetween(param1 []byte, param2 []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndBitTypeNotBetween(param1 []byte, param2 []byte) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and bitType not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeIsNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType is null", noValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeIsNotNull() *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType is not null", noValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeEqualTo(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType = ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotEqualTo(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType <> ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeGreaterThan(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType > ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeGreaterThanOrEqualTo(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType >= ?", Param1: param})
	return ex
}
func (ex *DemoExample) AndLongblobTypeLessThan(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType < ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeLessThanOrEqualTo(param string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType <= ?", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeIn(param []string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType in (?)", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotIn(param []string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType not in (?)", Param1: param})
	return ex
}

func (ex *DemoExample) AndLongblobTypeBetween(param1 string, param2 string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}

func (ex *DemoExample) AndLongblobTypeNotBetween(param1 string, param2 string) *DemoExample {
	ex.Criteria = append(ex.Criteria, struct {
		Fragment     string
		Param1       interface{}
		Param2       interface{}
		noValue      bool
		betweenValue bool
	}{Fragment: "and longblobType not between ? and ?", Param1: param1, Param2: param2, betweenValue: true})
	return ex
}
