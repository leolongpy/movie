package handler

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"movie/src/order-srv/db"
	"movie/src/order-srv/entity"
	"movie/src/share/errors"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
	"strconv"
	"time"
)

type OrderServiceExtHandler struct {
	logger *zap.Logger
}

func NewOrderServiceExtHandler() *OrderServiceExtHandler {
	return &OrderServiceExtHandler{
		logger: log.Instance(),
	}
}

func (o *OrderServiceExtHandler) WantTicket(ctx context.Context, req *pb.WantTicketReq, rsp *pb.WantTicketRsp) error {
	err := db.InsertWantSeeRecord(req.FilmId, req.UserId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	return nil
}

func (o *OrderServiceExtHandler) Ticket(ctx context.Context, req *pb.TicketReq, rsp *pb.TicketRsp) error {
	price, err := db.SelectFilmPrice(req.FilmId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	orderNum := time.Now().Unix()
	err = db.InsertOrder(strconv.Itoa(int(orderNum)), price, req.MhId, req.UserId, req.FilmId, req.X, req.Y, req.StartTime, req.EndTime)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	rsp.OrderNumD = orderNum
	return nil
}

func (o *OrderServiceExtHandler) PayOrder(ctx context.Context, req *pb.PayOrderReq, rsp *pb.PayOrderRsp) error {
	err := db.UpdateOrderStatus(req.OrderNum, req.UserId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	o.logger.Debug("debug", zap.Any("req.Phone", req.Phone))
	o.logger.Debug("debug", zap.Any("req.UserId", req.UserId))
	err = db.UpdateUserPhone(req.UserId, req.Phone)
	if err != nil {
		o.logger.Error("err", zap.Any("UpdateUserPhone", err))
		return errors.ErrorOrderFailed
	}
	return nil
}

func (o *OrderServiceExtHandler) UndoOrder(ctx context.Context, req *pb.UndoOrderReq, rsp *pb.UndoOrderRsp) error {
	return nil
}

func (o *OrderServiceExtHandler) LookOrders(ctx context.Context, req *pb.LookOrdersReq, rsp *pb.LookOrdersRsp) error {

	o.logger.Debug("debug", zap.Any("debug", req.UserId))
	userId := req.UserId
	orders, err := db.SelectOrderNumMovieIdMHIdStartTime(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	movieTicketsPB := []*pb.MovieTicket{}

	for _, order := range orders {
		cinemafilm, err := db.SelectFilmNameCinemaName(order.MhId, order.MovieId)
		if err != nil {
			o.logger.Error("err", zap.Any("order", err))
			return errors.ErrorOrderFailed
		}
		movieTicketPB := &pb.MovieTicket{
			FilmName:  cinemafilm.FilmName,
			Cinema:    cinemafilm.CinemaName,
			StartTime: order.StartTime,
			OrderNum:  order.OrderNum,
		}
		movieTicketsPB = append(movieTicketsPB, movieTicketPB)
	}
	rsp.MovieTickets = movieTicketsPB
	return nil
}

func (o *OrderServiceExtHandler) LookAlreadyOrders(ctx context.Context, req *pb.LookAlreadyOrdersReq, rsp *pb.LookAlreadyOrdersRsp) error {
	userId := req.UserId
	orders, err := db.SelectLookAlreadyOrders(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	var oneNoComment int64 = 0
	movies := []*pb.AlreadyMovie{}
	for _, order := range orders {
		actorNames := []string{}
		film, err := db.SelectFilmDetail(order.MovieId)
		if err != nil {
			o.logger.Error("err", zap.Any("order", err))
			return errors.ErrorOrderFailed
		}
		actors, err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			o.logger.Error("err", zap.Any("order", err))
			return errors.ErrorOrderFailed
		}
		for _, actor := range actors {
			actorNames = append(actorNames, actor.ActorName)
		}
		movie := &pb.AlreadyMovie{
			FilmImg:    film.Img,
			FilmName:   film.TitleCn,
			Time:       fmt.Sprintf("%d-%d-%d", film.RYear, film.RMonth, film.RDay),
			Director:   film.FilmDirector,
			ActorNames: actorNames,
			OrderNum:   order.OrderNum,
		}
		movies = append(movies, movie)
		if order.OrderScore == -1 {
			oneNoComment = oneNoComment + 1
		}
	}
	rsp.Movies = movies
	rsp.OneNoComment = oneNoComment
	rsp.TotalMovie = int64(len(orders))
	return nil
}

func (o *OrderServiceExtHandler) OrderComment(ctx context.Context, req *pb.OrderCommentReq, rsp *pb.OrderCommentRsp) error {
	userId := req.UserId
	score := req.Score
	orderNum := req.OrderNum
	content := req.CommentContent
	order, err := db.SelectOrderScore(orderNum, userId)
	if err != nil {
		o.logger.Error("err", zap.Any("SelectOrderScore", err))
		return errors.ErrorOrderFailed
	}
	if order == nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	if order.OrderScore != -1 {
		return errors.ErrorOrderAlreadyScore
	}
	err = db.UpdateOrderScore(score, userId, orderNum)
	if err != nil {
		o.logger.Error("err", zap.Any("UpdateOrderScore", err))
		return errors.ErrorOrderFailed
	}
	err = db.UpdateFilmScore(order.MovieId, score)
	if err != nil {
		o.logger.Error("err", zap.Any("UpdateFilmScore", err))
		return errors.ErrorOrderFailed
	}
	user, err := db.SelectUserNameByUserId(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("SelectUserNameByUserId", err))
		return errors.ErrorOrderFailed
	}
	comment := entity.Comment{
		FilmId:   order.MovieId,
		Title:    strconv.Itoa(int(score)),
		Content:  content,
		NickName: user.UserName,
		UserId:   userId,
	}
	err = db.InsertComment(&comment)
	if err != nil {
		o.logger.Error("err", zap.Any("InsertComment", err))
		return errors.ErrorOrderFailed
	}
	return nil
}

func (o *OrderServiceExtHandler) GetOrderMessage(ctx context.Context, req *pb.GetOrderMessageReq, rsp *pb.GetOrderMessageRsp) error {
	orderNum := req.OrderNum
	userId := req.UserId
	order, err := db.SelectOrderMessage(orderNum, userId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	film, err := db.SelectFilmMessage(order.MovieId)
	if err != nil {
		o.logger.Error("err", zap.Any("film", err))
		return errors.ErrorOrderFailed
	}
	hall, err := db.SelectMHNameMHId(order.MhId)
	if err != nil {
		o.logger.Error("err", zap.Any("hall", err))
		return errors.ErrorOrderFailed
	}

	ciname, err := db.SelectCinemaByCid(hall.CinemaId)
	if err != nil {
		o.logger.Error("err", zap.Any("cinema", err))
		return errors.ErrorOrderFailed
	}
	user, err := db.SelectUserPhoneByUserId(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("user", err))
		return errors.ErrorOrderFailed
	}
	ticketDetailPB := &pb.TicketDetail{
		FilmName:      film.TitleCn,
		FilmImg:       film.Img,
		StartTime:     order.StartTime,
		EndTime:       order.EndTime,
		CinemaName:    ciname.CinemaName,
		MhName:        hall.MhName,
		Seat:          fmt.Sprintf("%d排%d座", order.OrderX, order.OrderY),
		OrderNum:      orderNum,
		CinemaAddress: ciname.CinemaAdd,
		Price:         order.OrderPrice,
		CreateAt:      order.CreateAt,
		Phone:         user.Phone,
		CinemaPhone:   ciname.CinemaPhone,
	}
	rsp.Ticketdetail = ticketDetailPB
	return nil
}
