package helpers

import "context"

type key string

const (
	ctxUserIDKey key = "UserID"
)

func NewWithUserIDContext(ctx context.Context, v int) context.Context {
	return context.WithValue(ctx, ctxUserIDKey, v)
}

func ExtractUserID(ctx context.Context) (int, bool) {
	v, ok := ctx.Value(ctxUserIDKey).(int)
	return int(v), ok
}
