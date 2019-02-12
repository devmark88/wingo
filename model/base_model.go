package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Base : Base of all models
type Base struct {
	gorm.Model
}
