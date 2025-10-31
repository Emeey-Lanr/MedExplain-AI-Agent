package llm

import (
	"ai-agent/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


func GeminiAIRequest(content  models.RequestObject) {
	
    api_key := os.Getenv("GOOGLE_AI_API_KEY") //google gemini api key from dotenv
  
	url := fmt.Sprintf("https://genrativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s", api_key)

   reqBody, _:= json.Marshal(content) // encoding rquest into json 


  req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
  if err != nil{
	log.Fatal(err)
  }

  req.Header.Set("contet-type", "application/json")

   client := &http.Client{}
   resp, err := client.Do(req) 

   if err != nil{
	log.Fatal(err)
   }

   defer resp.Body.Close()
 
   respBody, _ := io.ReadAll(resp.Body)

   fmt.Println(string(respBody))
  

}