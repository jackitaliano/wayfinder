package tui

import (
	"github.com/jackitaliano/wayfinder/internal/term/app"
	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
)

type Screen struct {
    Width int
    Height int
    Buffer *buffer.Buffer
}

func NewScreen(borderChars BorderChars) Screen {
    width, height := app.GetSize()

    buffer := buffer.NewBuffer(0, 0, width, height)

	return Screen{
        width,
        height,
        &buffer,
    }
}

func (s Screen) Draw() {
    s.Buffer.Draw()
}
