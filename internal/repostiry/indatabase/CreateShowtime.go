package indatabase

import (
	"cinema-ticket/internal/model"
	postgres "cinema-ticket/internal/repostiry"
)

func SearchForSession(input model.CreateShowTimeRequest) error {
	db := postgres.GetDB()
	var session model.Session
	if err := db.First(&session, input.SessionID).Error; err != nil {
		return err
	}
	return nil
}

func SearchForMovie(input model.CreateShowTimeRequest) error {
	db := postgres.GetDB()
	var movie model.Movie
	if err := db.First(&movie, input.MovieID).Error; err != nil {
		return err
	}
	return nil
}

func CreateShowTime(showtime model.ShowTime) error {
	db := postgres.GetDB()
	if err := db.Create(&showtime).Error; err != nil {
		return err
	}
	return nil
}
