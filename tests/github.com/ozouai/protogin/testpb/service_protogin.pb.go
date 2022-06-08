// Code generated by protoc-gen-protogin. DO NOT EDIT.

package testpb

import "context"
import "github.com/gin-gonic/gin"
import "github.com/ozouai/protogin/protogingen"
import "github.com/ozouai/protogin"
import "github.com/golang/protobuf/jsonpb"

type TestService_GinHandler interface {
	First(context.Context, *FirstRequest) (*FirstResponse, error)
	First_Middleware() protogin.MiddlewareList
	Second(context.Context, *SecondRequest) (*SecondResponse, error)
	Second_Middleware() protogin.MiddlewareList
	SecondPost(context.Context, *SecondRequest) (*SecondResponse, error)
	SecondPost_Middleware() protogin.MiddlewareList
}

func NewTestServiceGinServer(handler TestService_GinHandler) *gin.Engine {
	engine := gin.New()
	engine.GET("/first", func(ginCtx *gin.Context) {
		var err error
		mainCtx := ginCtx.Request.Context()
		request := &FirstRequest{}
		var responseString string
		err = protogingen.ApplyMiddlewareList(mainCtx, handler.First_Middleware(), func(ctx context.Context) error {
			response, err := handler.First(ctx, request)
			if err != nil {
				return err
			}
			responseString, err = (&jsonpb.Marshaler{}).MarshalToString(response)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			ginCtx.AbortWithError(500, err)
			return
		}
		ginCtx.Status(200)
		ginCtx.Writer.WriteString(responseString)
	})
	engine.GET("/second/:id", func(ginCtx *gin.Context) {
		var err error
		mainCtx := ginCtx.Request.Context()
		request := &SecondRequest{}
		var responseString string
		request.Id = ginCtx.Param("id")
		err = protogingen.ApplyMiddlewareList(mainCtx, handler.Second_Middleware(), func(ctx context.Context) error {
			response, err := handler.Second(ctx, request)
			if err != nil {
				return err
			}
			responseString, err = (&jsonpb.Marshaler{}).MarshalToString(response)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			ginCtx.AbortWithError(500, err)
			return
		}
		ginCtx.Status(200)
		ginCtx.Writer.WriteString(responseString)
	})
	engine.POST("/second/:id", func(ginCtx *gin.Context) {
		var err error
		mainCtx := ginCtx.Request.Context()
		request := &SecondRequest{}
		var responseString string
		err = jsonpb.Unmarshal(ginCtx.Request.Body, request)
		if err != nil {
			ginCtx.AbortWithError(400, err)
			return
		}
		request.Id = ginCtx.Param("id")
		err = protogingen.ApplyMiddlewareList(mainCtx, handler.SecondPost_Middleware(), func(ctx context.Context) error {
			response, err := handler.SecondPost(ctx, request)
			if err != nil {
				return err
			}
			responseString, err = (&jsonpb.Marshaler{}).MarshalToString(response)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			ginCtx.AbortWithError(500, err)
			return
		}
		ginCtx.Status(200)
		ginCtx.Writer.WriteString(responseString)
	})
	return engine
}
