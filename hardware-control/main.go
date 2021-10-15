package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
	"strings"
)

func main() {
	node := NewControllerNode()
	node.Configure()
	node.Start()
}

type State struct {
	LatestCommands []string
}

type ControllerNode struct {
	*core.Node
	State State
}

func NewControllerNode() *ControllerNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &ControllerNode{
<<<<<<< HEAD
		Node: core.NewDefaultRaspberryPiNode(),
=======
		Node: core.NewDefaultRaspberryPiNode(true),
>>>>>>> 46d8092 (Add hardware control example.)
		State: State{
			LatestCommands: make([]string, 8),
		},
	}
}

func (pn *ControllerNode) Configure() {
	broadCastEventAction := &core.Action{
		Name: "BroadcastEvent",
		Do:   pn.BroadcastInputEvent,
	}

	// Binding actions to events
	pn.OnEventDo("I_UP_PRESSED", broadCastEventAction)
	pn.OnEventDo("I_DOWN_PRESSED", broadCastEventAction)
	pn.OnEventDo("I_LEFT_PRESSED", broadCastEventAction)
	pn.OnEventDo("I_RIGHT_PRESSED", broadCastEventAction)
}

func (pn *ControllerNode) BroadcastInputEvent(event *core.Event) {
	command := strings.ReplaceAll(event.Name, "I_", "")
	command = strings.ReplaceAll(command, "_PRESSED", "")
	pn.AddCommand(command)
	pn.BroadcastEvent(strings.ReplaceAll(event.Name, "I_", ""), "")
	if pn.CheckKonamiCode() {
		pn.BroadcastEvent("KONAMI_CODE", "")
	}
}

func (pn *ControllerNode) AddCommand(command string) {
<<<<<<< HEAD
	// Shift previous command
=======
	// Shift previous commands
>>>>>>> 46d8092 (Add hardware control example.)
	for k := 7; k > 0; k-- {
		pn.State.LatestCommands[k] = pn.State.LatestCommands[k-1]
	}
	pn.State.LatestCommands[0] = command
}

func (pn *ControllerNode) CheckKonamiCode() bool {
	return pn.State.LatestCommands[7] == "UP" &&
		pn.State.LatestCommands[6] == "DOWN" &&
		pn.State.LatestCommands[5] == "UP" &&
		pn.State.LatestCommands[4] == "DOWN" &&
		pn.State.LatestCommands[3] == "LEFT" &&
		pn.State.LatestCommands[2] == "RIGHT" &&
		pn.State.LatestCommands[1] == "LEFT" &&
		pn.State.LatestCommands[0] == "RIGHT"
}
