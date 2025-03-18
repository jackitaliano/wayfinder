package log

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/jackitaliano/wayfinder/internal/tui/events"
)

type Close func()

type BufferedLog struct {
    io io.Writer
    buffer *bytes.Buffer
}

func dumpBufferedLogs(bufferedLogs []BufferedLog) {
    for _, bufferedLog := range bufferedLogs {
        fmt.Fprint(bufferedLog.io, bufferedLog.buffer.String())
    }
}


func NewHandler(eventHandler *events.EventHandler) (*slog.Handler, Close) {
    now := time.Now()
	logFileName := fmt.Sprintf("logs%04d%02d%02d_%02d.json", now.Year(), now.Month(), now.Day(), now.Hour())
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
        fmt.Printf("Error opening log file to write: %v\n", err)
        os.Exit(1)
	}

    logLevel := slog.LevelDebug
    statusLevel := slog.LevelDebug

    fileLogOptions := slog.HandlerOptions{Level: logLevel}
    stdLogOptions := slog.HandlerOptions{Level: logLevel}
    tuiLogOptions := slog.HandlerOptions{Level: statusLevel}
    multiLogOptions := slog.HandlerOptions{Level: statusLevel}

    var stdBuffer bytes.Buffer

	stdoutHandler := slog.NewTextHandler(&stdBuffer, &stdLogOptions)
	fileHandler := slog.NewJSONHandler(logFile, &fileLogOptions)
    tuiHandler := NewTuiHandler(eventHandler, &tuiLogOptions)

    multiHandler := NewMultiHandler(&multiLogOptions, stdoutHandler, fileHandler, tuiHandler)

    bufferedLogs := []BufferedLog{
        { os.Stdout, &stdBuffer },
    }

    close := func() {
        dumpBufferedLogs(bufferedLogs)
    }

    return &multiHandler, close
}

