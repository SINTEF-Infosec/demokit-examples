package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
)

func main() {
	node := NewLightNode()
	node.Configure()
	node.Start()
}

type LightNode struct {
	*core.Node
}

func NewLightNode() *LightNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &LightNode{
		Node: core.NewDefaultNode(),
	}
}

func (n *LightNode) Configure() {
	n.OnEventDo("LIGHT_ON", &core.Action{
		Name: "TurnLightOn",
		Do: n.TurnLightOn,
	})
	n.OnEventDo("LIGHT_OFF", &core.Action{
		Name: "TurnLightOff",
		Do: n.TurnLightOff,
	})
}

func (n *LightNode) TurnLightOn(_ *core.Event) {
	n.Logger.Info("Turning light on...")
	if err := n.Hardware.LightOn(); err != nil {
		n.Logger.Errorf("could not turn light on: %v", err)
	}
}

func (n *LightNode) TurnLightOff(_ *core.Event) {
	n.Logger.Info("Turning light off...")
	if err := n.Hardware.LightOff(); err != nil {
		n.Logger.Errorf("could not turn light off: %v", err)
	}
}
