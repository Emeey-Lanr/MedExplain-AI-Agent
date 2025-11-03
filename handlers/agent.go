package handlers

import (
	"ai-agent/config"
	"ai-agent/helpers"
	"ai-agent/llm"
	"ai-agent/models"
	"ai-agent/utils"
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
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
 geminiResponse, err :=  llm.GeminiAIRequest([]models.ContentData{{Parts: []models.TextData{{Text: reqJsonRPC.Params.Message.Parts[0].Text}}}})
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
			"state":"completed",
			"timestamp":time.Now().UTC().Format(time.RFC3339), 
			"message":map[string]interface{}{
                "messageId":helpers.GenerateContextId("msg-"),
				"role":"agent",
				"parts": messageParts,
				"kind":"message",
				"taskId": taskId,
			},
		},
		"artifacts": []map[string]interface{}{{
			"artifactId":helpers.GenerateContextId("art-"),
			"name":"Gemini AI Respone",
			"parts":[]map[string]interface{}{
				{"kind":"text", "text":messageParts},
			},
		},
		}, 
    "history":   []map[string]interface{}{},
		"kind":"task",
	},
 }
 

 utils.Response(c, http.StatusOK, response)


	go func() { // using goroutine so it doesn't block
    pushUrl := reqJsonRPC.Params.Configuration.PushNotificationConfig.Url
    token := reqJsonRPC.Params.Configuration.PushNotificationConfig.Token
		
	fmt.Println("PushNotificationConfig:", reqJsonRPC.Params.Configuration.PushNotificationConfig.Token)

    payload, _ := json.Marshal(response)
    req, _ := http.NewRequest("POST", pushUrl, bytes.NewBuffer(payload))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending to Telex:", err)
        return
    }
    defer resp.Body.Close()
    fmt.Println("Sent to Telex:", resp.Status)
}()
 

}
