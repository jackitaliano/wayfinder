package input

import (
	"fmt"

	"github.com/jackitaliano/wayfinder/internal/tui"
	"github.com/jackitaliano/wayfinder/internal/tui/events"
	"github.com/jackitaliano/wayfinder/internal/tui/ops"
)

type InputHandler struct {
    ctx *tui.Context
    eventHandler *events.EventHandler
    normalKeys map[byte]events.InputEvent
    insertKeys map[byte]events.InputEvent
}

func NewInputHandler(ctx *tui.Context, eventHandler *events.EventHandler) InputHandler {
    normalKeys := defineNormalOps()
    insertKeys := defineInsertOps()

    return InputHandler{
        ctx,
        eventHandler,
        normalKeys,
        insertKeys,
    }
}

func (i *InputHandler) HandleKey(inputKey byte) error {
    i.eventHandler.PostEvent(events.LogEvent{Type: "INFO", Message: fmt.Sprintf("ctx.mode: %v", i.ctx.Mode)})
    if i.ctx.Mode == tui.NORMAL {
        event, handled := i.normalKeys[inputKey]

        if !handled {
            return &UnhandledKeyError{inputKey}
        }

        switch event.Op.(type) {
        case ops.ChangeModeOp:
            i.ctx.Mode = tui.INSERT
        }

        i.eventHandler.PostEvent(event)

        return nil
    }

    if i.ctx.Mode == tui.INSERT {
        event, handled := i.insertKeys[inputKey]

        if !handled {
            return &UnhandledKeyError{inputKey}
        }

        switch event.Op.(type) {
        case ops.ChangeModeOp:
            i.ctx.Mode = tui.NORMAL
        }

        i.eventHandler.PostEvent(event)

        return nil
    }

    i.eventHandler.PostEvent(events.LogEvent{Type: "ERROR", Message: "mode not selected"})

    return nil
}

type UnhandledKeyError struct {
    Key byte
}

func (e UnhandledKeyError) Error() string {
    return fmt.Sprintf("UnhandledKeyError: '%v'", e.Key)
}

