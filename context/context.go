package context

import "gitlab.com/mt-api/wingo/connectors"

type AppContext struct {
	Connections *connectors.Connections
}
