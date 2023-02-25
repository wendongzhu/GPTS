package endinterface

import (
	"context"
	"github.com/wendongzhu/GPTS/backend/service/drives"
	"time"
)

func init() {
	// websocket
	ws := drives.WebSocket{}
	go ws.Start()

	time.Sleep(time.Second * 2)
	// RS485

	rs := drives.RS485{}
	go rs.Start()

}

type Backend struct {
	ctx context.Context
}

func StartBackend(ctx *context.Context) *Backend {
	return &Backend{}
}
