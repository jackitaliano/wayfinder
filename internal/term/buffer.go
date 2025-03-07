package term

import (
    "fmt"
    "io"
)

const (
    AltBuf TermSpecifier = "\033[?1049h" 
    MainBuf TermSpecifier = "\033[?1049l"
)

func OpenAltBuf(io io.Writer) {
    fmt.Fprint(io, AltBuf)
}

func OpenMainBuf(io io.Writer) {
    fmt.Fprint(io, MainBuf)
}

