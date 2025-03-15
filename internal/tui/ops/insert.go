package ops

import (
	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
)


type InsertStringOp struct {
    Key string
}

func (o InsertStringOp) Run(buffer *buffer.Buffer) error {
    buffer.InsertChar(o.Key)

    return nil
}

type NormalModeOp struct {
    Key string
}

func (o NormalModeOp) Run(buffer *buffer.Buffer) error {
    buffer.CursorNormalMode()

    // ctx.Mode = tui.NORMAL

    return nil
}

type CarryLineOp struct {
    Key string
}

func (o CarryLineOp) Run(buffer *buffer.Buffer) error {
    buffer.CarryLine()

    return nil
}

type BackspaceOp struct {
    Key string
}

func (o BackspaceOp) Run(buffer *buffer.Buffer) error {
    buffer.Backspace()

    return nil
}

type TabOp struct {
    Key string
}

func (o TabOp) Run(buffer *buffer.Buffer) error {
    buffer.InsertChar(" ")

    return nil
}
