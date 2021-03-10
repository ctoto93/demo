package sqlite

import (
	"github.com/ctoto93/demo/service"
	"github.com/jinzhu/gorm"
)

type repository struct {
	service.UnimplementedRepository
	db gorm.DB
}

func NewRepository() *repository {
	var db gorm.DB
	return &repository{db: db}

}
