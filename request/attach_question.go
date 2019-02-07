package request

import (
	"encoding/json"
	"fmt"

	"gitlab.com/mt-api/wingo/helpers"

	"gitlab.com/mt-api/wingo/messages"

	"gitlab.com/mt-api/wingo/model"
)

type AttachQuestion struct {
	ContestID      uint64     `json:"contestId"`
	CorrectAnswers []byte     `json:"correctAnswers"`
	Questions      []Question `json:"questions"`
}

type Question struct {
	Order   byte                    `json:"order"`
	Body    string                  `json:"text"`
	Options []Option                `json:"options"`
	Level   model.QuestionLevelEnum `json:"level"`
}

type Option struct {
	Body string `json:"text"`
	Hit  uint   `json:"hit"`
}

func (a *AttachQuestion) ToModel() (*model.Contest, error) {
	m := model.Contest{Questions: []model.Question{}}
	m.ContestMetaID = a.ContestID
	m.CorrectAnswersIndices = helpers.ByteArrayToString(a.CorrectAnswers, ",")

	for _, q := range a.Questions {
		mq := new(model.Question)
		mq.Body = q.Body
		mq.ContestID = m.ID
		mq.Level = q.Level
		if len(q.Level) == 0 {
			mq.Level = model.Unknown
		}
		mq.Order = q.Order
		b, err := json.Marshal(q.Options)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf(messages.MAPPING_ERROR, "request.AttachQuestion", "model.Contest", err))
		}
		mq.Answers = string(b)
		m.Questions = append(m.Questions, *mq)
	}
	return &m, nil

}