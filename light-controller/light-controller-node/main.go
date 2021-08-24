package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	node := NewLightControllerNode()
	node.Configure()
	node.Start()
}

type LightControllerNode struct {
	*core.Node
}

func NewLightControllerNode() *LightControllerNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &LightControllerNode{
		Node: core.NewDefaultNode(),
	}
}

func (n *LightControllerNode) Configure() {
	n.Router.GET("/turn_light_on", func(ctx *gin.Context) {
		n.BroadcastEvent("LIGHT_ON", "")
	})
	n.Router.GET("/turn_light_off", func(ctx *gin.Context) {
		n.BroadcastEvent("LIGHT_OFF", "")
	})
}
