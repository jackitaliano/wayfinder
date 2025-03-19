package events

import (
	"log"

	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
	"github.com/jackitaliano/wayfinder/internal/tui/context"
)

type EventPriority int

const (
    CRITICAL EventPriority = 0
    TIME_SENSITIVE EventPriority = 1
    NORMAL EventPriority = 5
    LAZY EventPriority = 10
    NONE EventPriority = 100
)

type EventHandler struct {
    activeBuffer *buffer.Buffer
    inputEvents []InputEvent
    bufferEvents []BufferEvent
    otherEvents []Event
    drawEvents []DrawEvent
}

func NewEventHandler(buffer *buffer.Buffer) *EventHandler {
    return &EventHandler{
        buffer,
        []InputEvent{},
        []BufferEvent{},
        []Event{},
        []DrawEvent{},
    }
}

type Event interface {
    Handle(*buffer.Buffer) error
}

func (e *EventHandler) SetBuffer(buffer *buffer.Buffer) {
    e.activeBuffer = buffer
}

func (e *EventHandler) PostEvent(event Event) {
    if e.activeBuffer == nil {


    }

    switch ev := event.(type) {
    case InputEvent:
        e.inputEvents = append(e.inputEvents, ev)
    case BufferEvent:
        e.bufferEvents = append(e.bufferEvents, ev)
    case DrawEvent:
        e.drawEvents = append(e.drawEvents, ev)
    default:
        e.otherEvents = append(e.otherEvents, ev)
    }
}

func (e *EventHandler) HandlePendingEvents(ctx *context.Context) {
    e.HandlePendingInputEvents(ctx)
    e.HandlePendingBufferEvents(ctx)
    e.HandlePendingOtherEvents(ctx)
    e.HandlePendingDrawEvents(ctx)
}

func (e *EventHandler) clearEvents() {
    e.clearInputEvents()
    e.clearBufferEvents()
    e.clearOtherEvents()
    e.clearDrawEvents()
}

func (e *EventHandler) HandlePendingInputEvents(ctx *context.Context) {
    for _, event := range e.inputEvents {
        err := event.Handle(e.activeBuffer)

        if err != nil {
            log.Printf("ERROR: %v\n", err)
        }
    }

    e.clearInputEvents()
}

func (e *EventHandler) HandlePendingBufferEvents(ctx *context.Context) {
    for _, event := range e.bufferEvents {
        err := event.Handle(e.activeBuffer)

        if err != nil {
            log.Printf("ERROR: %v\n", err)
        }
    }

    e.clearBufferEvents()
}

func (e *EventHandler) HandlePendingOtherEvents(ctx *context.Context) {
    for _, event := range e.otherEvents {
        err := event.Handle(e.activeBuffer)

        if err != nil {
            log.Printf("ERROR: %v\n", err)
        }
    }

    e.clearOtherEvents()
}

func (e *EventHandler) HandlePendingDrawEvents(ctx *context.Context) {
    for _, event := range e.bufferEvents {
        err := event.Handle(e.activeBuffer)

        if err != nil {
            log.Printf("ERROR: %v\n", err)
        }
    }

    e.clearDrawEvents()
}

func (e *EventHandler) clearInputEvents() {
    e.inputEvents = []InputEvent{}
}

func (e *EventHandler) clearBufferEvents() {
    e.bufferEvents = []BufferEvent{}
}

func (e *EventHandler) clearOtherEvents() {
    e.otherEvents = []Event{}
}

func (e *EventHandler) clearDrawEvents() {
    e.drawEvents = []DrawEvent{}
}

type LogEvent struct {
    Level string
    Msg string
}

func (e LogEvent) Handle(buffer *buffer.Buffer) error {

    buffer.StatusPrintf("%v: %v", e.Level, e.Msg)

    return nil
}
