package handler

import (
	"context"
	"go.uber.org/zap"
	"movie/src/share/errors"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
	"movie/src/user-srv/db"
)

type UserServiceExtHandler struct {
	logger *zap.Logger
}

func NewUserServiceExtHandler() *UserServiceExtHandler {
	return &UserServiceExtHandler{
		logger: log.Instance(),
	}
}

// RegistAccount 用户注册
func (u *UserServiceExtHandler) RegistAccount(ctx context.Context, req *pb.RegistAccountReq, rsp *pb.RegistAccountRsp) error {
	userName := req.UserName
	password := req.Password
	email := req.Email
	user, err := db.SelectUserByEmail(email)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user != nil {
		return errors.ErrorUserAlready
	}
	err = db.InsertUser(userName, password, email)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	return nil
}

// LoginAccount 用户登录
func (u *UserServiceExtHandler) LoginAccount(ctx context.Context, req *pb.LoginAccountReq, rsp *pb.LoginAccountRsp) error {
	email := req.Email
	password := req.Password
	user, err := db.SelectUserByPasswordName(email, password)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user == nil {
		return errors.ErrorUserLoginFailed
	}
	rsp.Email = user.Email
	rsp.Phone = user.Phone
	rsp.UserID = user.UserId
	rsp.UserName = user.UserName
	return nil
}

func (u *UserServiceExtHandler) ResetAccount(ctx context.Context, req *pb.ResetAccountReq, rsp *pb.ResetAccountRsp) error {
	return nil
}

func (u *UserServiceExtHandler) WantScore(ctx context.Context, req *pb.WantScoreReq, rsp *pb.WantScoreRsp) error {

	orderNum, err := db.SelectOrderByUidMid(req.MovieId, req.UserID)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if orderNum == 0 {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorScoreForbid
	}
	err = db.UpdateOrderScore(req.MovieId, req.Score)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	return nil
}

func (u *UserServiceExtHandler) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileReq, rsp *pb.UpdateUserProfileRsp) error {

	userEmail := req.UserEmail
	userName := req.UserName
	userPhone := req.UserPhone
	userID := req.UserID
	if userEmail != "" {
		err := db.UpdateUserEmailProfile(userEmail, userID)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userName != "" {
		err := db.UpdateUserNameProfile(userName, userID)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userPhone != "" {
		err := db.UpdateUserPhoneProfile(userPhone, userID)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	return nil
}
