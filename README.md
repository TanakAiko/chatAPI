# chatAPI
A go API for the handling of chat within a forum

## Before starting
import package: 
- [sqlite3](github.com/mattn/go-sqlite3) : `go get github.com/mattn/go-sqlite3`

## Note
The stucture of the body of the request is the following :
- ### created chat
```
{
    action: "createChat"
    body: {
	    senderID      int   
        receiverID    int
	    content       string
    }
}
```

- ### get all chats
```
{
    action: "getChats"
}
```

- ### update statusReceived
```
{
    action: "updateStatusReceived"
    body: {
	    id      int 
    }
}
```

- ### update statusRead
```
{
    action: "updateStatusRead"
    body: {
	    id      int 
    }
}
```

## Testing
### createChat
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action":"createChat", 
  "body": 
    { 
      "senderID": 1,
      "receiverID": 2
      "content": "Hello Casca, it's Guts"
    }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusCreated (201)
- body    : "New chat created"
```

### getChats
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action":"getChats"
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : all the chat data
```

### updateStatusReceived
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action": "updateStatusReceived",
  "boby": {
    "id": 1
  }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : "statusReceived is updated"
```

### updateStatusRead
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action": "updateStatusRead",
  "boby": {
    "id": 1
  }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : "statusRead is updated"
```