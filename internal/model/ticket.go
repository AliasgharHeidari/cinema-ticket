package model

type Ticket struct {
	ID         int          `json:"id"`
	ShowTimeID int          `json:"showTimeId"`
	UserID     int          `json:"userId"`
	Seat       ShowTimeSeat `json:"Seat"`
}
