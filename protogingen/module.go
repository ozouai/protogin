package protogingen

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ozouai/protogin"
	"github.com/ozouai/protogin/protoginctx"
)

func ApplyMiddlewareList(ctx context.Context, middleware protogin.MiddlewareList, next func(context.Context) error) error {
	return createMiddlewareChain(ctx, middleware, 0, next)
}

func createMiddlewareChain(ctx context.Context, middleware protogin.MiddlewareList, i int, next func(context.Context) error) error {
	if i == len(middleware) {
		return next(ctx)
	}
	return middleware[i](ctx, func(ctx context.Context) error {
		return createMiddlewareChain(ctx, middleware, i+1, next)
	})
}

func AddGinContext(ctx context.Context, ginCtx *gin.Context) context.Context {
	return context.WithValue(ctx, protoginctx.GinCtxKey, ginCtx)
}
