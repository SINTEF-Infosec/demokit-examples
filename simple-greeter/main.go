package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
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
	pn.OnEventDo("GREET", &core.Action{
		Name: "GreetSender",
		Do:   pn.GreetSender,
	})
}

func (pn *PongNode) GreetSender(event *core.Event) {
	if event.Receiver == pn.Info.Name {
		pn.SendEventTo(event.Emitter, "GREET", "")
	}
}
