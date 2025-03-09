package color

import (
    "fmt"
    "io"
)

const esc = "\033"

const reset string = "[0m"
const foreground = "[38;2;"
const background = "[48;2;"
const colorRGB string = "%v;%v;%vm"

const (
    white string = "255;255;255m"
    black string = "0;0;0m"
    red string = "255;0;0m"
    green string = "0;255;0m"
    blue string = "0;0;255m"
    gray string = "40;40;40m"
)

const (
    WhiteFg = esc + foreground + white
    BlackFg = esc + foreground + black
    GreenFg = esc + foreground + green
    BlueFg = esc + foreground + blue
    GrayFg = esc + foreground + gray
)

const (
    WhiteBg = esc + background + white
    BlackBg = esc + background + black
    GreenBg = esc + background + green
    BlueBg = esc + background + blue
    GrayBg = esc + background + gray
)

const (
    ResetFormat = esc + reset
)

func GetFg(io io.Writer, red uint8, green uint8, blue uint8) string {
    return fmt.Sprintf(esc + foreground + colorRGB, red, green, blue)
}

func GetBg(io io.Writer, red uint8, green uint8, blue uint8) string {
    return fmt.Sprintf(esc + background + colorRGB, red, green, blue)
}

func SetFg(io io.Writer, red uint8, green uint8, blue uint8) {
    fmt.Fprintf(io, esc + foreground + colorRGB, red, green, blue)
}

func SetBg(io io.Writer, red uint8, green uint8, blue uint8) {
    fmt.Fprintf(io, esc + background + colorRGB, red, green, blue)
}

func Reset(io io.Writer) {
    fmt.Fprint(io, esc + reset)
}
