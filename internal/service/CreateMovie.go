package service

import (
	"cinema-ticket/internal/model"
	"cinema-ticket/internal/repostiry/indatabase"
)

func CreateMovie(input model.CreateMovieRequest) (model.Movie, error) {

	var movie = model.Movie{
		Name:     input.Name,
		Genre:    input.Genre,
		Year:     input.Year,
		Director: input.Director,
	}

	err := indatabase.CreateMovie(movie)
	if err != nil {
		return model.Movie{}, ErrInternal
	}

	return movie, nil
}
