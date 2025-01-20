package utils

import (
	"reflect"
)

type Response struct {
	Code    int         `json:"code"`
	Meta    string      `json:"meta"`
	Errors  interface{} `json:"errors"`
	Payload interface{} `json:"payload"`
}

func BuildResponse(code int, message string, errors interface{}, payload interface{}) Response {
	if payload == nil || reflect.ValueOf(payload).IsNil() {
		payload = make([]interface{}, 0)
	} else {
		v := reflect.ValueOf(payload)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Slice {
			payload = []interface{}{payload}
		}
	}

	if errors == nil {
		errors = make([]interface{}, 0)
	} else if reflect.TypeOf(errors).Kind() != reflect.Slice {
		errors = []interface{}{errors}
	}

	return Response{
		Code:    code,
		Meta:    message,
		Errors:  errors,
		Payload: payload,
	}
}
