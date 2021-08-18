package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
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

	// Adding an entry point
	pn.SetEntryPoint(sendPingAction)
}

func (pn *PongNode) SendPing() {
	time.Sleep(time.Second)
	pn.Logger.Info("Sending ping...")
	pn.BroadcastEvent("PING", "")
}

func (pn *PongNode) SendPong() {
	time.Sleep(time.Second)
	pn.Logger.Info("Sending pong...")
	pn.BroadcastEvent("PONG", "")
}
