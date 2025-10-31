package models


// Request 
type TextData struct {
    Text  string `json:"text"`
}
type ContentData struct {
    Parts [] TextData `json:"parts"`
}

type RequestObject struct {
	Contents []ContentData `json:"contents"`
}



// Response
