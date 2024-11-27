package helper

import (
	"errors"
	"koriebruh/cqrs/dto"
	"net/http"
)

var (
	ErrBadRequest        = errors.New("BAD REQUEST")
	ErrInternalServerErr = errors.New("INTERNAL SERVER ERROR")
	ErrNotFound          = errors.New("NOT FOUND")
)

func ErrorResponseMsg(err error, msgErr error) dto.WebRes {
	if errors.Is(err, ErrBadRequest) {
		return dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   msgErr.Error(),
		}
	} else if errors.Is(err, ErrNotFound) {
		return dto.WebRes{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   msgErr.Error(),
		}
	} else {
		return dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   msgErr.Error(),
		}
	}
}
