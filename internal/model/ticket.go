package model

import "time"

type Ticket struct {
	ID         int `gorm:"primaryKey"`
	ShowTimeID int `gorm:"uniqueIndex:idx_showtime_seat"`
	SeatID     int `gorm:"uniqueIndex:idx_showtime_seat"`
	UserID     int `gorm:"index"`
	IssueDate  time.Time
}