package models



type PartData struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

type MessageData struct {
  Role string `json:"Role"`
  Parts []PartData `json:"parts"`
}

type MessageObj struct{
	Message MessageData `json:"message"`
}

type JSONRPC_REQUEST struct {
  Jsonrpc string `json:"jsonrpc"`
  Id string `json:"id"`
  Method string `json:"method"`
  Params MessageObj `json:"params"`
}


type JSONRPC_SUCCESS_RESPONSE struct {
  Jsonrpc string `json:"jsonrpc"`
  Id string `json:"id"`
  Method string `json:"method"`
  Params MessageObj `json:"params"`
}

type JSONRPC_ERROR_RESPONSE struct {
  Jsonrpc string `json:"jsonrpc"`
   
}