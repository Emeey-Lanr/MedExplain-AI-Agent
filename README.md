## Medical AI Agent
This an AI agent that explains medical terms and health related question

## Getting Started
```bash
git clone https://github.com/Emeey-Lanr/MedExplain-AI-Agent.git
```

```bash
cd  MedExplain-AI-Agent
```

##To Start Testing Locally
Run
```bash
go run main.go 
```

### Endpoints to Interact with The Agent
#### POST /a2a/medic

##### Request Body
```json
{
"jsonrpc":"2.0",
"id":"test-001",
"method":"message/send",
"params":{
    "message":{
     "kind":"message",
     "role":"user",
    "parts":[{"kind":"text","text":"Explain arthritis."}],
    "taskId":"task-001",
    "messageId":"msg-001"},

    "configuration":{
        "blocking":true
    }
}
```


##### Response Body
```json
{
  "id": "test-001",
  "jsonrpc": "2.0",
  "result": {
     "contextId": "ctx-uuid",
    "history": [],
    "id": "task-uuid",
    "kind": "task",
    "status": {
         "state": "completed",
      "timestamp": "2025-11-03T10:32:14Z",
      "message": {
        "kind": "message",
        "messageId": "msg-uuid",
        "parts": [
          {
            "kind": "text",
            "text": "Arthritis is a condition that causes inflammation in one or more joints, leading to pain, stiffness, and sometimes swelling. It makes it difficult for people to move the affected body parts easily."
          }
        ],
        "role": "agent",
        "taskId": "task-001"
      },
    "artifacts": [
      {
        "artifactId": "artifact-uuid",
        "name": "Gemini-AI-Response",
        "parts": [
          {
            "kind": "text",
            "text": "Arthritis is a condition that causes inflammation in one or more joints, leading to pain, stiffness, and sometimes swelling. It makes it difficult for people to move the affected body parts easily."
          }
        ]
      }
    ],
   
    }
  }
}


```




