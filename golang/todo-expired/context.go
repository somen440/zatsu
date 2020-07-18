package main

import "context"

type contextKey int

const (
	uidKey contextKey = iota
)

func GetUID(ctx context.Context) string {
	return ctx.Value(uidKey).(string)
}

func SetUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, uidKey, uid)
}
