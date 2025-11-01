package models



type PartData struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}



type MessageData struct {
  Role string `json:"Role"`
  Parts []PartData `json:"parts"`
  MessageId string `json:"messageId"`
  TaskId  string `json:"taskId"`

}


type Config struct {
	Blocking bool `json:"blocking"`
}

type MessageObj struct{
   ContextId string `json:"contextId"`
	Message MessageData `json:"message"`
	  Configuration  Config `json:"configuration"`
}

type JSONRPC_REQUEST struct {
  Jsonrpc string `json:"jsonrpc"`
  Id string `json:"id"`
  Method string `json:"method"`
  Params MessageObj `json:"params"`
}



// Success Response

type JSONRPC_SUCCESS_RESPONSE struct {
  Jsonrpc string `json:"jsonrpc"`
  Id string `json:"id"`
  Method string `json:"method"`
  Params MessageObj `json:"params"`
}


// Error Response

type ErrorReponse struct {
    Code int `json:"code"`
	Message string `json:"message"`
	Data string `json:"data,omitempty"`
}

type JSONRPC_ERROR_RESPONSE struct {
  Jsonrpc string `json:"jsonrpc"`
  Id string `json:"id"`
  Error ErrorReponse `json:"error"`
  
}