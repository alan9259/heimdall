package platform

import (
	"fmt"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type HttpError struct {
	HttpErrors map[string]interface{} `json:"errors"`
}

func NewHttpError(err error) HttpError {
	e := HttpError{}
	e.HttpErrors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.HttpErrors["body"] = v.Message
	default:
		e.HttpErrors["body"] = v.Error()
	}
	return e
}

func NewValidatorError(err error) HttpError {
	e := HttpError{}
	e.HttpErrors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.HttpErrors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}

func AccessForbidden() HttpError {
	e := HttpError{}
	e.HttpErrors = make(map[string]interface{})
	e.HttpErrors["body"] = "Access forbidden."
	return e
}

func NotFound() HttpError {
	e := HttpError{}
	e.HttpErrors = make(map[string]interface{})
	e.HttpErrors["body"] = "Resource not found."
	return e
}

func AlreadyRegistered() HttpError {
	e := HttpError{}
	e.HttpErrors = make(map[string]interface{})
	e.HttpErrors["code"] = "4001"
	e.HttpErrors["body"] = "The email address is taken."
	return e
}
