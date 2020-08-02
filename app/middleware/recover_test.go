package middleware_test

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/leoleoasd/EduOJBackend/app/middleware"
	"github.com/leoleoasd/EduOJBackend/app/response"
	"github.com/pkg/errors"
	"testing"
)

func TestRecover(t *testing.T) {
	e := echo.New()

	e.Use(middleware.Recover)
	e.POST("/panics_with_error", func(context echo.Context) error {
		panic(errors.New("123"))
	})
	e.POST("/panics_with_other", func(context echo.Context) error {
		panic("123")
	})

	req := MakeReq(t, "POST", "/panics_with_error", &bytes.Buffer{})
	resp := MakeResp(req, e)
	JsonEQ(t, response.MakeInternalErrorResp(), resp)
	req = MakeReq(t, "POST", "/panics_with_other", &bytes.Buffer{})
	resp = MakeResp(req, e)
	JsonEQ(t, response.MakeInternalErrorResp(), resp)

}
