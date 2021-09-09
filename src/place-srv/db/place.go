package db

import (
	"database/sql"
	"movie/src/place-srv/entity"
)

func SelectPlaces() ([]*entity.Place, error) {
	places := []*entity.Place{}
	err := db.Select(&places, "SELECT * FROM `place` ")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return places, err
}
