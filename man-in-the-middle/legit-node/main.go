package main

import (
	"fmt"
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	node := NewDiscussionNode()
	node.Configure()
	node.Start()
}

type DiscussionNode struct {
	*core.Node
}

func NewDiscussionNode() *DiscussionNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &DiscussionNode{
		Node: core.NewDefaultNode(),
	}
}

func (pn *DiscussionNode) Configure() {
	// Binding actions to events
	pn.OnEventDo("DISCUSSION", &core.Action{
		Name: "Discussion",
		Do:   pn.Discussion,
	})
}

func (pn *DiscussionNode) Discussion(event *core.Event) {
	if event.Receiver == pn.Info.Name {
		time.Sleep(time.Second)
		pn.SendEventTo(event.Emitter, "DISCUSSION", fmt.Sprintf("{\"Message\": \"Hello %s\"}", event.Emitter))
	}
}
