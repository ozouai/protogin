package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ozouai/protogin"
	"github.com/ozouai/protogin/tests/github.com/ozouai/protogin/testpb"
)

func main() {

	handler := &Handler{}
	engine := testpb.NewTestServiceGinServer(handler)
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

func (m *Handler) First_Middleware() protogin.MiddlewareList {
	return protogin.MiddlewareList{
		func(ctx context.Context, f func(ctx context.Context) error) error {
			err := f(ctx)
			if err != nil {
				fmt.Println(err)
			}
			return err
		},
	}
}

func (m *Handler) First(ctx context.Context, request *testpb.FirstRequest) (*testpb.FirstResponse, error) {
	return &testpb.FirstResponse{}, nil
}

func (m *Handler) Second_Middleware() protogin.MiddlewareList {
	return protogin.MiddlewareList{
		func(ctx context.Context, f func(ctx context.Context) error) error {
			err := f(ctx)
			if err != nil {
				fmt.Println(err)
			}
			return err
		},
	}
}

func (m *Handler) Second(ctx context.Context, request *testpb.SecondRequest) (*testpb.SecondResponse, error) {
	return &testpb.SecondResponse{
		Response: "Got ID of " + request.GetId(),
	}, nil
}
