package input

import (
	"fmt"

	"github.com/jackitaliano/wayfinder/internal/tui/context"
	"github.com/jackitaliano/wayfinder/internal/tui/events"
	"github.com/jackitaliano/wayfinder/internal/tui/ops"
)

type InputHandler struct {
    mode string
    eventHandler *events.EventHandler
    normalKeys map[byte]events.InputEvent
    insertKeys map[byte]events.InputEvent
}

func NewInputHandler(eventHandler *events.EventHandler) InputHandler {
    normalKeys := defineNormalOps()
    insertKeys := defineInsertOps()

    return InputHandler{
        "NORMAL",
        eventHandler,
        normalKeys,
        insertKeys,
    }
}

func (i *InputHandler) HandleKey(ctx *context.Context, inputKey byte) error {
    if i.mode == "NORMAL" {
        event, handled := i.normalKeys[inputKey]

        if !handled {
            return &UnhandledKeyError{inputKey}
        }

        switch event.Op.(type) {
        case ops.ChangeModeOp:
            i.mode = "INSERT"
        }

        i.eventHandler.PostEvent(event)

        return nil
    }

    if i.mode == "INSERT" {
        event, handled := i.insertKeys[inputKey]

        if !handled {
            return &UnhandledKeyError{inputKey}
        }

        switch event.Op.(type) {
        case ops.ChangeModeOp:
            i.mode = "NORMAL"
        }

        i.eventHandler.PostEvent(event)

        return nil
    }

    i.eventHandler.PostEvent(events.LogEvent{Level: "ERROR", Msg: "mode not selected"})

    return nil
}

type UnhandledKeyError struct {
    Key byte
}

func (e UnhandledKeyError) Error() string {
    return fmt.Sprintf("UnhandledKeyError: '%v'", e.Key)
}

