package term

import (
    "fmt"
    "io"
)

type Color struct {
    red uint8
    green uint8
    blue uint8
}

// fg: \033[38;2;R;G;Bm
// bg: \033[48;2;R;G;Bm
// reset: \033[0m
const (
    WhiteFg TermSpecifier = "\033[38;2;255;255;255m"
    BlackFg TermSpecifier = "\033[38;2;0;0;0m"
    RedFg TermSpecifier = "\033[38;2;255;0;0m"
    GreenFg TermSpecifier = "\033[38;2;0;255;0m"
    BlueFg TermSpecifier = "\033[38;2;0;0;255m"
    GrayFg TermSpecifier = "\033[38;2;40;40;40m"

    WhiteBg TermSpecifier = "\033[48;2;255;255;255m"
    BlackBg TermSpecifier = "\033[48;2;0;0;0m"
    RedBg TermSpecifier = "\033[48;2;255;0;0m"
    GreenBg TermSpecifier = "\033[48;2;0;255;0m"
    BlueBg TermSpecifier = "\033[48;2;0;0;255m"
    GrayBg TermSpecifier = "\033[48;2;40;40;40m"

    Reset TermSpecifier = "\033[0m"
)

func ColorFgRGB(red uint8, green uint8, blue uint8) string {
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
}

func ColorBgRGB(red uint8, green uint8, blue uint8) string {
    return fmt.Sprintf("\033[48;2;%d;%d;%dm", red, green, blue)
}

func SetFg(io io.Writer, r uint8, g uint8, b uint8) {
    color := ColorFgRGB(r, g, b)
    fmt.Fprint(io, color)
}

func SetBg(io io.Writer, r uint8, g uint8, b uint8) {
    color := ColorBgRGB(r, g, b)
    fmt.Fprint(io, color)
}

func ResetColor(io io.Writer) {
    fmt.Fprint(io, Reset)
}

