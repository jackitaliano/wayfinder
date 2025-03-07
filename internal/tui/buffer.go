package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/jackitaliano/wayfinder/internal/term"
)

type Buffer struct {
    X int
    Y int
    Width int
    Height int
    Lines []Line
    borderChars BorderChars
    CursorLine int
    CursorX int
}

func NewBuffer(x int, y int, width int, height int, borderChars BorderChars) Buffer {
    lines := make([]Line, height)
    for i := 0; i < height; i++ {
        lines[i] = Line{"", "", strings.Repeat(" ", width)}
    }


    buffer := Buffer{
        x,
        y,
        width,
        height,
        lines,
        borderChars,
        0,
        0,
    }

    return buffer
}

func (b *Buffer) MoveCursorDown() {
    if b.CursorLine == b.Height - 1 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorLine += 1

    b.DrawCursor()
}

func (b *Buffer) MoveCursorUp() {
    if b.CursorLine == 0 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorLine -= 1

    b.DrawCursor()
}

func (b *Buffer) MoveCursorLeft() {
    if b.CursorX == 0 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorX -= 1

    b.DrawCursor()
}

func (b *Buffer) MoveCursorRight() {
    if b.CursorX == b.Width - 1 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorX += 1

    b.DrawCursor()
}

func (b Buffer) DrawCursor() {
    line := b.Lines[b.CursorLine]

    term.SetCursor(os.Stdin, b.X + 1, b.Y + b.CursorLine + 1)

    fmt.Fprint(os.Stdin, line.Fg, term.BlueBg, line.Content, term.Reset)

    term.SetCursor(os.Stdin, b.X + b.CursorX + 1, b.Y + b.CursorLine + 1)

    fmt.Fprint(os.Stdin, line.Fg, term.RedBg, string(line.Content[b.CursorX]), term.Reset)
}

func (b Buffer) DrawLine(lineNum int) {
    line := b.Lines[lineNum]

    term.SetCursor(os.Stdin, b.X + 1, b.Y + lineNum + 1)
    fmt.Fprint(os.Stdin, line.Fg, line.Bg, line.Content, term.Reset)
}

func (b Buffer) Draw() {
    for i := 0; i < b.CursorLine; i++ {
        b.DrawLine(i)
    }

    b.DrawCursor()

    for i := b.CursorLine + 1; i < b.Height; i++ {
        b.DrawLine(i)
    }
}

// func (b *Buffer) applyBorders() {
//
//     b.Grid[0][0].Char = b.borderChars.NW;
//     b.Grid[0][b.Width - 1].Char = b.borderChars.NE;
//     b.Grid[b.Height - 1][b.Width - 1].Char = b.borderChars.SE;
//     b.Grid[b.Height - 1][0].Char = b.borderChars.SW;
//
//     var i int
//     for i = 1; i < b.Width - 1; i++ {
//         b.Grid[0][i].Char = b.borderChars.N;
//         b.Grid[b.Height - 1][i].Char = b.borderChars.S;
//     }
//     for i = 1; i < b.Height - 1; i++ {
//         b.Grid[i][0].Char = b.borderChars.W;
//         b.Grid[i][b.Width - 1].Char = b.borderChars.E;
//     }
// }
