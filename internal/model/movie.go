package model


type Movie struct {
	ID 		 int 	`json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Year     int64  `json:"year"`
	Director string `json:"director"`
}