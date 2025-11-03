package handlers

import (
	"ai-agent/config"
	"ai-agent/helpers"
	"ai-agent/llm"
	"ai-agent/models"
	"ai-agent/utils"
	"fmt"
	"net/http"
	"time"
    "github.com/gin-gonic/gin"
	
)






func Inquire(c *gin.Context) {
   
   var reqJsonRPC models.JSONRPC_REQUEST
  
   

 if err := c.ShouldBindJSON(&reqJsonRPC); err != nil{
   utils.Response(c, http.StatusBadRequest, utils.ErrorResponse{Jsonrpc: "2.0", Id:reqJsonRPC.Id, Error:utils.ErrorData{Code:-32700, Message: "Parse Error", Data: err.Error()}})
	return
 }


  

  if reqJsonRPC.Jsonrpc != "2.0" || reqJsonRPC.Jsonrpc == ""{
	utils.Response(c, http.StatusBadRequest, utils.ErrorResponse{Jsonrpc: "2.0", Id:reqJsonRPC.Id, Error:utils.ErrorData{Code:-32600, Message: "Invalid Request", Data: "Unsupported JSON-RPC Version"}})
	return
 
  }



// Create Context Id if it's a begining of a new chat
if reqJsonRPC.Params.ContextId == "" {
   reqJsonRPC.Params.ContextId = helpers.GenerateContextId("ctx-")
}


// Add to the history body which will also serve as as a content for gemini
  config.History.AddHistory(reqJsonRPC.Params.ContextId, "user", reqJsonRPC.Params.Message.Parts[0].Text)

//  Gemini Response
 geminiResponse, err :=  llm.GeminiAIRequest(config.History.GetHistory(reqJsonRPC.Params.ContextId))
 if err != nil{
   utils.Response(c, http.StatusInternalServerError,utils.ErrorResponse{Jsonrpc: "2.0", Id:reqJsonRPC.Id, Error:utils.ErrorData{Code:-32603, Message: "Invalid Server Error", Data: err.Error()}})
	fmt.Println(err)
	return
 }

// add gemini reponse to memory db 
 config.History.AddHistory(reqJsonRPC.Params.ContextId, "agent", geminiResponse.Candidates[0].Contents.Parts[0].Text )

 messageParts :=  []map[string]interface{}{{"kind":"text", "text":geminiResponse.Candidates[0].Contents.Parts[0].Text}}
taskId := helpers.GenerateContextId("task-")
 response := map[string]interface{}{
	"jsonrpc":reqJsonRPC.Jsonrpc,
	"id":reqJsonRPC.Id,
	"result":map[string]interface{}{
		"id":taskId,
		"contextId":reqJsonRPC.Params.ContextId,
		"status":map[string]interface{}{
			"state":"input-required",
			"timestamp":time.Now().UTC().Format(time.RFC3339), 
			"message":map[string]interface{}{
                "messageId":helpers.GenerateContextId("msg-"),
				"role":"agent",
				"parts": messageParts,
				"kind":"message",

			},
		},
		"artifacts": []map[string]interface{}{}, 
    "history":   []map[string]interface{}{},
		"kind":"task",
	},
 }
 

 utils.Response(c, http.StatusOK, response)
 

}
