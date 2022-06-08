package protogin

import "context"

type Middleware func(ctx context.Context, f func(ctx context.Context) error) error

type MiddlewareList []Middleware

func ApplyMiddlewareList(ctx context.Context, middleware MiddlewareList, next func(context.Context) error) error {
	return createMiddlewareChain(ctx, middleware, 0, next)
}

func createMiddlewareChain(ctx context.Context, middleware MiddlewareList, i int, next func(context.Context) error) error {
	if i == len(middleware) {
		return next(ctx)
	}
	return middleware[i](ctx, func(ctx context.Context) error {
		return createMiddlewareChain(ctx, middleware, i+1, next)
	})
}
