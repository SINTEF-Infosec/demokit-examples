package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/sirupsen/logrus"
)

func main() {
	node := NewVideoNode()
	node.Configure()
	node.Start()
}

type VideoNode struct {
	*core.Node
}

func NewVideoNode() *VideoNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &VideoNode{
		Node: core.NewDefaultNodeWithVideo(),
	}
}

func (n *VideoNode) Configure() {
	n.SetEntryPoint(&core.Action{
		Name: "LoadBooVideo",
		Do:   n.LoadBooVideo,
	})

	n.OnEventDo("PLAY_VIDEO", &core.Action{
		Name: "PlayPumpkins",
		Do:   n.PlayVideo,
	})

	n.OnEventDo("PAUSE_VIDEO", &core.Action{
		Name: "PauseVideo",
		Do:   n.PauseVideo,
	})

	n.OnEventDo("STOP_VIDEO", &core.Action{
		Name: "StopVideo",
		Do:   n.StopVideo,
	})

	n.OnEventDo("LOAD_PUMPKINS", &core.Action{
		Name: "LoadPumpkins",
		Do:   n.LoadPumpkinsVideo,
	})

	n.OnEventDo("LOAD_BOO", &core.Action{
		Name: "LoadBoo",
		Do:   n.LoadBooVideo,
	})
}

func (n *VideoNode) LoadBooVideo(_ *core.Event) {
	if err := n.MediaController.LoadMediaFromPath("/home/pi/boo_scare_1.mp4"); err != nil {
		n.Logger.Errorf("could not load video: %v", err)
	}
}

func (n *VideoNode) LoadPumpkinsVideo(_ *core.Event) {
	if err := n.MediaController.LoadMediaFromPath("/home/pi/Pumpkins.VOB"); err != nil {
		n.Logger.Errorf("could not load video: %v", err)
	}
}

func (n *VideoNode) PlayVideo(_ *core.Event) {
	n.Logger.Info("Playing video...")
	if err := n.MediaController.Play(); err != nil {
		n.Logger.Errorf("could not play video: %v", err)
	}
}

func (n *VideoNode) StopVideo(_ *core.Event) {
	n.Logger.Info("Stopping video...")
	if err := n.MediaController.Stop(); err != nil {
		n.Logger.Errorf("could not stop video: %v", err)
	}
}

func (n *VideoNode) PauseVideo(_ *core.Event) {
	n.Logger.Info("Pausing video...")
	if err := n.MediaController.Pause(); err != nil {
		n.Logger.Errorf("could not pause video: %v", err)
	}
}
