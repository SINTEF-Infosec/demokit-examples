package main

import (
	"encoding/json"
	"fmt"
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func main() {
	node := NewRingNode()
	node.Configure()
	node.Start()
}

type NextPayload struct {
	CurrentId int
}

type State struct {
	Id     int
	NbNode int
}

type RingNode struct {
	*core.Node
	State State
}

func NewRingNode() *RingNode {
	logrus.SetLevel(logrus.DebugLevel)
	id, err := strconv.Atoi(os.Getenv("NODE_ID"))
	if err != nil {
		logrus.Fatalf("could not get node ID: %v", err)
	}

	nbNode, err := strconv.Atoi(os.Getenv("NODE_NB"))
	if err != nil {
		logrus.Fatalf("could not get node number: %v", err)
	}

	return &RingNode{
		Node: core.NewDefaultNode(),
		State: State{
			Id:     id,
			NbNode: nbNode,
		},
	}
}

func (pn *RingNode) GetNextPayloadFromEvent(event *core.Event) (*NextPayload, error) {
	var payload NextPayload
	if event.Payload == "" {
		return nil, fmt.Errorf("empty payload")
	}

	if err := json.Unmarshal([]byte(event.Payload), &payload); err != nil {
		pn.Logger.Error("could not unmarshal payload: %v", err)
		return nil, fmt.Errorf("malformed payload")
	}

	return &payload, nil
}

func (pn *RingNode) IsItMyTurn(event *core.Event) bool {
	payload, err := pn.GetNextPayloadFromEvent(event)
	if err != nil {
		return false
	}

	if payload.CurrentId%pn.State.NbNode == pn.State.Id {
		return true
	}

	return false
}

func (pn *RingNode) Configure() {
	// Binding actions to events
	pn.OnEventDo("NEXT", &core.Action{
		Name:        "SendNext",
		Do:          pn.SendNext,
		DoCondition: pn.IsItMyTurn,
	})

	if pn.State.Id == 0 {
		pn.SetEntryPoint(&core.Action{
			Name: "SendNext",
			Do:   pn.SendNext,
		})
	}
}

func (pn *RingNode) SendNext(event *core.Event) {
	pn.Logger.Info("It is my turn ! Doing my stuff...")
	time.Sleep(time.Second)
	pn.Logger.Info("I'm done, passing to the next one...")

	// entry point payload
	nextPayload := NextPayload{
		CurrentId: 1,
	}

	// if not entry point
	if event != nil {
		payload, err := pn.GetNextPayloadFromEvent(event)
		if err != nil {
			return
		}

		nextPayload.CurrentId = payload.CurrentId + 1
	}

	strPayload, err := json.Marshal(nextPayload)
	if err != nil {
		pn.Logger.Errorf("error while marshaling next payload: %v", err)
	}

	pn.BroadcastEvent("NEXT", string(strPayload))
}
