package service

import (
	"cinema-ticket/internal/model"
	"cinema-ticket/internal/repostiry/indatabase"
	"log"
)

func CreateSession(input model.CreateSessionRequest) (model.Session, error) {

	session := model.Session{
		RoomID: input.RoomID,
		Date: model.SessionDate{
			Start: input.Date.Start,
			End:   input.Date.End,
		},
	}

	err := indatabase.CreateSession(session)
	if err != nil {
		return model.Session{}, ErrInternal
	}

	log.Println(int(session.ID))
	return session, nil
}
