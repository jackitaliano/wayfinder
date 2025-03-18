package buffer

import (
    "os"
    "strings"
    "fmt"

	"github.com/jackitaliano/wayfinder/internal/term/cursor"
	"github.com/jackitaliano/wayfinder/internal/term/color"
)

func (b Buffer) DrawCursor() {
    pad := strings.Repeat(" ", b.Width - len(b.CurrentLine.Content) - len(b.CurrentLine.Gutter))

    visibleContent := b.CurrentLine.Content[:min(len(b.CurrentLine.Content), b.Width + len(b.CurrentLine.Gutter))]

    cursor.SetPos(os.Stdout, b.TermLine + b.CursorLine + 1, b.TermCol + 1)

    fmt.Fprint(os.Stdout, b.CurrentLine.Fg, color.GrayBg, b.CurrentLine.Gutter, visibleContent, pad, color.ResetFormat)

    b.DrawStatusLine()
    cursor.SetPos(os.Stdout, b.TermLine + b.CursorLine + 1, b.TermCol + b.CursorCol + len(b.CurrentLine.Gutter) + 1)
}

func (b Buffer) DrawLine(lineNum int) {
    line := b.Lines[lineNum]

    pad := strings.Repeat(" ", b.Width - len(line.Content) - len(line.Gutter))

    visibleContent := line.Content[:min(len(line.Content), b.Width + len(b.CurrentLine.Gutter))]

    cursor.SetPos(os.Stdout, b.TermLine + lineNum + 1, b.TermCol + 1)
    fmt.Fprint(os.Stdout, line.Fg, line.Bg, line.Gutter, visibleContent, pad, color.ResetFormat)

    cursor.SetPos(os.Stdout, b.TermLine + b.CursorLine + 1, b.TermCol + b.CursorCol + len(b.CurrentLine.Gutter) + 1)
}

func (b Buffer) DrawStatusLine() {
    b.StatusLine.Row = b.CursorLine
    b.StatusLine.Col = b.CursorCol
    cursor.SetPos(os.Stdout, b.TermLine + b.Height, b.TermCol + 1)

    position := fmt.Sprintf("%v:%v", b.StatusLine.Row, b.StatusLine.Col)
    input := fmt.Sprintf("%v:%v", b.StatusLine.LastInput, b.StatusLine.LastInputMap)

    padLen := b.Width - len(b.StatusLine.Mode) - 1 - len(b.StatusLine.Content) - len(position) - len(input) - 2 * len(b.StatusLine.Separator)
    pad := strings.Repeat(" ", padLen)

    fmt.Fprint(os.Stdout, b.StatusLine.Fg, b.StatusLine.Bg, " ", b.StatusLine.Mode, color.ResetFormat)

    fmt.Fprint(os.Stdout, b.StatusLine.Fg, b.StatusLine.Bg, b.StatusLine.Separator, b.StatusLine.Content, pad, color.ResetFormat)

    fmt.Fprint(os.Stdout, b.StatusLine.Fg, b.StatusLine.Bg, input, b.StatusLine.Separator, position, color.ResetFormat)
    cursor.SetPos(os.Stdout, b.TermLine + b.CursorLine + 1, b.TermCol + b.CursorCol + len(b.CurrentLine.Gutter) + 1)
}

func (b Buffer) DrawFillLine(lineNum int) {
    cursor.SetPos(os.Stdout, b.TermLine + lineNum + 1, b.TermCol + 1)

    fmt.Fprint(os.Stdout, b.fillLine.Fg, b.fillLine.Bg, b.fillLine.Content, color.ResetFormat)
    cursor.SetPos(os.Stdout, b.TermLine + b.CursorLine + 1, b.TermCol + b.CursorCol + 1)
}

func (b Buffer) Draw() {
    for i := 0; i < b.CursorLine; i++ {
        b.Lines[i].Gutter = fmt.Sprintf(" %v  ", i + 1)
        b.DrawLine(i)
    }

    b.CurrentLine.Gutter = fmt.Sprintf(" %v  ", b.CursorLine + 1)
    b.DrawCursor()

    for i := b.CursorLine + 1; i < len(b.Lines); i++ {
        b.Lines[i].Gutter = fmt.Sprintf(" %v  ", i + 1)
        b.DrawLine(i)
    }

    for i := len(b.Lines); i < b.Height - 1; i ++ {
        b.DrawFillLine(i)
    }
    cursor.SetPos(os.Stdout, b.TermLine + b.CursorLine + 1, b.TermCol + b.CursorCol + len(b.CurrentLine.Gutter) + 1)
}
