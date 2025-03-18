package log

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jackitaliano/wayfinder/internal/tui/events"
)

type TuiHandler struct {
    eventHandler *events.EventHandler
	mu       sync.Mutex
    opts *slog.HandlerOptions
}

func NewTuiHandler(eventHandler *events.EventHandler, opts *slog.HandlerOptions) slog.Handler {
    return &TuiHandler{eventHandler: eventHandler, opts: opts }
}

func (h *TuiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *TuiHandler) Handle(ctx context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if r.Level < h.opts.Level.Level() {
		return nil
	}

    logEvent := events.LogEvent{Level: r.Level.String(), Msg: r.Message}
    h.eventHandler.PostEvent(logEvent)

	return nil
}

func (h *TuiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *TuiHandler) WithGroup(name string) slog.Handler {
	return h
}
