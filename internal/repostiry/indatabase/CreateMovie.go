package indatabase

import (
	"cinema-ticket/internal/model"
	postgres "cinema-ticket/internal/repostiry"
)

func CreateMovie(movie model.Movie) error {
	db := postgres.GetDB()

	err := db.Create(&movie).Error
	if err != nil {
		return err
	}
	return nil
}
