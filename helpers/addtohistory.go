package helpers

import "ai-agent/models"

func AddToHistory(role string, text string) models.ContentData {
	

 	addContent := models.ContentData{Role: role, Parts: append([]models.TextData{}, models.TextData{Text: text})}
 
	return addContent
}