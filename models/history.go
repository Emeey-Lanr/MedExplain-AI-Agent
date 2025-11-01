package models


type Text struct {
	  Text string `json:"text"`
}


type History2 struct{
	Role string `json:"role"`
	Parts []Text `json:"parts"`
}

