package events

import (
	// "log"

	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
)

type DrawEvent struct {
    // BufferId int
    LineNumber int
    StartCol int
    Content string
}

func (e DrawEvent) Handle(buffer *buffer.Buffer) error {
    // err := e.Op.Run(buffer)
    //
    // if err != nil {
    //     log.Printf("ERROR: Input Event: %v\n", err)
    // }

    return nil
}
