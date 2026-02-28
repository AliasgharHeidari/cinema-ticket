package service

import (
	"cinema-ticket/internal/model"
	"cinema-ticket/internal/repostiry/indatabase"
	"errors"
)

var (
	ErrInternal        = errors.New("internal server error")
	ErrSessionNotFound = errors.New("session not found error")
	ErrMovieNotFound   = errors.New("Movie not found error")
)

func CreateShowTime(input model.CreateShowTimeRequest) (model.ShowTime, error) {

	err := indatabase.SearchForSession(input)
	if err != nil {
		return model.ShowTime{}, ErrSessionNotFound
	}

	err = indatabase.SearchForMovie(input)
	if err != nil {
		return model.ShowTime{}, ErrMovieNotFound
	}

	showtime := model.ShowTime{
		MovieID:   input.MovieID,
		SessionID: input.SessionID,
		Price:     input.Price,
	}

	err = indatabase.CreateShowTime(showtime)

	return showtime, nil
}
