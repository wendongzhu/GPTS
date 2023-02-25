package drives

import (
	"github.com/gin-gonic/gin"
	"github.com/wendongzhu/GPTS/backend/service/middlewares"
	"github.com/wendongzhu/GPTS/backend/service/router"
)

type WebSocket struct {
}

func (ws *WebSocket) Start() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	//设置全局跨域访问-通过中间件
	cors := middlewares.Cors{}
	app.Use(cors.CorsMiddleWares())

	r := router.Router{}
	r.RegisterRouter(app)
	err := app.Run("0.0.0.0:8090")
	if err != nil {
		return
	}

}
