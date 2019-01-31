package model

import (
	"context"
	"fmt"
	"time"
)

//MetaContest => Meta Contest Model
type MetaContest struct {
	id                         uint64
	appID                      string
	title                      string
	prize                      int
	beginTime                  time.Time
	duration                   int
	itemDuration               int
	neededCorrectors           int
	allowedCorrectorUsageTimes int
	allowCorrectTilQuestion    int
	neededTickets              int
}

func (meta MetaContest) validate() error {
	return nil
}
func (meta *MetaContest) save(ctx context.Context, manager Manager) error {
	var id uint64
	id = 0
	q := saveMetaQuery()
	err := manager.DB.QueryRowContext(ctx, q, meta.appID, meta.title, meta.prize, meta.beginTime, meta.duration, meta.itemDuration, meta.neededCorrectors, meta.allowedCorrectorUsageTimes, meta.allowCorrectTilQuestion, meta.neededTickets).Scan(&id)
	if err != nil {
		e := fmt.Errorf(fmt.Sprintf("Error in Meta Validation: %v", err))
		return e
	}
	meta.id = id
	return nil
}
