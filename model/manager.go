package model

import (
	"context"
	"database/sql"
	"fmt"
)

//Manager => Manager for everything which related to database
type Manager struct {
	DB   *sql.DB
	Meta *MetaContest
}

//SaveMeta => Add meta to database and return meta with id
func (mm Manager) SaveMeta(ctx context.Context, meta MetaContest) (MetaContest, error) {
	err := meta.validate()
	if err != nil {
		e := fmt.Errorf(fmt.Sprintf("Error in Meta Validation: %v", err))
		return meta, e
	}
	err = meta.save(ctx, mm)
	if err != nil {
		return meta, err
	}
	return meta, nil

}
