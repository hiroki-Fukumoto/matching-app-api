package error_handler

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrRecordNotFound = errors.New("record not found")
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

func ApiErrorHandle(err error, options ...Option) *ErrorResponse {
	fmt.Println(err)

	opts := &Options{
		Messages: nil,
	}

	for _, option := range options {
		option(opts)
	}

	switch err {
	case ErrBadRequest:
		fmt.Println(opts.Messages)
		if opts.Messages == nil {
			opts.Messages = &[]string{"不正なリクエストです"}
		}
		return &ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "bad request",
			ErrorDetail: *opts.Messages,
		}
	case ErrUnauthorized:
		if opts.Messages == nil {
			opts.Messages = &[]string{"認証が必要です"}
		}
		return &ErrorResponse{
			Status:      http.StatusUnauthorized,
			Error:       "unauthorized",
			ErrorDetail: *opts.Messages,
		}
	case ErrForbidden:
		if opts.Messages == nil {
			opts.Messages = &[]string{"実行権限がありません"}
		}
		return &ErrorResponse{
			Status:      http.StatusForbidden,
			Error:       "forbidden",
			ErrorDetail: *opts.Messages,
		}
	case ErrRecordNotFound:
		if opts.Messages == nil {
			opts.Messages = &[]string{"指定されたレコードは存在しません"}
		}
		return &ErrorResponse{
			Status:      http.StatusNotFound,
			Error:       "not found",
			ErrorDetail: *opts.Messages,
		}
	default:
		if opts.Messages == nil {
			opts.Messages = &[]string{"サーバーエラーが発生しました"}
		}
		return &ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "internal server error",
			ErrorDetail: *opts.Messages,
		}
	}
}
