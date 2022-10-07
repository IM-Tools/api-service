package bootstrap

import (
	"im-services/internal/config"
	"im-services/internal/middleware"
	router2 "im-services/internal/router"
	"im-services/internal/service/client"
	"im-services/internal/service/queue/nsq_queue"
	"im-services/pkg/coroutine_poll"
	"im-services/pkg/logger"
	"im-services/pkg/model"
	"im-services/pkg/nsq"
	"im-services/pkg/redis"
	"im-services/server"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

// Start 启动服务方法
func Start() {

	r := gin.Default()

	go client.ImManager.Start()

	r.Use(middleware.Recover)

	setRoute(r)

	gin.SetMode(config.Conf.Server.Mode)

	server.StartGrpc()

	_ = r.Run(config.Conf.Server.Listen)
}

// LoadConfiguration 加载连接池
func LoadConfiguration() {

	setUpLogger()

	model.InitDb()

	redis.InitClient()

	coroutine_poll.ConnectPool()

	_ = nsq.InitNewProducerPoll()
	// todo 消费逻辑可以单独抽离
	go nsq_queue.ConsumersPrivateMessageInit()
	go nsq_queue.ConsumersPrivateMessageInit()

}

// 初始化日志方法
func setUpLogger() {
	logger.InitLogger(
		config.Conf.Log.FileName,
		config.Conf.Log.MaxSize,
		config.Conf.Log.MaxBackup,
		config.Conf.Log.MaxAge,
		config.Conf.Log.Compress,
		config.Conf.Log.Type,
		config.Conf.Log.Level,
	)
}

// 注册路由方法
func setRoute(r *gin.Engine) {
	router2.RegisterApiRoutes(r)
	router2.RegisterWsRouters(r)
}
