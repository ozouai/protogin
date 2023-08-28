package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozouai/protogin"
	"github.com/ozouai/protogin/tests/github.com/ozouai/protogin/testpb"
)

func main() {

	handler := &Handler{}
	engine := gin.New()
	testpb.NewTestServiceGinServer(handler, engine)
	go engine.Run(":8000")
	timer := time.NewTimer(time.Millisecond * 200)
	<-timer.C
	res, err := http.Get("http://127.0.0.1:8000/second/blah")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	_ = res
}

type Handler struct {
}

func (m *Handler) First() testpb.FirstHandler {
	return testpb.FirstHandler{
		Middleware: protogin.MiddlewareList{
			func(ctx context.Context, f func(ctx context.Context) error) error {
				err := f(ctx)
				if err != nil {
					fmt.Println(err)
				}
				return err
			},
		},
		Handler: func(ctx context.Context, fr *testpb.FirstRequest) (*testpb.FirstResponse, error) {
			return &testpb.FirstResponse{}, nil
		},
	}
}

func (m *Handler) Second() testpb.SecondHandler {
	return testpb.SecondHandler{
		Middleware: protogin.MiddlewareList{
			func(ctx context.Context, f func(ctx context.Context) error) error {
				return protogin.Error403Forbidden(nil, protogin.ResponseJSON(map[string]interface{}{"error": "Unauthenticated"}))
				err := f(ctx)
				if err != nil {
					fmt.Println(err)
				}
				return err
			},
		},
		Handler: func(ctx context.Context, request *testpb.SecondRequest) (*testpb.SecondResponse, error) {

			return &testpb.SecondResponse{
				Response: "Got ID of " + request.GetId(),
			}, nil
		},
	}
}

func (m *Handler) SecondPost() testpb.SecondPostHandler {
	return testpb.SecondPostHandler{
		Middleware: protogin.MiddlewareList{
			func(ctx context.Context, f func(ctx context.Context) error) error {
				err := f(ctx)
				if err != nil {
					fmt.Println(err)
				}
				return err
			},
		},
		Handler: func(ctx context.Context, request *testpb.SecondRequest) (*testpb.SecondResponse, error) {

			return &testpb.SecondResponse{
				Response: "Got ID of " + request.GetId(),
			}, nil
		},
	}
}
