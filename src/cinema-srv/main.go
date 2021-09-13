package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/cli/v2"
	"go.uber.org/zap"
	"movie/src/cinema-srv/db"
	"movie/src/cinema-srv/handler"
	"movie/src/share/config"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
)

func main() {

	log.Init("cinema")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameCinema),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("cinema-srv", "cinema-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterCinemaServiceExtHandler(service.Server(), handler.NewCinemaServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("cinema-srv", "cinema-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("cinema-srv服务启动失败 ...")
	}
}
