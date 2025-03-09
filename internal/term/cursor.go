package term

import (
    "fmt"
    "io"
)

// cursor position: \033[%y;%xH
const (
    Cursor00 TermSpecifier = "\033[H"
    CursorHide TermSpecifier = "\033[?25l"
    CursorReveal TermSpecifier = "\033[?25h"
    CursorBlock TermSpecifier = "\033[2 q"
    CursorLine TermSpecifier = "\033[6 q"
)

func CursorPos(x int, y int) string {
    return fmt.Sprintf("\033[%d;%dH", y, x)
}

func HideCursor(io io.Writer) {
    fmt.Fprint(io, CursorHide)
}

func RevealCursor(io io.Writer) {
    fmt.Fprint(io, CursorReveal)
}

func SetBlockCursor(io io.Writer) {
    fmt.Fprint(io, CursorBlock)
}

func SetLineCursor(io io.Writer) {
    fmt.Fprint(io, CursorLine)
}

func SetCursor(io io.Writer, x int, y int) {
    pos := CursorPos(x, y)
    fmt.Fprint(io, pos)
}

func ResetCursor(io io.Writer) {
    pos := Cursor00
    fmt.Fprint(io, pos)
}

