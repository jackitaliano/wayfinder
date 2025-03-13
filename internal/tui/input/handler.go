package input

import (
	"fmt"

	"github.com/jackitaliano/wayfinder/internal/tui"
)

type InputHandler struct {
    normalHandler ModeHandler
    insertHandler ModeHandler
}

func NewInputHandler() InputHandler {

    normalHandler := &NormalHandler{defineNormalOps()}
    insertHandler := &InsertHandler{defineInsertOps()}

    return InputHandler{
        normalHandler,
        insertHandler,
    }

}

func (i *InputHandler) HandleKey(ctx *tui.Context, inputKey byte) {
    if ctx.Mode == tui.NORMAL {
        err := i.normalHandler.Handle(ctx, inputKey)

        if err != nil {
            ctx.ActiveBuffer.StatusPrint(err)
        }
        return
    }

    if ctx.Mode == tui.INSERT {
        err := i.insertHandler.Handle(ctx, inputKey)

        if err != nil {
            ctx.ActiveBuffer.StatusPrint(err)
        }
        return
    }


    ctx.ActiveBuffer.StatusPrint("Error: mode not selected")

}

type OpHandler func(ctx *tui.Context, op string) error

func NoOp(ctx *tui.Context, op string) error { return nil }

type ModeHandler interface {
    Handle(*tui.Context, byte) error
}

type Op struct {
    Keys string
    Handler OpHandler
    IsChorded bool
}

type HandlerError struct {
    Message string
    Context *tui.Context
}

func (e HandlerError) Error() string {
    return fmt.Sprintf("HandlerError: %v\nContext: %v", e.Message, e.Context )
}

type UnhandledKeyError struct {
    Key byte
}

func (e UnhandledKeyError) Error() string {
    return fmt.Sprintf("UnhandledKeyError: '%v'", e.Key)
}

