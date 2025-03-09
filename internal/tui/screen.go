package tui

import "github.com/jackitaliano/wayfinder/internal/term"

type Screen struct {
    Width int
    Height int
    Buffer *Buffer
}

func NewScreen(borderChars BorderChars) Screen {
    width, height := term.GetTermSize()

    buffer := NewBuffer(0, 0, width, height, borderChars)
    buffer.Lines[0].Content = "hello"

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
