package llm

import (
	"ai-agent/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)


func GeminiAIRequest(contentData []models.ContentData) models.GeminiReponseObject{
	
   for _, value := range contentData{
      if value.Role == "agent" {
         value.Role = "model"
      }

   }

   content := models.GeminiRequestObject{Contents: contentData}
      
  
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=%s", os.Getenv("GOOGLE_AI_API_KEY"))

   reqBody, _:= json.Marshal(content) 
	
       fmt.Println(string(reqBody))

  req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
  if err != nil{
	log.Fatal(err)
  }

  req.Header.Set("Content-Type", "application/json")


   client := &http.Client{}
   resp, err := client.Do(req) 
   if err != nil{
	log.Fatal(err)
   }

   defer resp.Body.Close()

   
  var responseData models.GeminiReponseObject
   if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil{
      log.Fatal(err.Error())
   }

  return responseData
  
}