package models



type PartData struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}



type MessageData struct {
  Kind string `json:"kind,omitempty"`
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



type StatusData struct {
   State string `json:"state"`
   TimeStamp string `json:"timestamp"`
   Message MessageData `json:"message"`
}

type ResultData struct {
  Id string `json:"id"`
  ContextId string `json:"contextId"`
  Status StatusData `json:"status"`
}

type JSONRPC_SUCCESS_RESPONSE struct {
  Jsonrpc string `json:"jsonrpc"`
  Id string `json:"id"`
 Result ResultData `json:"result"`
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