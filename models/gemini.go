package models


// Request 
type TextData struct {
    Text  string `json:"text"`
}
type ContentData struct {
    Role  string `json:"role,omitempty"`
    Parts []TextData `json:"parts"`
}

type GeminiRequestObject struct {
	Contents []ContentData `json:"contents"`
}



// Response
type ContentObj  struct {
	Contents ContentData `json:"content"`
}


type GeminiReponseObject struct {
    Candidates [] ContentObj `json:"candidates"`
}



// System Instruction And Rquest

// type PartsForSystem struct {
//     Parts []TextData `json:"parts"`
// }

type LLMSystemInstruction struct {
     SystemInstructions ContentData `json:"system_instruction"`
     Contents []ContentData `json:"contents"`
}
