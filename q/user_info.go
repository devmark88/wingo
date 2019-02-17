package q

import (
	"github.com/RichardKnop/machinery/v1"
	"gitlab.com/mt-api/wingo/model"
	"gitlab.com/mt-api/wingo/response"
)

// PublishUserInfo : Publish User corrector and ticket
func (q QueueManager) PublishUserInfo(c *model.UserInfo, srv *machinery.Server) error {
	pub := Pub{Server: srv}
	res := response.UserInfoPayload{}
	res.Type = response.UserInfoPayloadEnum
	res.Corrector = c.Correctors
	res.ID = c.ID
	res.Ticket = c.Tickets
	res.CanPlay = c.CanPlay

	e := pub.Publish(getUserTopic(c.ID), res)
	return e
}
