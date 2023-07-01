package utils

import (
	"context"
	"regexp"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	reFn = regexp.MustCompile(`([^/]+)/?$`)
)

func Trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}

func NewSpan(ctx context.Context, fn string) (context.Context, trace.Span) {
	tr := otel.Tracer("")
	fn = reFn.FindString(fn)
	return tr.Start(ctx, fn)
}
