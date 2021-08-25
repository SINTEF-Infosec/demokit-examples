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
		Node: core.NewDefaultNode(),
	}
}

func (n *VideoNode) Configure() {
	n.SetEntryPoint(&core.Action{
		Name: "LoadVideo",
		Do: n.LoadVideo,
	})

	n.OnEventDo("PLAY_VIDEO", &core.Action{
		Name: "PlayVideo",
		Do: n.PlayVideo,
	})

	n.OnEventDo("PAUSE_VIDEO", &core.Action{
		Name: "PauseVideo",
		Do: n.PauseVideo,
	})

	n.OnEventDo("STOP_VIDEO", &core.Action{
		Name: "StopVideo",
		Do: n.StopVideo,
	})
}

func (n *VideoNode) LoadVideo(_ *core.Event) {
	if err := n.MediaController.LoadMediaFromPath("/home/pi/boo_scare_1.mp4"); err != nil {
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
