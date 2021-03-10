package db

import (
	"github.com/jinzhu/gorm"
)

type MySQL struct {
	db gorm.DB
}

func NewMySQLRepository() *MySQL {
	var db gorm.DB
	return &MySQL{db: db}

}

func (m *MySQL) GetStudent() {

}
