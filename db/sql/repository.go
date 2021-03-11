package sql

import (
	"github.com/ctoto93/demo"
	"gorm.io/gorm"
)

type repository struct {
	demo.UnimplementedRepository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
