package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackitaliano/wayfinder/internal/term"
	"github.com/jackitaliano/wayfinder/internal/tui"
)

func main(){
    borderChars := tui.BorderChars{
        '-',
        'x',
        '|',
        'x',
        '-',
        'x',
        '|',
        'x',
    }

    screen := tui.NewScreen(borderChars)
    term.Startup()
    defer term.Cleanup()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigChan
        term.Cleanup()
        fmt.Println("Lines: ", len( screen.Buffer.Lines ) )
        os.Exit(0)
    }()

	buf := make([]byte, 1)

    screen.Draw()

    for {
        os.Stdin.Read(buf) // Read single byte

        if buf[0] == 106 { // j
            screen.MoveCursorDown()
        }
        if buf[0] == 107 { // k
            screen.MoveCursorUp()
        }

        if buf[0] == 104 { // h
            screen.MoveCursorLeft()
        }
        if buf[0] == 108 { // l
            screen.MoveCursorRight()
        }

        if buf[0] == 27 { // ESC key
            term.Cleanup()
            os.Exit(0)
        }
    }
}
