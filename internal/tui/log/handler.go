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

type Open func() (io.Writer, error)

type BufferedLog struct {
    Open Open
    buffer *bytes.Buffer
}

func dumpBufferedLogs(bufferedLogs []BufferedLog) {
    for _, bufferedLog := range bufferedLogs {
        if bufferedLog.buffer.Len() == 0 {
            return
        }

        bufferIO, err := bufferedLog.Open()

        if err != nil {
            fmt.Printf("Unable to open log buffered IO: %v\n", err)
            continue
        }

        if closer, ok := bufferIO.(io.Closer); ok {
            defer closer.Close()
        }

        fmt.Fprint(bufferIO, bufferedLog.buffer.String())
    }
}

type MultiOptions struct {
    GlobalOpts *slog.HandlerOptions
    StdOpts *slog.HandlerOptions
    FileOpts *slog.HandlerOptions
    StatusOpts *slog.HandlerOptions
}

func NewHandler(eventHandler *events.EventHandler, opts *MultiOptions) (*slog.Handler, Close) {
    now := time.Now()

    var stdBuffer bytes.Buffer
    var fileBuffer bytes.Buffer

	fileHandler := slog.NewJSONHandler(&fileBuffer, opts.FileOpts)
	stdoutHandler := slog.NewTextHandler(&stdBuffer, opts.StdOpts)
    statusHandler := NewStatusHandler(eventHandler, opts.StatusOpts)

    multiHandler := NewMultiHandler(opts.GlobalOpts, fileHandler, stdoutHandler, statusHandler)

    fileOpen := func() (io.Writer, error) {
        path := "logs"
        logFileName := fmt.Sprintf("%v/%04d%02d%02d_%02d.json", path, now.Year(), now.Month(), now.Day(), now.Hour())

        os.Mkdir(path, os.ModePerm)
        logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

        if err != nil {
            return nil, err
        }

        return logFile, nil
    }

    stdoutOpen := func() (io.Writer, error) {
        return os.Stdout, nil
    }

    bufferedLogs := []BufferedLog{
        { fileOpen, &fileBuffer },
        { stdoutOpen, &stdBuffer },
    }

    close := func() {
        dumpBufferedLogs(bufferedLogs)
    }

    return &multiHandler, close
}

