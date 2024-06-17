package common

import "github.com/gin-gonic/gin"

type Resp struct {
	Code    int
	Data    interface{}
	Message string
}

func ErrResp(err error) interface{} {
	return Resp{Code: -1, Message: err.Error()}
}

func OkResp(data interface{}) interface{} {
	return gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
	}
}
