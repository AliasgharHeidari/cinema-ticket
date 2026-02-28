package indatabase

import (
	"cinema-ticket/internal/model"
	postgres "cinema-ticket/internal/repostiry"
)

func CreateSession(session model.Session) error {

	db := postgres.GetDB()
	if err := db.Create(&session).Error; err != nil {
		return err
	}

	return nil
}
