package main

import (
	"encoding/json"
	"fmt"
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
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

func (pn *DiscussionNode) handleEvent(event *core.Event) {
	// Ignoring events sent by this node
	if event.Emitter == pn.Info.Name {
		return
	}

	// Need to parse the payload
	type MsgPayload struct {
		Message string
	}

	var payload MsgPayload
	if event.Payload != "" {
		err := json.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			pn.Logger.Error("error while unmarshalling payload: %v", err)
		}
	}

	pn.BroadcastEvent("INTERCEPT",
		fmt.Sprintf("{\"From\": \"%s\", \"To\": \"%s\", \"MsgContent\": \"%s\",}", event.Emitter, event.Receiver, payload.Message))
}

func (pn *DiscussionNode) Configure() {
	// We set a new event handler
	pn.EventNetwork.SetReceivedEventCallback(pn.handleEvent)
}
