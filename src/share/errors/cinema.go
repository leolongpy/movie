package errors

import (
	"github.com/asim/go-micro/v3/errors"
	"movie/src/share/config"
)

const (
	errorCodeCinemaSuccess = 200
)

var (
	ErrorCinemaFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeCinemaSuccess,
	)
	ErrorCinemaNotFound = errors.New(
		config.ServiceNameUser, "找不到对应的影院", errorCodeCinemaSuccess,
	)
)
