package handlers

import (
	"ai-agent/helpers"
	"ai-agent/llm"
	"ai-agent/models"
	"ai-agent/utils"
	"fmt"

	"encoding/json"

	"github.com/gin-gonic/gin"
)






func Inquire(c *gin.Context) {
	history := map[string][]models.ContentData{}

   var reqJsonRPC models.JSONRPC_REQUEST
 

 if err := c.ShouldBindJSON(&reqJsonRPC); err != nil{
   utils.Response(c, 200, utils.ErrorResponse{Jsonrpc: "2.0", Id:"", Error:utils.ErrorData{Code:-32700, Message: "Parse Error", Data: err.Error()}})
	return
 }


  req, _ := json.Marshal(reqJsonRPC)
  
  fmt.Println(string(req))

  if reqJsonRPC.Jsonrpc != "2.0" || reqJsonRPC.Jsonrpc == ""{
	utils.Response(c, 200, utils.ErrorResponse{Jsonrpc: "2.0", Id:"", Error:utils.ErrorData{Code:-32600, Message: "Invalid Request", Data: "Unsupported JSON-RPC Version"}})
	return
 
  }

  if reqJsonRPC.Method != "method/send" {
  utils.Response(c, 200, utils.ErrorResponse{Jsonrpc: "2.0", Id:"", Error:utils.ErrorData{Code:-32601, Message: "Method not found", Data: "The method int the request doesn't exist or is not available"}})
	return
  }



// Create Context Id if it's begining of a new chat
if reqJsonRPC.Params.ContextId == "" {
   reqJsonRPC.Params.ContextId = helpers.GenerateContextId()
}

// Add to the history body which will also serve as as a content for gemini
 history[reqJsonRPC.Params.ContextId]  = append(history[reqJsonRPC.Params.ContextId], helpers.AddToHistory("user", reqJsonRPC.Params.Message.Parts[0].Text)) 
 
 geminiResponse :=  llm.GeminiAIRequest(history[reqJsonRPC.Params.ContextId])

// add gemini reponse to memory db 
 history[reqJsonRPC.Params.ContextId] = append(history[reqJsonRPC.Params.ContextId], helpers.AddToHistory("agent", geminiResponse.Candidates[0].Contents.Parts[0].Text)) 
 

 fmt.Println(geminiResponse)
 

 return

}