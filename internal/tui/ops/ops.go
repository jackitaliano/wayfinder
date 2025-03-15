package ops

import "github.com/jackitaliano/wayfinder/internal/tui/buffer"

type Op interface {
    Run(*buffer.Buffer) error
}

type NoOp struct {
    Key string
}

func (o NoOp) Run(buffer *buffer.Buffer) error {
    return nil
}
