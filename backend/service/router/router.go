package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wendongzhu/GPTS/backend/service/controller"
)

type Router struct {
}

func (r *Router) RegisterRouter(engine *gin.Engine) {

	new(controller.PushController).Router(engine)

}
