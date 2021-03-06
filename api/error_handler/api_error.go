package error_handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrRecordNotFound      = errors.New("record not found")
	ErrInternalServerError = errors.New("internal server error")
)

type ErrorResponse struct {
	Status      int      `json:"status"`
	Error       string   `json:"error"`
	ErrorDetail []string `json:"error_detail"`
}

type Options struct {
	Messages *[]string
}

type Option func(*Options)

func ErrorMessage(messages []string) Option {
	return func(opts *Options) {
		opts.Messages = &messages
	}
}

func ApiErrorHandle(err string, errorKind error, options ...Option) *ErrorResponse {
	log.Println("API Exception: %s", err)

	opts := &Options{
		Messages: nil,
	}

	for _, option := range options {
		option(opts)
	}

	switch errorKind {
	case ErrBadRequest:
		fmt.Println(opts.Messages)
		if opts.Messages == nil {
			opts.Messages = &[]string{enum.BadRequest.String()}
		}
		return &ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "bad request",
			ErrorDetail: *opts.Messages,
		}
	case ErrUnauthorized:
		if opts.Messages == nil {
			opts.Messages = &[]string{enum.Unauthorized.String()}
		}
		return &ErrorResponse{
			Status:      http.StatusUnauthorized,
			Error:       "unauthorized",
			ErrorDetail: *opts.Messages,
		}
	case ErrForbidden:
		if opts.Messages == nil {
			opts.Messages = &[]string{enum.Forbidden.String()}
		}
		return &ErrorResponse{
			Status:      http.StatusForbidden,
			Error:       "forbidden",
			ErrorDetail: *opts.Messages,
		}
	case ErrRecordNotFound:
		if opts.Messages == nil {
			opts.Messages = &[]string{enum.NotFound.String()}
		}
		return &ErrorResponse{
			Status:      http.StatusNotFound,
			Error:       "not found",
			ErrorDetail: *opts.Messages,
		}
	default:
		if opts.Messages == nil {
			opts.Messages = &[]string{enum.InternalServerError.String()}
		}
		return &ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "internal server error",
			ErrorDetail: *opts.Messages,
		}
	}
}
