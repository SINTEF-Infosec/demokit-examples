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

type State struct {
	VideoToLoad int
}

type VideoNode struct {
	*core.Node
	State State
}

func NewVideoNode() *VideoNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &VideoNode{
		Node: core.NewDefaultNodeWithVideo(),
		State: State{
			VideoToLoad: 0,
		},
	}
}

func (n *VideoNode) Configure() {
	n.SetEntryPoint(&core.Action{
		Name: "LoadVideoAndEmitLightSignal",
		Do:   n.LoadVideoAndEmitLightSignal,
	})

	n.OnEventDo("PLAY_VIDEO", &core.Action{
		Name: "PlayVideo",
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

	n.OnEventDo("I_MEDIA_ENDED", &core.Action{
		Name: "LoadNextVideo",
		Do:   n.LoadVideoAndEmitLightSignal,
		Then: &core.Action{
			Name: "AutomaticPlay",
			Do:   n.PlayVideo,
		},
	})
}

func (n *VideoNode) LoadVideoAndEmitLightSignal(_ *core.Event) {
	videosPaths := []string{
		"/home/pi/boo_scare_1.mp4",
		"/home/pi/Pumpkins.VOB",
	}

	n.Logger.Debugf("Loading video %d, from %s", n.State.VideoToLoad, videosPaths[n.State.VideoToLoad])

	if err := n.MediaController.LoadMediaFromPath(videosPaths[n.State.VideoToLoad]); err != nil {
		n.Logger.Errorf("could not load video: %v", err)
	}

	// We turn the light on or on
	if n.State.VideoToLoad == 0 {
		n.BroadcastEvent("LIGHT_ON", "")
	} else {
		n.BroadcastEvent("LIGHT_OFF", "")
	}

	n.State.VideoToLoad = (n.State.VideoToLoad + 1) % 2
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
