package util

import "errors"

type ResponseWrapper struct {
    Code  	string      `json:"code"`
    Message string      `json:"message"`
    Result  interface{} `json:"data,omitempty"`
}

var (
    ErrInvalidInput = errors.New("invalid input")
    ErrNotFound     = errors.New("resource not found")
    ErrServer       = errors.New("internal server error")
)
