package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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

type Repository struct {
	CustomerRepo CustomerRepository
	DemoRepo     DemoRepository
}

type CustomerRepository interface {
	Insert(*Customer) error
	Update(*Customer) error
	DeleteByPrimaryKey(int64) error
	SelectByPrimaryKey(int64) (*Customer, error)
}

type DemoRepository interface {
	Insert(*Demo) error
	Update(*Demo) error
	DeleteByPrimaryKey(int32) error
	SelectByPrimaryKey(int32) (*Demo, error)
}
