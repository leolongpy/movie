package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/cli/v2"
	"go.uber.org/zap"
	"movie/src/film-srv/db"
	"movie/src/film-srv/handler"
	"movie/src/share/config"
	"movie/src/share/pb"
	"movie/src/share/utils/log"
)

func main() {

	log.Init("film")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameFilm),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("comment-srv", "comment-srv is start ..."))
			// 注册redis
			//redisPool := share.NewRedisPool(3, 3, 1,300*time.Second,":6379","redis")
			// 先注册db
			db.Init(config.MysqlDSN)
			pb.RegisterFilmServiceExtHandler(service.Server(), handler.NewFilmServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("comment-srv", "comment-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("comment-srv服务启动失败 ...")
	}
}
