package contest

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"gitlab.com/mt-api/wingo/model"
)

type MetaRepository struct{}

func (r *MetaRepository) SaveMeta(m *model.ContestMeta, db *gorm.DB) {
	db.Create(m)
	fmt.Println(m)
	return
}
