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
    // buffer.StatusLine.Content = fmt.Sprintf("width: %v, height: %v", width, height)

	return Screen{
        width,
        height,
        &buffer,
    }
}

func (s Screen) Draw() {
    s.Buffer.Draw()
}

func (s Screen) MoveCursorDown() {
    s.Buffer.MoveCursorDown()
}

func (s Screen) MoveCursorUp() {
    s.Buffer.MoveCursorUp()
}

func (s Screen) MoveCursorLeft() {
    s.Buffer.MoveCursorLeft()
}

func (s Screen) MoveCursorRight() {
    s.Buffer.MoveCursorRight()
}

func (s Screen) InsertChar(char string) {
    s.Buffer.InsertChar(char)
}
