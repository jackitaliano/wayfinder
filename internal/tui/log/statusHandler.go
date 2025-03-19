package log

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jackitaliano/wayfinder/internal/tui/events"
)

type StatusHandler struct {
    eventHandler *events.EventHandler
	mu       sync.Mutex
    opts *slog.HandlerOptions
}

func NewStatusHandler(eventHandler *events.EventHandler, opts *slog.HandlerOptions) slog.Handler {
    return &StatusHandler{eventHandler: eventHandler, opts: opts }
}

func (h *StatusHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *StatusHandler) Handle(ctx context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if r.Level < h.opts.Level.Level() {
		return nil
	}

    logEvent := events.LogEvent{Level: r.Level.String(), Msg: r.Message}

    h.eventHandler.PostEvent(logEvent)

	return nil
}

func (h *StatusHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *StatusHandler) WithGroup(name string) slog.Handler {
	return h
}
