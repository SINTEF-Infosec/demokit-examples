package main

import (
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	node := NewVideoNode()
	node.Configure()
	node.Start()
}

type State struct {
	interruptMode bool
	mediaPosition float32
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
			interruptMode: false,
			mediaPosition: 0.0,
		},
	}
}

func (n *VideoNode) Configure() {
	n.SetEntryPoint(&core.Action{
		Name: "PlayPumpkins",
		Do:   n.PlayPumpkins,
	})

	n.OnEventDo("INTERRUPT_WITH_GHOSTS", &core.Action{
		Name: "InterruptWithGhosts",
		Do:   n.InterruptWithGhosts,
	})

	n.OnEventDo("I_MEDIA_ENDED", &core.Action{
		Name: "HandleInterruptEnded",
		Do:   n.HandleInterruptEnded,
	})

	n.Router.GET("/ghosts", func(ctx *gin.Context) {
		n.InterruptWithGhosts(nil)
		ctx.String(http.StatusOK, "PLAYING GHOSTS")
	})
}

func (n *VideoNode) PlayPumpkins(_ *core.Event) {
	if !n.State.interruptMode {
		n.Logger.Debug("Loading pumpkins video...")
		if err := n.MediaController.LoadMediaFromPath("/home/pi/Pumpkins.VOB"); err != nil {
			n.Logger.Errorf("could not load video: %v", err)
		}

		n.Logger.Info("Playing pumpkins video...")
		if err := n.MediaController.Play(); err != nil {
			n.Logger.Errorf("could not play video: %v", err)
		}
	} else {
		n.Logger.Info("not playing video, currently in interrupt mode")
	}
}

func (n *VideoNode) InterruptWithGhosts(_ *core.Event) {
	if !n.State.interruptMode {
		// switching to interrupt mode
		n.State.interruptMode = true

		n.Logger.Debug("Pausing pumpkins video...")
		if err := n.MediaController.Pause(); err != nil {
			n.Logger.Errorf("could not pause video: %v", err)
		}

		// Getting current position
		position, err := n.MediaController.GetCurrentMediaPosition()
		if err != nil {
			n.Logger.Errorf("could not get current media position: %v", err)
		}
		n.State.mediaPosition = position

		n.Logger.Info("Loading ghosts video...")
		if err := n.MediaController.LoadMediaFromPath("/home/pi/boo_scare_1.mp4"); err != nil {
			n.Logger.Errorf("could not load video: %v", err)
		}

		n.Logger.Info("Playing ghosts video...")
		if err := n.MediaController.Play(); err != nil {
			n.Logger.Errorf("could not play video: %v", err)
		}
	} else {
		n.Logger.Debug("Already interrupted by ghost!")
	}
}

func (n *VideoNode) HandleInterruptEnded(_ *core.Event) {
	if n.State.interruptMode {
		n.Logger.Info("Loading pumpkins video back...")
		if err := n.MediaController.LoadMediaFromPath("/home/pi/Pumpkins.VOB"); err != nil {
			n.Logger.Errorf("could not load video: %v", err)
		}

		n.Logger.Info("Playing pumpkins video back...")
		if err := n.MediaController.Play(); err != nil {
			n.Logger.Errorf("could not play video: %v", err)
		}

		// setting position for pumpkins video
		if err := n.MediaController.SetCurrentMediaPosition(n.State.mediaPosition); err != nil {
			n.Logger.Errorf("could not set media position: %v", err)
		}

		// switching back to normal mode
		n.State.interruptMode = false
		n.State.mediaPosition = 0.0
	} else {
		// The regular video ended, we just load it back
		n.PlayPumpkins(nil)
	}
}
