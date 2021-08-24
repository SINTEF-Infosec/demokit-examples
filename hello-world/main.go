package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	node := NewHelloNode()
	node.Configure()
	node.Start()
}

type HelloNode struct {
	*core.Node
}

func NewHelloNode() *HelloNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &HelloNode{
		Node: core.NewDefaultNode(),
	}
}

func (n *HelloNode) Configure() {
	n.SetEntryPoint(&core.Action{
		Name: "HelloWorld",
		Do: n.HelloWorld,
	})
}

func (n *HelloNode) HelloWorld(_ *core.Event) {
	for {
		n.Logger.Info("Broadcasting hello world...")
		n.BroadcastEvent("HELLO_WORLD", "")
		time.Sleep(2 * time.Second)
	}
}
