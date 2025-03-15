package events

import (
	"log"

	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
	"github.com/jackitaliano/wayfinder/internal/tui/ops"
)


type InputEvent struct {
    Priority EventPriority
    Op ops.Op
}

func (e InputEvent) Handle(buffer *buffer.Buffer) error {
    err := e.Op.Run(buffer)

    if err != nil {
        log.Printf("ERROR: Input Event: %v\n", err)
    }

    return err
}

func NormalInput(op ops.Op) InputEvent {
    return InputEvent{
        Priority: NORMAL,
        Op: op,
    }
}

func InsertInput(key string) InputEvent {
    return InputEvent{
        Priority: NORMAL,
        Op: ops.InsertStringOp{Key: key},
    }
}

func MoveOpInput(key string) InputEvent {
    return InputEvent{
        Priority: NORMAL,
        Op: ops.MoveOp{Key: key},
    }
}

func ChangeModeInput(key string) InputEvent {
    return InputEvent{
        Priority: NORMAL,
        Op: ops.ChangeModeOp{Key: key},
    }
}

func NoOpInput(key string) InputEvent {
    return InputEvent{
        Priority: NONE,
        Op: ops.NoOp{Key: key},
    }
}

func DeleteInput(key string) InputEvent {
    return InputEvent {
        Priority: NORMAL,
        Op: ops.DeleteOp{Key: key},
    }
}


