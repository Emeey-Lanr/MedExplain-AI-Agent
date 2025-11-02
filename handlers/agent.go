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
   utils.Response(c, 200, utils.ErrorResponse{Jsonrpc: "2.0", Id:"", Error:utils.ErrorData{Code:-32700, Message: "Parse Error", Data: err.Error()}})
	return
 }


  

  if reqJsonRPC.Jsonrpc != "2.0" || reqJsonRPC.Jsonrpc == ""{
	utils.Response(c, 200, utils.ErrorResponse{Jsonrpc: "2.0", Id:"", Error:utils.ErrorData{Code:-32600, Message: "Invalid Request", Data: "Unsupported JSON-RPC Version"}})
	return
 
  }

  if reqJsonRPC.Method != "method/send" {
  utils.Response(c, 200, utils.ErrorResponse{Jsonrpc: "2.0", Id:"", Error:utils.ErrorData{Code:-32601, Message: "Method not found", Data: "The method int the request doesn't exist or is not available"}})
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
	fmt.Println(err)
	return
 }

// add gemini reponse to memory db 
 config.History.AddHistory(reqJsonRPC.Params.ContextId, "agent", geminiResponse.Candidates[0].Contents.Parts[0].Text )

 messageParts :=  []map[string]interface{}{{"kind":"text", "text":geminiResponse.Candidates[0].Contents.Parts[0].Text}}

 response := map[string]interface{}{
	"jsonrpc":reqJsonRPC.Jsonrpc,
	"id":reqJsonRPC.Id,
	"result":map[string]interface{}{
		"id":reqJsonRPC.Params.Message.TaskId,
		"contextId":reqJsonRPC.Params.ContextId,
		"status":map[string]interface{}{
			"state":"input-required",
			"timestamp":time.Now().UTC().Format(time.RFC3339),
			"message":map[string]interface{}{
                "messageId":helpers.GenerateContextId("msg-"),
				"role":"agent",
				"parts": messageParts,
				"kind":"message",
				"taskId":reqJsonRPC.Params.Message.TaskId,
			},
		},
	},
 }
 

 utils.Response(c, http.StatusOK, response)
 

}
