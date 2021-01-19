package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DemoRepo DemoRepository
}

func NewRepo(db *sqlx.DB) Repository {
	return Repository{
		DemoRepo: NewDemoRepo(db),
	}
}

func NewDemoRepo(db *sqlx.DB) DemoRepository {
	return &demoRepo{
		db: db,
	}
}

type DemoRepository interface {
	Insert(*Demo) error
	Update(*Demo) error
	DeleteByPrimaryKey(int32) error
	SelectByPrimaryKey(int32) (*Demo, error)
}
