package model

import "time"

type Session struct {
	ID     int         `json:"id"`
	RoomID int         `json:"roomId"`
	Date   SessionDate `json:"date"`
}

type SessionDate struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type CreateSessionRequest struct {
	RoomId int 
	Date   SessionDate
}




func (s *Session) Duration() time.Duration {
	return s.Date.End.Sub(s.Date.Start)

}

