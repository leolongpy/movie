package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/cli/v2"
	"go.uber.org/zap"
	"movie/src/share/config"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
	"movie/src/user-srv/db"
	"movie/src/user-srv/handler"
)

func main() {
	log.Init("user")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameUser),
		micro.Version("latest"),
	)
	service.Init(

		//micro.Action(func(context *cli2.Context) error {

		//}),
		micro.Action(func(context *cli.Context) error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is start"))

			db.Init(config.MysqlDSN)
			pb.RegisterUserServiceExtHandler(service.Server(), handler.NewUserServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),

		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is stop..."))
			return nil
		}),

		micro.AfterStart(func() error {
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败")
	}
}
