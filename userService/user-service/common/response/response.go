package response

import "github.com/gin-gonic/gin"


type Response struct {
	Status string `json:"status"`
	Message any `json:"message"`
	Data interface{} `json:"data"`
	Token *string `json:"token,omitempty"`
}

type ParamHttpRes struct {
	Code int
	Err error
	Message *string
	Gin *gin.Context
	Data interface{}
	Token *string
}

