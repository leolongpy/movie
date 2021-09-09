package handler

import (
	"context"
	"go.uber.org/zap"
	"movie/src/place-srv/db"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
)

type PlaceServiceExtHandler struct {
	logger *zap.Logger
}

func NewPlaceServiceExtHandler() *PlaceServiceExtHandler {
	return &PlaceServiceExtHandler{
		logger: log.Instance(),
	}
}

//HotCitiesByCinema(context.Context, *HotCitiesByCinemaReq, *HotCitiesByCinemaReq) error
func (p *PlaceServiceExtHandler) HotCitiesByCinema(ctx context.Context, req *pb.HotCitiesByCinemaReq, res *pb.HotCitiesByCinemaRep) error {
	places, err := db.SelectPlaces()
	if err != nil {
		p.logger.Error("err", zap.Any("places", err))
		return err
	}
	placesPB := []*pb.Place{}
	for _, place := range places {
		placePB := place.ToProtoDBHotPlayMovies()
		placesPB = append(placesPB, placePB)
	}
	res.P = placesPB
	return nil
}
