package service

import (
	"cinema-ticket/internal/model"
	postgres "cinema-ticket/internal/repostiry"
	"errors"
	"log"
)

var (
	ErrInternal = errors.New("internal server error")
	ErrSessionNotFound = errors.New("session not found error")
	ErrMovieNotFound = errors.New("Movie not found error")
)

func CreateSession(input model.CreateSessionRequest) (model.Session, error) {
	db := postgres.GetDB()

	session := model.Session{
		RoomID: input.RoomID,
		Date: model.SessionDate{
			Start: input.Date.Start,
			End:   input.Date.End,
		},
	}

	if err := db.Create(&session).Error; err != nil {
		return model.Session{}, ErrInternal
	}

	log.Println(int(session.ID))
	return session, nil
}

func CreateShowTime(input model.CreateShowTimeRequest) (model.ShowTime, error) {
    db := postgres.GetDB()

    var session model.Session
    if err := db.First(&session, input.SessionID).Error; err != nil {
        return model.ShowTime{}, ErrSessionNotFound
    }
    var movie model.Movie
    if err := db.First(&movie, input.MovieID).Error; err != nil {
        return model.ShowTime{}, ErrMovieNotFound
    }

    showtime := model.ShowTime{
        MovieID:   input.MovieID,
        SessionID: input.SessionID,
        Price:     input.Price,
    }

    if err := db.Create(&showtime).Error; err != nil {
        return model.ShowTime{}, ErrInternal
    }

    return showtime, nil
}

func CreateMovie(input model.CreateMovieRequest) (model.Movie, error) {
	db := postgres.GetDB()

	var movie = model.Movie{
		Name:     input.Name,
		Genre:    input.Genre,
		Year:     input.Year,
		Director: input.Director,
	}

	err := db.Create(&movie).Error
	if err != nil {
		return model.Movie{}, ErrInternal
	}

	return movie, nil
}
