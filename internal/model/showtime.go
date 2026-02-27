package model


type ShowTime struct {
	ID        int           `gorm:"primaryKey" json:"id"`
	MovieID   int           `gorm:"index" json:"movieId"`
	SessionID int           `gorm:"index" json:"sessionId"`
	Price     float64       `json:"price"`
}

type CreateShowTimeRequest struct {
	SessionID int     `json:"sessionId"`
	MovieID   int     `json:"movieId"`
	Price     float64 `json:"price"`
}