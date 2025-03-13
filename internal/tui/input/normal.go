package input

import (
	"fmt"

	"github.com/jackitaliano/wayfinder/internal/tui"
)


type NormalHandler struct {
    keys map[byte]Op
}

func (i NormalHandler) Handle(context *tui.Context, input byte) error {
    op, handled := i.keys[input]

    if !handled {
        return &UnhandledKeyError{input}
    }

    err := op.Handler(context, op.Keys)

    if err != nil {
        return &HandlerError{
            fmt.Sprint("normal op handler error: ", err),
            context,
        }
    }

    return nil
}

func MoveOp(ctx *tui.Context, op string) error {
    if op == "j" {
        ctx.ActiveBuffer.MoveCursorDown()
    } else if op == "k" {
        ctx.ActiveBuffer.MoveCursorUp()
    } else if op == "h" {
        ctx.ActiveBuffer.MoveCursorLeft()
    } else if op == "l" {
        ctx.ActiveBuffer.MoveCursorRight()
    } else if op == "0" {
        ctx.ActiveBuffer.CursorHome()
    } else if op == "$" {
        ctx.ActiveBuffer.CursorEnd()
    }

    return nil
}

func ChangeModeOp(ctx *tui.Context, op string) error {
    if op == "i" {
        ctx.ActiveBuffer.CursorInsertMode()
    } else if op == "I" {
        ctx.ActiveBuffer.CursorHome()
        ctx.ActiveBuffer.CursorInsertMode()
    } else if op == "a" {
        ctx.ActiveBuffer.CursorAppendMode()
    } else if op == "A" {
        ctx.ActiveBuffer.CursorEnd()
        ctx.ActiveBuffer.CursorAppendMode()
    }

    ctx.Mode = tui.INSERT

    return nil
}

func OpenLineOp(ctx *tui.Context, op string) error {
    if op == "o" {
        ctx.ActiveBuffer.OpenLineBelow()
        ctx.ActiveBuffer.CursorInsertMode()
    } else if op == "O" {
        ctx.ActiveBuffer.OpenLineAbove()
        ctx.ActiveBuffer.CursorInsertMode()
    }

    ctx.Mode = tui.INSERT

    return nil
}

func DeleteOp(ctx *tui.Context, op string) error {
    if op == "D" {
        ctx.ActiveBuffer.DeleteToEnd()
    } else if op == "x" {
        ctx.ActiveBuffer.DeleteChar()
    }

    return nil
}
