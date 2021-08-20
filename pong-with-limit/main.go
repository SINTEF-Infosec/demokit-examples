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

type State struct {
	Limit int
}

type PongNode struct {
	*core.Node
	State State
}

func NewPongNode() *PongNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &PongNode{
		Node: core.NewDefaultNode(),
		State: State{
			Limit: 5,
		},
	}
}

func (pn *PongNode) Configure() {
	// Binding actions to events
	pn.OnEventDo("PING", &core.Action{
		Name:        "SendPong",
		Do:          pn.SendPong,
		DoCondition: pn.CanSend,
	})
	sendPingAction := &core.Action{
		Name:        "SendPing",
		Do:          pn.SendPing,
		DoCondition: pn.CanSend,
	}
	pn.OnEventDo("PONG", sendPingAction)

	// Adding an entry point
	pn.SetEntryPoint(sendPingAction)

	// Serving the state
	pn.ServeState(&pn.State, true)
}

func (pn *PongNode) CanSend() bool {
	return pn.State.Limit > 0
}

func (pn *PongNode) SendPing(_ *core.Event) {
	time.Sleep(time.Second)
	pn.Logger.Info("Sending ping...")
	pn.BroadcastEvent("PING", "")
	pn.State.Limit--
}

func (pn *PongNode) SendPong(_ *core.Event) {
	time.Sleep(time.Second)
	pn.Logger.Info("Sending pong...")
	pn.BroadcastEvent("PONG", "")
	pn.State.Limit--
}
