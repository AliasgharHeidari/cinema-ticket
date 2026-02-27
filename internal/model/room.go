package model

type Room struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Number int64  `gorm:"uniqueIndex" json:"number"`
	Seat   []Seat `gorm:"foreignKey:RoomID" json:"seat"`
}
type Seat struct {
	ID     int   `gorm:"primaryKey" json:"id"`
	RoomID int   `gorm:"index" json:"roomId"`
	Column int `json:"number"`
	Row    int64 `json:"row"`
}

type CreateRoomRequest struct {
	Number int64
	Seat []Seat	
}