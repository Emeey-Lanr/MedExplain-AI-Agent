package handlers

import (
	"ai-agent/models"

	"github.com/gin-gonic/gin"
)

func Inquire(c *gin.Context) {
   var reqJsonRPC models.JSONRPC_REQUEST
   
 if err := c.ShouldBindJSON(&reqJsonRPC); err != nil{

 }


}