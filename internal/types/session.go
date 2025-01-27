package types

import "github.com/google/uuid"

type Session struct {
	RefreshToken string
	AccessToken  string
	Username     string
	UserId       uuid.UUID
}
