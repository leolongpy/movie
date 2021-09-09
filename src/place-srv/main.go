package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/cli/v2"
	"go.uber.org/zap"
	"movie/src/place-srv/db"
	"movie/src/place-srv/handler"
	"movie/src/share/config"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
)

func main() {
	log.Init("place")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNamePlace),
		micro.Version("latest"),
	)
	service.Init(

		//micro.Action(func(context *cli2.Context) error {

		//}),
		micro.Action(func(context *cli.Context) error {
			logger.Info("Info", zap.Any("place-srv", "place-srv is start"))

			db.Init(config.MysqlDSN)
			pb.RegisterPlaceServiceExtHandler(service.Server(), handler.NewPlaceServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),

		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("place-srv", "place-srv is stop..."))
			return nil
		}),

		micro.AfterStart(func() error {
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		logger.Panic("place-srv服务启动失败")
	}
}
