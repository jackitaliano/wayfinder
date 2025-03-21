package buffer

import (
    "fmt"
)

func (b *Buffer) StatusPrintf(format string, a...any) {
    b.StatusLine.Content = fmt.Sprintf(format, a...)
    b.DrawStatusLine()
}

func (b *Buffer) StatusPrint(a...any) {
    b.StatusLine.Content = fmt.Sprint(a...)
    b.DrawStatusLine()
}

func (b *Buffer) insertLine(position int) {
    b.Lines = append(b.Lines[:position], append([]Line{BlankLine()}, b.Lines[position:]...)...)
}

func (b *Buffer) openLine(offset int) {
    b.insertLine(b.CursorLine + offset)

    b.CursorLine += offset
    b.CursorCol = b.TermCol

    b.CurrentLine = &b.Lines[b.CursorLine]

    b.Draw()
}

func (b *Buffer) OpenLineBelow() {
    b.openLine(1)
}

func (b *Buffer) OpenLineAbove() {
    b.openLine(0)
}


func (b *Buffer) DeleteToEnd() {
    line := b.CurrentLine
    if len(line.Content) == 0 {
        return
    }

    b.CurrentLine.Content = line.Content[:b.CursorCol]

    b.CursorCol -= 1

    if b.CursorCol < 0 {
        b.CursorCol = 0
    }

    b.DrawCursor()
}

func (b *Buffer) DeleteChar() {
    line := b.CurrentLine
    if len(line.Content) == 0 {
        return
    }

    b.CurrentLine.Content = line.Content[:b.CursorCol] + line.Content[b.CursorCol + 1:]

    b.CursorCol -= 1
    if b.CursorCol < 0 {
        b.CursorCol = 0
    }

    b.DrawCursor()
}

func (b *Buffer) Backspace() {
    if b.CursorCol > 0 {
        line := b.CurrentLine
        b.CurrentLine.Content = line.Content[:b.CursorCol - 1] + line.Content[b.CursorCol:]
        b.CursorCol -= 1
        b.DrawCursor()

        return
    }


    if b.CursorLine == 0 {
        return
    }

    b.CursorCol = len(b.Lines[b.CursorLine - 1].Content)

    if len(b.CurrentLine.Content) > 0 {
        b.Lines[b.CursorLine - 1].Content += b.CurrentLine.Content
    }

    b.Lines = append(b.Lines[:b.CursorLine], b.Lines[b.CursorLine + 1:]...)

    b.CursorLine -= 1
    b.CurrentLine = &b.Lines[b.CursorLine]

    b.Draw()
}

func (b *Buffer) CarryLine() {
    remains := b.CurrentLine.Content[:b.CursorCol]
    carry := b.CurrentLine.Content[b.CursorCol:]


    newLine := BlankLine()
    newLine.Content = carry

    b.CurrentLine.Content = remains

    b.Lines = append(b.Lines[:b.CursorLine + 1], append([]Line{newLine}, b.Lines[b.CursorLine + 1:]...)...)

    b.CursorLine += 1
    b.CursorCol = b.TermCol

    b.CurrentLine = &b.Lines[b.CursorLine]

    b.Draw()
}

func (b *Buffer) InsertChar(char string) {
    b.CurrentLine.Content = b.CurrentLine.Content[:b.CursorCol] + char + b.CurrentLine.Content[b.CursorCol:]
    b.CursorCol += 1
    b.DrawCursor()
}

