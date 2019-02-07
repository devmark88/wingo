package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
	"gitlab.com/mt-api/wingo/repository/contest"
)

type Connections struct {
	DB *gorm.DB
}

func (cn *Connections) AddMeta(m *model.ContestMeta) {
	r := contest.MetaRepository{}
	r.SaveMeta(m, cn.DB)
}
