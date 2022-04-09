package model

import "net/http"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
}
