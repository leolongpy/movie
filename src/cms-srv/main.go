package main

import (
	"movie/src/cms-srv/db"
	"movie/src/cms-srv/handler"
	"movie/src/share/config"
	"movie/src/share/pb"
	"movie/src/share/utils/log"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/cli/v2"
	"go.uber.org/zap"
)

func main() {

	log.Init("cms")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameCMS),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("cms-srv", "cms-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterCMSServiceExtHandler(service.Server(), handler.NewCMSServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("cms-srv", "cms-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("cms-srv服务启动失败 ...")
	}
}
