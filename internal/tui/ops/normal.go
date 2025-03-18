package ops

import (
	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
)

type MoveOp struct {
    Key string
}

func (o MoveOp) Run(buffer *buffer.Buffer) error {
    if o.Key == "j" {
        buffer.MoveCursorDown()
    } else if o.Key == "k" {
        buffer.MoveCursorUp()
    } else if o.Key == "h" {
        buffer.MoveCursorLeft()
    } else if o.Key == "l" {
        buffer.MoveCursorRight()
    } else if o.Key == "0" {
        buffer.CursorHome()
    } else if o.Key == "$" {
        buffer.CursorEnd()
    }

    return nil
}

type ChangeModeOp struct {
    Key string
}

func (o ChangeModeOp) Run(buffer *buffer.Buffer) error {
    if o.Key == "ESC" {
        buffer.CursorNormalMode()

    } else if o.Key == "i" {
        buffer.CursorInsertMode()
    } else if o.Key == "I" {
        buffer.CursorHome()
        buffer.CursorInsertMode()
    } else if o.Key == "a" {
        buffer.CursorAppendMode()
    } else if o.Key == "A" {
        buffer.CursorEnd()
        buffer.CursorAppendMode()
    } else if o.Key == "o" {
        buffer.OpenLineBelow()
        buffer.CursorInsertMode()
    } else if o.Key == "O" {
        buffer.OpenLineAbove()
        buffer.CursorInsertMode()
    }

    return nil
}

type DeleteOp struct {
    Key string
}

func (o DeleteOp) Run(buffer *buffer.Buffer) error {
    if o.Key == "D" {
        buffer.DeleteToEnd()
    } else if o.Key == "x" {
        buffer.DeleteChar()
    }

    return nil
}

