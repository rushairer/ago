package bootstrap

import (
	"log"
	"net/http"

	apiControllers "PACKAGENAME/api/controllers/api"
	"PACKAGENAME/core/auth"
	"PACKAGENAME/core/config"
	"PACKAGENAME/core/databases"

	"github.com/gin-gonic/gin"
)

// 在这里初始化各种服务对象
func SetupServer(server *gin.Engine) {

	log.Println("connecting databases...")

	mysql, err := databases.NewMysql(config.MysqlDSN)
	if err != nil {
		log.Panicln(err)
	}

	// redis := databases.NewRedis(config.RedisDSN)

	// routers
	testGroup := server.Group("/test")
	{
		testGroup.GET(
			"/alive",
			func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			},
		)
	}

	viewerService := auth.NewViewerService(mysql)
	apiGroup := server.Group("/api")
	{
		userController := apiControllers.NewUserController(viewerService)
		apiGroup.GET(
			"/viewer",
			userController.Viewer,
		)
	}

}
