package errors

import (
	"github.com/asim/go-micro/v3/errors"
	"movie/src/share/config"
)

const (
	errorCodeFilmSuccess = 200
)

var (
	ErrorFilmFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeFilmSuccess,
	)
)
