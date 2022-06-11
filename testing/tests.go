package main

import (
	"context"
	"fmt"
	"github.com/NubeIO/lib-bus/eventbus"
	"github.com/mustafaturan/bus/v3"
	"os"
	"os/signal"
)

const allTopic = "*"
const topic = "test"

type test struct {
	EventBus *eventbus.EventBus
}

type name struct {
	name string
	age  string
}

func main() {
	s, _ := eventbus.New()
	s.RegisterTopic(topic)
	test := &test{
		EventBus: s,
	}
	test.registerSubscriber()
	err := s.Emit(topic, &name{name: "i love coding on the weekend"})
	fmt.Println(err)
	keepRunning()

}

func (inst *test) registerSubscriber() {
	handler := bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			go func() {
				switch e.Topic {
				case topic:
					payload, ok := e.Data.(*name)
					msg := fmt.Sprintf("topic: %s msg: %s", topic, payload.name)
					fmt.Println(msg)
					if !ok {
						return
					}
				}
			}()
		},
		Matcher: topic,
	}
	inst.EventBus.Bus.RegisterHandler(allTopic, handler)
}

func keepRunning() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
