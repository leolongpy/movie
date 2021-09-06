package errors

import (
	"github.com/asim/go-micro/v3/errors"
	"movie/src/share/config"
)

const (
	errorCodeCommentSuccess = 200
)

var (
	ErrorCommentFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeCommentSuccess,
	)
)
