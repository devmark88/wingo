package request

// UpdateUserInfoRequest : update user info request
type UpdateUserInfoRequest struct {
	Correctors uint `json:"corrector"`
	Tickets    uint `json:"ticket"`
}

// ReferralRequest : referral request
type ReferralRequest struct {
	ReferralUserID string `json:"referral"`
	ReferrerUserID string `json:"referrer"`
}

// NewUserRequest : new user request
type NewUserRequest struct {
	UserID string `json:"userId"`
}
