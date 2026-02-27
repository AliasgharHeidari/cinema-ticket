package model

type Room struct {
	Number int64  `json:"number"`
	ID     int    `json:"id"`
	Seat   []Seat `json:"seat"`
}

type Seat struct {
	ID     int   `json:"id"`
	Number int64 `json:"number"`
	Row    int64 `json:"row"`
}
