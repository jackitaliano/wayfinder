package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jackitaliano/wayfinder/internal/term"
	"github.com/jackitaliano/wayfinder/internal/tui"
)

func main(){
    term.Startup()
    defer term.Cleanup()
    keyChan := make(chan byte)
    defer close(keyChan)

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigChan
        term.Cleanup()
        close(keyChan)
        os.Exit(0)
    }()

    borderChars := tui.BorderChars{
        N:'-',
        NE:'x',
        E:'|',
        SE:'x',
        S:'-',
        SW:'x',
        W:'|',
        NW:'x',
    }

    screen := tui.NewScreen(borderChars)
    input := tui.NewInput(&screen)
    tui.ListenForKeys(keyChan)

    screen.Draw()


    for key := range keyChan {
        input.HandleKey(key)
    }
}
