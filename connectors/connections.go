package connectors

import (
	"github.com/jinzhu/gorm"
)

type Connections struct {
	Database *gorm.DB
}
