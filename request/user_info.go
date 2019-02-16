package request

// UpdateUserInfoRequest : update user info request
type UpdateUserInfoRequest struct {
	Correctors uint `json:"corrector"`
	Tickets    uint `json:"ticket"`
}
