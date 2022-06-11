package eventbus

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"
)

var ctx = context.Background()

type EventBus struct {
	Bus *bus.Bus
}

// New ...
func New() (*EventBus, error) {
	b, err := newBus()
	if err != nil {
		return nil, err
	}
	serv := &EventBus{
		Bus: b,
	}
	return serv, nil
}

// EmitString emits an event to the eventbus
func (inst *EventBus) EmitString(topicName string, data string) error {
	ctx = context.WithValue(ctx, bus.CtxKeyTxID, "")
	err := inst.Bus.Emit(ctx, topicName, data)
	if err != nil {
		return err
	}
	return err
}

// Emit emits an event to the eventbus
func (inst *EventBus) Emit(topicName string, data interface{}) error {
	err := inst.Bus.Emit(ctx, topicName, data)
	if err != nil {
		return err
	}
	return err
}

// RegisterTopic registers a topic for publishing
func (inst *EventBus) RegisterTopic(ds string) {
	inst.Bus.RegisterTopics(fmt.Sprintf("%s", ds))
}

// RegisterTopicParent registers a topic for publishing
func (inst *EventBus) RegisterTopicParent(parent string, child string) {
	topic := fmt.Sprintf("%s.%s", parent, child)
	inst.Bus.RegisterTopics(topic)
}

// UnregisterTopic removes a topic from consumer
func (inst *EventBus) UnregisterTopic(topic string) {
	inst.Bus.DeregisterTopics(topic)
}

// UnregisterTopicChild removes a topic from consumer
func (inst *EventBus) UnregisterTopicChild(parent string, child string) {
	topic := fmt.Sprintf("%s.%s", parent, child)
	inst.Bus.DeregisterTopics(topic)
}

// UnsubscribeHandler removes handler
func (inst *EventBus) UnsubscribeHandler(id string) {
	inst.Bus.DeregisterHandler(id)
}

func newBus() (*bus.Bus, error) {
	// configure id generator
	node := uint64(1)
	initialTime := uint64(1577865600000) // set 2020-01-01 PST as initial time
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		return nil, err
	}
	// init an id generator
	var idGenerator bus.Next = m.Next
	// create a new eventbus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		return nil, err
	}
	return b, nil
}
