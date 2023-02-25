package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type PushController struct {
}

var uGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (pc *PushController) Router(engine *gin.Engine) {
	engine.GET("/ws/v1/test", pc.pushTest)
	engine.GET("/ws/v1/index", pc.pushIndex)

}

func (pc *PushController) pushTest(c *gin.Context) {
	client, _ := uGrader.Upgrade(c.Writer, c.Request, nil)
	defer client.Close()
	for {
		err := client.WriteJSON(1)
		if err != nil {
			log.Println(err)
			return
		}
	}

}

func (pc *PushController) pushIndex(c *gin.Context) {
	client, _ := uGrader.Upgrade(c.Writer, c.Request, nil)
	defer client.Close()
	for {
		err := client.WriteJSON(1)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Millisecond * 100)
	}

}
