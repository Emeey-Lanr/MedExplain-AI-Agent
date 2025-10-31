package handlers

import (
	"ai-agent/llm"
	"ai-agent/models"

	"github.com/gin-gonic/gin"
)

func Inquire(c *gin.Context) {
   var reqJsonRPC models.JSONRPC_REQUEST

 if err := c.ShouldBindJSON(&reqJsonRPC); err != nil{

 }
 

 parts := models.ContentData{
	Parts: append([]models.TextData{}, models.TextData{Text: "What's animal"}),
 }

geminiRequestBody := models.RequestObject{Contents: append([]models.ContentData{}, parts)} 

// Gemini Request 
 llm.GeminiAIRequest(geminiRequestBody)

}