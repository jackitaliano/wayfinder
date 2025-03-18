package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackitaliano/wayfinder/internal/term/app"
	"github.com/jackitaliano/wayfinder/internal/tui/buffer"
	"github.com/jackitaliano/wayfinder/internal/tui/context"
	"github.com/jackitaliano/wayfinder/internal/tui/events"
	"github.com/jackitaliano/wayfinder/internal/tui/input"
	"github.com/jackitaliano/wayfinder/internal/tui/log"
)

func main(){
    app.Startup()
    keyChan := make(chan byte)

    width, height := app.GetSize()
    buffer := buffer.NewBuffer(0, 0, width, height)

    ctx := context.NewContext()

    eventHandler := events.NewEventHandler(ctx, buffer)
    logHandler, logCloser := log.NewHandler(eventHandler)

    logger := slog.New(*logHandler)

    slog.SetDefault(logger)


    inputHandler := input.NewInputHandler(ctx, eventHandler)
    input.ListenForKeys(keyChan)

    buffer.Draw()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        for sig := range sigChan {
            if sig == syscall.SIGINT {
                keyChan <- 3
            }
        }
    }()

    for key := range keyChan {
        if key == 3 {
            close(sigChan)
            close(keyChan)
            app.Cleanup()
            logCloser()
            return
        }

        // logger.Info(fmt.Sprintf("Key: %v", key))

        inputHandler.HandleKey(key)
        eventHandler.HandlePendingEvents()
    }
}
