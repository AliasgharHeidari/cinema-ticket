package model


type Movie struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Year     int64  `json:"year"`
	Director string `json:"director"`
}

type CreateMovieRequest struct {
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Year     int64  `json:"year"`
	Director string `json:"director"`
}