package main

import (
	"context"
	"encoding/json"
	"github.com/SINTEF-Infosec/demokit/core"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	node := NewPongNode()
	node.Configure()
	node.Start()
}

type State struct {
	IsStarted bool
}

type TemperatureNode struct {
	*core.Node
	State State
}

func NewPongNode() *TemperatureNode {
	logrus.SetLevel(logrus.DebugLevel)
	return &TemperatureNode{
		Node: core.NewDefaultNode(),
		State: State{
			IsStarted: false,
		},
	}
}

func (pn *TemperatureNode) Configure() {
	pn.ServeState(pn.State, true)

	temperatureCtx, cancel := context.WithCancel(context.Background())
	pn.Router.GET("/start", func(ctx *gin.Context) {
		if !pn.State.IsStarted {
			pn.Logger.Info("Starting temperature sensor")
			temperatureCtx, cancel = context.WithCancel(context.Background())
			go pn.ReadTemperatureAndBroadcast(temperatureCtx)
		}
	})
	pn.Router.GET("/stop", func(ctx *gin.Context) {
		if pn.State.IsStarted {
			pn.Logger.Info("Stopping temperature sensor")
			cancel()
		}
	})
}

type TemperatureEvent struct {
	Temperature float64
	Timestamp   time.Time
}

func (pn *TemperatureNode) ReadTemperatureAndBroadcast(ctx context.Context) {
	pn.State.IsStarted = true
	for {
		select {
		case <-ctx.Done():
			pn.State.IsStarted = false
			return
		default:
			// read temperature
			temperature, err := pn.Hardware.ReadTemperature()
			if err != nil {
				pn.Logger.Errorf("could not read temperature: %v", err)
				pn.State.IsStarted = false
				return
			}

			tempEvent := TemperatureEvent{
				Temperature: temperature,
				Timestamp: time.Now(),
			}

			payload, _ := json.Marshal(tempEvent)
			pn.Logger.Info("Broadcasting temperature data...")
			pn.BroadcastEvent("TEMPERATURE_DATA", string(payload))
			time.Sleep(2 * time.Second)
		}
	}
}
