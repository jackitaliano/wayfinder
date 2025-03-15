package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jackitaliano/wayfinder/internal/term/app"
	"github.com/jackitaliano/wayfinder/internal/tui"
	"github.com/jackitaliano/wayfinder/internal/tui/events"
	"github.com/jackitaliano/wayfinder/internal/tui/input"
)

func main(){
    app.Startup()
    defer app.Cleanup()
    keyChan := make(chan byte)
    defer close(keyChan)

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigChan
        app.Cleanup()
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
    ctx := tui.NewContext(screen.Buffer)
    eventHandler := events.NewEventHandler(screen.Buffer)

    input := input.NewInputHandler(ctx, eventHandler)
    tui.ListenForKeys(keyChan)

    screen.Draw()


    for key := range keyChan {
        input.HandleKey(key)
        eventHandler.HandlePendingEvents()
    }
}
