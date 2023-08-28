package protogin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ProtoginError struct {
	e        error
	code     int
	response ProtoginResponse
}

func (m ProtoginError) HTTP(ginCtx *gin.Context) {
	if v, ok := m.response.(*jsonResponse); ok {
		ginCtx.AbortWithStatusJSON(m.code, v.d)
	} else {
		ginCtx.AbortWithStatus(m.code)
	}

}

type ProtoginResponse interface {
	isProtoginResponse()
}

type jsonResponse struct {
	d map[string]interface{}
}

func (m jsonResponse) isProtoginResponse() {

}

func ResponseJSON(j map[string]interface{}) ProtoginResponse {
	return &jsonResponse{d: j}
}

func (e ProtoginError) Error() string {
	if e.e != nil {
		return e.e.Error()
	}
	return fmt.Sprintf("http error: %d", e.code)
}

func Error400BadRequest(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     400,
		response: response,
	}
}

func Error401Unauthorized(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     401,
		response: response,
	}
}

func Error402PaymentRequired(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     402,
		response: response,
	}
}

func Error403Forbidden(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     403,
		response: response,
	}
}

func Error404NotFound(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     404,
		response: response,
	}
}

func Error405MethodNotAllowed(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     405,
		response: response,
	}
}

func Error406NotAcceptable(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     406,
		response: response,
	}
}

func Error407ProxyAuthenticationRequired(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     407,
		response: response,
	}
}

func Error408RequestTimeout(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     408,
		response: response,
	}
}

func Error409Conflict(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     409,
		response: response,
	}
}

func Error410Gone(err error, response ProtoginResponse) error {
	return &ProtoginError{
		e:        err,
		code:     410,
		response: response,
	}
}
