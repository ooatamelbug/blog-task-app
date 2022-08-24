package services

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

func ReturnResponse(status bool, message string, data interface{}, token string, err string) Response {
	var errorMessage interface{}
	if err != "" {
		errorMessage = strings.Split(err, "\n")
	}
	res := Response{
		Status:  status,
		Message: message,
		Errors:  errorMessage,
		Data:    data,
		Token:   token,
	}
	return res
}
