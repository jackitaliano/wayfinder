package log

import (
	"context"
	"log/slog"
	"sync"
)

type MultiHandler struct {
	mu       *sync.Mutex
    opts *slog.HandlerOptions
    handlers []slog.Handler
}

func NewMultiHandler(opts *slog.HandlerOptions, handlers ...slog.Handler) slog.Handler {
    return &MultiHandler{mu: &sync.Mutex{}, opts: opts, handlers: handlers}
}

func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

    for _, h := range h.handlers {
		_ = h.Handle(ctx, r)
	}

	return nil
}

func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
    if len(attrs) == 0 {
        return h
    }

    h2 := *h
    
    handlers := make([]slog.Handler, len(h.handlers))

    copy(handlers, h.handlers)

    for i, handler := range h2.handlers {
        handlers[i] = handler.WithAttrs(attrs)
    }

    h2.handlers = handlers

	return &h2
}

func (h *MultiHandler) WithGroup(name string) slog.Handler {
    if len(name) == 0 {
        return h
    }

    h2 := *h
    
    handlers := make([]slog.Handler, len(h.handlers))

    copy(handlers, h.handlers)

    for i, handler := range h2.handlers {
        handlers[i] = handler.WithGroup(name)
    }

    h2.handlers = handlers

	return &h2
}
