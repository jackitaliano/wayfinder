package input

import (
	"fmt"

	"github.com/jackitaliano/wayfinder/internal/tui"
)

type InsertHandler struct {
    keys map[byte]Op
}

func (i InsertHandler) Handle(context *tui.Context, input byte) error {
    op, handled := i.keys[input]

    if !handled {
        return &UnhandledKeyError{input}
    }

    err := op.Handler(context, op.Keys)

    if err != nil {
        return &HandlerError{
            fmt.Sprint("insert op handler error: ", err),
            context,
        }
    }

    return nil
}

func InsertStringOp(ctx *tui.Context, op string) error {
    ctx.ActiveBuffer.InsertChar(op)

    return nil
}

func NormalModeOp(ctx *tui.Context, op string) error {
    ctx.ActiveBuffer.CursorNormalMode()

    ctx.Mode = tui.NORMAL

    return nil
}

func CarryLineOp(ctx *tui.Context, op string) error {
    ctx.ActiveBuffer.CarryLine()

    return nil
}

func BackspaceOp(ctx *tui.Context, op string) error {
    ctx.ActiveBuffer.Backspace()

    return nil
}

func TabOp(ctx *tui.Context, op string) error {
    ctx.ActiveBuffer.InsertChar(" ")

    return nil
}
