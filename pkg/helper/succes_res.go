package helper

import "koriebruh/cqrs/dto"

func SuccessRes(code int, status string, data interface{}) dto.WebRes {
	return dto.WebRes{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
