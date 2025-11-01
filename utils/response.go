package utils

import ("github.com/gin-gonic/gin")



type ErrorData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data string `json:"data"`
}

type ErrorResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id string `json:"id"`
	Error ErrorData `json:"error"`
}

func Response(c *gin.Context, code int, data interface{}){
	 c.JSON(code, data)
}



