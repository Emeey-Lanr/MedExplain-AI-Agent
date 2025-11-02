package llm

import (
	"ai-agent/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


func GeminiAIRequest(contentData []models.ContentData)( models.GeminiReponseObject, error){
	
   for _, value := range contentData{
      if value.Role == "agent" {
         value.Role = "model"
      }

   }

   content := models.GeminiRequestObject{Contents: contentData}
 
   instruction := "You're are MedExplain, an AI that only explains medical terms clearly and kindly. You can always show tips and what to do and what not to do in relation to the medical term if any. If a user asks about unrelated topics, politely refuse and remind them this chat is for medical explantion purpose"

   systemInstructionAndContentData := models.LLMSystemInstruction{
      SystemInstructions: models.ContentData{Parts: append([]models.TextData{}, models.TextData{Text: instruction})},
      Contents: content.Contents,
   }
      
  
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=%s", os.Getenv("GOOGLE_AI_API_KEY"))

   reqBody, _:= json.Marshal(systemInstructionAndContentData) 
	
       fmt.Println(string(reqBody))

  req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
  if err != nil{
   return models.GeminiReponseObject{}, fmt.Errorf("%s", err.Error())
	
  }

  req.Header.Set("Content-Type", "application/json")


   client := &http.Client{}
   resp, err := client.Do(req) 
   if err != nil{
   return models.GeminiReponseObject{}, fmt.Errorf("%s", err.Error())
   }

   defer resp.Body.Close()

   
  var responseData models.GeminiReponseObject
   if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil{
        return models.GeminiReponseObject{}, fmt.Errorf("%s", err.Error())
   }

  return responseData, nil
  
}