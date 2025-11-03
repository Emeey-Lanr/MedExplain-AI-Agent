package models

// Each part of the message (text, data, etc.)
type PartData struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

// The main message object sent between user and agent
type MessageData struct {
	Kind      string     `json:"kind,omitempty"`
	Role      string     `json:"role"` // fixed lowercase JSON key
	Parts     []PartData `json:"parts"`
	MessageId string     `json:"messageId"`
	TaskId    string     `json:"taskId"`
}

// Authentication config for push notifications
type AuthenticationD struct {
	Schemes []string `json:"schemes"`
}

// Push notification config from Telex
type PushNotificationConfigD struct {
	Url            string         `json:"url"`
	Token          string         `json:"token"`
	Authentication AuthenticationD `json:"authentication"`
}

// Configuration block that includes blocking and webhook info
type Config struct {
	AcceptedOutputModes   []string              `json:"acceptedOutputModes,omitempty"`
	HistoryLength         int                   `json:"historyLength,omitempty"`
	PushNotificationConfig PushNotificationConfigD `json:"pushNotificationConfig"`
	Blocking              bool                  `json:"blocking"`
}

// The params object in the JSON-RPC request
type MessageObj struct {
	ContextId     string  `json:"contextId"`
	Message       MessageData `json:"message"`
	Configuration Config  `json:"configuration"`
}

// The main JSON-RPC request wrapper
type JSONRPC_REQUEST struct {
	Jsonrpc string     `json:"jsonrpc"`
	Id      string     `json:"id"`
	Method  string     `json:"method"`
	Params  MessageObj `json:"params"`
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
