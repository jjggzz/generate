package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	CustomerRepo CustomerRepository
	DemoRepo     DemoRepository
}

func NewRepo(db *sqlx.DB) Repository {
	return Repository{
		CustomerRepo: NewCustomerRepo(db),
		DemoRepo:     NewDemoRepo(db),
	}
}

func NewCustomerRepo(db *sqlx.DB) CustomerRepository {
	return &customerRepo{
		db: db,
	}
}

func NewDemoRepo(db *sqlx.DB) DemoRepository {
	return &demoRepo{
		db: db,
	}
}

type CustomerRepository interface {
	Count() (int64, error)
	Insert(*Customer) (int64, error)
	UpdateByPrimaryKey(int64, *Customer) (int64, error)
	DeleteByPrimaryKey(int64) (int64, error)
	SelectByPrimaryKey(int64) (*Customer, error)
	DeleteByExample(example *CustomerExample) (int64, error)
	SelectByExample(*CustomerExample) ([]*Customer, error)
}

type DemoRepository interface {
	Count() (int64, error)
	Insert(*Demo) (int64, error)
	UpdateByPrimaryKey(int32, *Demo) (int64, error)
	DeleteByPrimaryKey(int32) (int64, error)
	SelectByPrimaryKey(int32) (*Demo, error)
	DeleteByExample(example *DemoExample) (int64, error)
	SelectByExample(*DemoExample) ([]*Demo, error)
}
