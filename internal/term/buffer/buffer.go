package buffer

import (
    "fmt"
    "io"
)

const esc string = "\033"

const (
    enableAlternate string = "[?1049h" 
    disableAlternate string = "[?1049l"
)

func EnableAlternate(io io.Writer) {
    fmt.Fprintf(io, esc + enableAlternate)
}

func DisableAlternate(io io.Writer) {
    fmt.Fprintf(io, esc + disableAlternate)
}

