package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	node := NewPongNode()
	node.Configure()
	node.Start()
}

type PongNode struct {
	*core.Node
}

func NewPongNode() *PongNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &PongNode{
		Node: core.NewDefaultNode(),
	}
}

func (pn *PongNode) Configure() {
	// Binding actions to events
	pn.OnEventDo("PING", &core.Action{
		Name: "SendPong",
		Do:   pn.SendPong,
	})
	sendPingAction := &core.Action{
		Name: "SendPing",
		Do:   pn.SendPing,
	}
	pn.OnEventDo("PONG", sendPingAction)

	// Adding endpoints to manually send ping and pong events
	pn.Router.GET("/ping", func(c *gin.Context) {
		pn.SendPing(nil)
		c.String(http.StatusOK, "PING Event sent")
	})
	pn.Router.GET("/pong", func(c *gin.Context) {
		pn.SendPong(nil)
		c.String(http.StatusOK, "PONG Event sent")
	})
}

func (pn *PongNode) SendPing(_ *core.Event) {
	time.Sleep(time.Second)
	pn.Logger.Info("Sending ping...")
	pn.BroadcastEvent("PING", "")
}

func (pn *PongNode) SendPong(_ *core.Event) {
	time.Sleep(time.Second)
	pn.Logger.Info("Sending pong...")
	pn.BroadcastEvent("PONG", "")
}
