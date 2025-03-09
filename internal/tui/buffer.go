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
    CursorColor term.TermSpecifier
    CursorLine int
    CursorX int
    StatusLine StatusLine
}

func NewBuffer(x int, y int, width int, height int, borderChars BorderChars) Buffer {
    lines := make([]Line, height)
    for i := 0; i < height; i++ {
        lines[i] = Line{"", "", "", ""}
    }

    statusLine := StatusLine{"", "", NORMAL, 0, 0, "", 0, " ", " "}

    buffer := Buffer{
        x,
        y,
        width,
        height,
        lines,
        borderChars,
        term.RedBg,
        0,
        0,
        statusLine,
    }

    return buffer
}

func (b *Buffer) MoveCursorDown() {
    if b.CursorLine == b.Height - 1 {
        return
    }

    if b.CursorX > len(b.Lines[b.CursorLine + 1].Content) {
        b.CursorX = len(b.Lines[b.CursorLine + 1].Content)
    }

    b.DrawLine(b.CursorLine)

    b.CursorLine += 1

    b.DrawCursor()
}

func (b *Buffer) MoveCursorUp() {
    if b.CursorLine == 0 {
        return
    }

    if b.CursorX > len(b.Lines[b.CursorLine - 1].Content) {
        b.CursorX = len(b.Lines[b.CursorLine - 1].Content)
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
    if b.CursorX >= len(b.Lines[b.CursorLine].Content) - 1 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorX += 1

    b.DrawCursor()
}

func (b *Buffer) CursorNormalMode() {
    // b.CursorColor = term.RedBg
    b.StatusLine.Mode = NORMAL
    term.SetBlockCursor(os.Stdin)
    b.MoveCursorLeft()
    b.DrawCursor()
}

func (b *Buffer) CursorInsertMode() {
    // b.CursorColor = term.GreenBg
    term.SetLineCursor(os.Stdin)
    b.StatusLine.Mode = INSERT
    b.DrawCursor()
}

func (b *Buffer) CursorAppendMode() {

    if len(b.Lines[b.CursorLine].Content) > 0 {
        b.CursorX += 1
    }
    term.SetLineCursor(os.Stdin)
    b.StatusLine.Mode = INSERT
    b.DrawCursor()
}

func (b *Buffer) CursorHome() {
    b.CursorX = b.X
    b.DrawCursor()
}

func (b *Buffer) CursorEnd() {
    b.CursorX = len(b.Lines[b.CursorLine].Content) - 1
    b.DrawCursor()
}

func (b *Buffer) AppendLineBelow() {
    newLine := Line{"", "", "", ""}

    b.Lines = append(b.Lines[:b.CursorLine + 1], append([]Line{newLine}, b.Lines[b.CursorLine + 1:]...)...)

    b.CursorLine += 1
    b.CursorX = b.X

    b.Draw()
}

func (b *Buffer) AppendLineAbove() {
    newLine := Line{"", "", "", ""}

    b.Lines = append(b.Lines[:b.CursorLine], append([]Line{newLine}, b.Lines[b.CursorLine:]...)...)

    b.CursorX = b.X

    b.Draw()
}

func (b *Buffer) DeleteToEnd() {
    line := b.Lines[b.CursorLine]
    if len(line.Content) == 0 {
        return
    }

    b.Lines[b.CursorLine].Content = line.Content[:b.CursorX]

    b.CursorX -= 1

    if b.CursorX < 0 {
        b.CursorX = 0
    }

    b.DrawCursor()
}

func (b *Buffer) DeleteChar() {
    line := b.Lines[b.CursorLine]
    if len(line.Content) == 0 {
        return
    }

    b.Lines[b.CursorLine].Content = line.Content[:b.CursorX] + line.Content[b.CursorX + 1:]

    b.CursorX -= 1
    if b.CursorX < 0 {
        b.CursorX = 0
    }

    b.DrawCursor()
}

func (b *Buffer) Backspace() {
    if b.CursorX == 0 {
        if b.CursorLine == 0 {
            return
        }

        line := b.Lines[b.CursorLine]
        b.CursorX = len(b.Lines[b.CursorLine - 1].Content)

        if len(line.Content) == 0 {
            b.Lines = append(b.Lines[:b.CursorLine], b.Lines[b.CursorLine + 1:]...)

            b.CursorLine -= 1

            b.Draw()

            return
        }

        b.Lines[b.CursorLine - 1].Content += b.Lines[b.CursorLine].Content

        b.Lines = append(b.Lines[:b.CursorLine], b.Lines[b.CursorLine + 1:]...)

        b.CursorLine -= 1

        b.Draw()

        return
    }

    line := b.Lines[b.CursorLine]
    b.Lines[b.CursorLine].Content = line.Content[:b.CursorX - 1] + line.Content[b.CursorX:]
    b.CursorX -= 1
    b.DrawCursor()
}

func (b *Buffer) CarryLine() {
    line := b.Lines[b.CursorLine]

    remains := line.Content[:b.CursorX]
    carry := line.Content[b.CursorX:]

    newLine := Line{"", "", carry, ""}

    b.Lines[b.CursorLine].Content = remains

    b.Lines = append(b.Lines[:b.CursorLine + 1], append([]Line{newLine}, b.Lines[b.CursorLine + 1:]...)...)

    b.CursorLine += 1
    b.CursorX = b.X

    b.Draw()
}

func (b *Buffer) InsertChar(char string) {
    line := b.Lines[b.CursorLine]
    b.Lines[b.CursorLine].Content = line.Content[:b.CursorX] + char + line.Content[b.CursorX:]
    b.CursorX += 1
    b.DrawCursor()
}

func (b Buffer) DrawCursor() {
    line := b.Lines[b.CursorLine]
    line.Content = line.Content + strings.Repeat(" ", b.Width - len(line.Content) + 1)

    visibleContent := line.Content[:b.Width]

    term.SetCursor(os.Stdin, b.X + 1, b.Y + b.CursorLine + 1)

    fmt.Fprint(os.Stdin, line.Fg, term.GrayBg, visibleContent, term.Reset)

    // term.SetCursor(os.Stdin, b.X + b.CursorX + 1, b.Y + b.CursorLine + 1)
    //
    // fmt.Fprint(os.Stdin, line.Fg, b.CursorColor, string(visibleContent[b.CursorX]), term.Reset)

    b.DrawStatusLine()
    term.SetCursor(os.Stdin, b.X + b.CursorX + 1, b.Y + b.CursorLine + 1)
}

func (b Buffer) DrawLine(lineNum int) {
    line := b.Lines[lineNum]
    line.Content = line.Content + strings.Repeat(" ", b.Width - len(line.Content))

    visibleContent := line.Content[:b.Width]

    term.SetCursor(os.Stdin, b.X + 1, b.Y + lineNum + 1)
    fmt.Fprint(os.Stdin, line.Fg, line.Bg, visibleContent, term.Reset)

    term.SetCursor(os.Stdin, b.X + b.CursorX + 1, b.Y + b.CursorLine + 1)
}

func (b Buffer) DrawStatusLine() {


    b.StatusLine.Row = b.CursorLine
    b.StatusLine.Col = b.CursorX
    term.SetCursor(os.Stdin, b.X + 1, b.Y + b.Height)

    position := fmt.Sprintf("%v:%v", b.StatusLine.Row, b.StatusLine.Col)
    // input := fmt.Sprintf("%b:%v:%v", b.StatusLine.LastInput, b.StatusLine.LastInputKey, b.StatusLine.LastInputMap)
    input := fmt.Sprintf("%v:%v | ", b.StatusLine.LastInput, b.StatusLine.LastInputMap)

    padLen := b.Width - len(b.StatusLine.Mode) - len(b.StatusLine.Content) - len(position) - len(input)
    pad := strings.Repeat(" ", padLen)

    // fmt.Fprint(os.Stdin, b.StatusLine.Fg, b.StatusLine.Bg, b.StatusLine.Mode, " | ", b.StatusLine.Content, " | ", position)

    fmt.Fprint(os.Stdin, b.StatusLine.Fg, b.StatusLine.Bg, b.StatusLine.Mode, term.Reset)

    fmt.Fprint(os.Stdin, b.StatusLine.Fg, b.StatusLine.Bg, b.StatusLine.Content, pad, term.Reset)

    fmt.Fprint(os.Stdin, b.StatusLine.Fg, b.StatusLine.Bg, input, position, term.Reset)
}

func (b Buffer) Draw() {
    for i := 0; i < b.CursorLine; i++ {
        b.DrawLine(i)
    }

    b.DrawCursor()

    for i := b.CursorLine + 1; i < b.Height - 1; i++ {
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
