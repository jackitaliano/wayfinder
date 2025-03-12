package buffer

import (
    "os"

	"github.com/jackitaliano/wayfinder/internal/term/cursor"
)

func (b *Buffer) MoveCursorVertical(numLines int) {
    b.DrawLine(b.CursorLine)

    b.CursorLine += numLines
    b.CurrentLine = &b.Lines[b.CursorLine]

    if b.heldCursorCol <= len(b.CurrentLine.Content) - 1 {
        b.CursorCol = b.heldCursorCol

    } else if b.CursorCol > len(b.CurrentLine.Content) - 1 {
        b.heldCursorCol = b.CursorCol
        b.CursorCol = max(len(b.CurrentLine.Content) - 1, 0)
    }

    b.DrawCursor()
}

func (b *Buffer) MoveCursorDown() {
    if b.CursorLine == len(b.Lines) - 1 {
        return
    }

    b.MoveCursorVertical(1)
}

func (b *Buffer) MoveCursorUp() {
    if b.CursorLine == 0 {
        return
    }

    b.MoveCursorVertical(-1)
}

func (b *Buffer) MoveCursorLeft() {
    if b.CursorCol == 0 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorCol -= 1
    b.heldCursorCol = b.CursorCol

    b.DrawCursor()
}

func (b *Buffer) MoveCursorRight() {
    if b.CursorCol >= len(b.CurrentLine.Content) - 1 {
        return
    }

    b.DrawLine(b.CursorLine)

    b.CursorCol += 1
    b.heldCursorCol = b.CursorCol

    b.DrawCursor()
}

func (b *Buffer) CursorNormalMode() {
    b.StatusLine.Mode = NORMAL
    cursor.SetBlock(os.Stdin)
    b.MoveCursorLeft()
    b.DrawCursor()
}

func (b *Buffer) CursorInsertMode() {
    cursor.SetBar(os.Stdin)
    b.StatusLine.Mode = INSERT
    b.DrawCursor()
}

func (b *Buffer) CursorAppendMode() {

    if len(b.CurrentLine.Content) > 0 {
        b.CursorCol += 1
    }
    cursor.SetBar(os.Stdin)
    b.StatusLine.Mode = INSERT
    b.DrawCursor()
}

func (b *Buffer) CursorHome() {
    b.CursorCol = b.TermCol
    b.DrawCursor()
}

func (b *Buffer) CursorEnd() {
    b.CursorCol = len(b.CurrentLine.Content) - 1

    if b.CursorCol < 0 {
        b.CursorCol = 0;
    }
    b.DrawCursor()
}
