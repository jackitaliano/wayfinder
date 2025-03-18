package cursor

import (
    "fmt"
    "io"
)

const esc = "\033"

// cursor position: \033[%y;%xH
const (
    hide string = "[?25l"
    reveal string = "[?25h"
    block string = "[2 q"
    bar string = "[6 q"
    home string = "[H" // moves cursor to home position (0, 0)
    move string = "[%v;%vH" // moves cursor to line #, column #
    up string = "[%vA" // moves cursor up # lines
    down string = "[%vB" // moves cursor down # lines
    right string = "[%vC" // moves cursor right # columns
    left string = "[%vD" // moves cursor left # columns
    // BeginningNext string = "[#E" // moves cursor to beginning of next line, # lines down
    // BeginningPrev string = "[#F" // moves cursor to beginning of previous line, # lines up
    // Column string = "[#G" // moves cursor to column #
    position string = "[6n" // request cursor position (reports as ESC[#;#R)
    // UpScroll string = "M" // moves cursor one line up, scrolling if needed
    save string = " 7"// save cursor position (DEC)
    restoreSave string = " 8" // restores the cursor to the last saved position (DEC)
    // SaveSCO string = "[s" // save cursor position (SCO)
    // RestoreSaveSCO string = "[u" // restores the cursor to the last saved position (SCO)
)

func SetPos(io io.Writer, line int, column int) {
    fmt.Fprintf(io, esc + move, line, column)
}

func MoveUp(io io.Writer, n int) {
    fmt.Fprintf(io, esc + up, n)
}

func MoveDown(io io.Writer, n int) {
    fmt.Fprintf(io, esc + down, n)
}

func MoveLeft(io io.Writer, n int) {
    fmt.Fprintf(io, esc + left, n)
}

func MoveRight(io io.Writer, n int) {
    fmt.Fprintf(io, esc + right, n)
}

func Hide(io io.Writer) {
    fmt.Fprintf(io, esc + hide)
}

func Reveal(io io.Writer) {
    fmt.Fprintf(io, esc + reveal)
}

func SetBlock(io io.Writer) {
    fmt.Fprintf(io, esc + block)
}

func SetBar(io io.Writer) {
    fmt.Fprintf(io, esc + bar)
}

func SaveCursorPos(io io.Writer) {
    fmt.Fprintf(io, esc + save)
}

func RestoreCursorPos(io io.Writer) {
    fmt.Fprintf(io, esc + restoreSave)
}
