package model

import "time"

type Session struct {
	ID     int         `gorm:"primaryKey" json:"id"`
	RoomID int         `gorm:"index" json:"roomId"`
	Date   SessionDate `gorm:"embedded" json:"date"`
}

type SessionDate struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type CreateSessionRequest struct {
	RoomID int         `json:"roomId"`
	Date   SessionDate `json:"date"`
}

func (s *Session) Duration() time.Duration {
	return s.Date.End.Sub(s.Date.Start)

}
