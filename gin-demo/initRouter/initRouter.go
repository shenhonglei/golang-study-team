package initRouter

import (
	"gin-demo/handler/syslog"
	"gin-demo/handler/vm"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	articleRouter := router.Group("/api")
	{
		// 通过获取单条记录
		articleRouter.GET("/querybyid/:id", syslog.QuerybyID)
		// 根据VMname 查summary
		articleRouter.GET("/queryvmbyname/:name", vm.QueryVMbyName)

	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
