package protogin

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ozouai/protogin/protoginctx"
)

type Middleware func(ctx context.Context, f func(ctx context.Context) error) error

type MiddlewareList []Middleware

func IsGinRequest(ctx context.Context) bool {
	val := ctx.Value(protoginctx.GinCtxKey)
	if val == nil {
		return false
	}
	_, ok := val.(*gin.Context)
	return ok
}

func GetGinRequest(ctx context.Context) *gin.Context {
	val := ctx.Value(protoginctx.GinCtxKey)
	if val == nil {
		return nil
	}
	if ginCtx, ok := val.(*gin.Context); ok {
		return ginCtx
	}
	return nil
}

// func Add
