package sql

import (
	"github.com/ctoto93/demo/service"
	"gorm.io/gorm"
)

type repository struct {
	service.UnimplementedRepository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
