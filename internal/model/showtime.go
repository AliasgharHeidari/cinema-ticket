package model

import "time"

type ShowTime struct {
	ID        int           `json:"id"`
	MovieID   int           `json:"movieId"`
	SessionID int           `json:"sessionId"`
	Price     float64       `json:"price"`
	Seats     ShowTimeSeats `json:"seats"`
}

type ShowTimeSeat struct {
	Row       int       `json:"row"`
	Column    int       `json:"column"`
	IssueDate time.Time `json:"issueDate"`
}

type ShowTimeSeats struct {
	Reserved []ShowTimeSeat `json:"reserved"`
	Sold     []ShowTimeSeat `json:"sold"`
}

type CreateShowTimeRequest struct {
	SessionID int
	MovieID   int
	Price     float64
}
