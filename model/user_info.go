package model

// UserInfo : UserInfo model
type UserInfo struct {
	ID         string `gorm:"primary_key"`
	Correctors uint
	Tickets    uint
	CanPlay    bool
}
