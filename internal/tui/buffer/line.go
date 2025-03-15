package buffer

import (
	"fmt"
	"io"
	"strings"

	"github.com/jackitaliano/wayfinder/internal/term/color"
)

type Line struct {
    Fg string
    Bg string
    Content string
    Gutter string
}

func BlankLine() Line {
    return Line{"", "", "", ""} 
}

func FillLine(width int) Line {
    return Line{color.BlackBg, "", strings.Repeat(" ", width), ""}
}

type StatusLine struct {
    Fg string
    Bg string
    Separator string
    Mode Mode
    Row int
    Col int
    Content string
    LastInput byte
    LastInputMap string
}

func (c Line) Draw(io io.Writer) {
    fmt.Fprint(io, c.Fg, c.Bg, c.Content)
}
