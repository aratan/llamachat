
# ollama4go

`ollama4go` es un paquete de Go que facilita la interacción con la API del chat en Ollama.
Cuando hagas tu app tendras que instalar el paquete así:

is a Go package that facilitates interaction with the Ollama chat API. When you make your app you will need to install the package like this:
```go
go get github.com/aratan/ollama4go
```

## Uso

```go
package main

import (
    "fmt"
    "github.com/aratan/ollama4go"
)

func main() {
    client := llamachat.NewChatClient("http://localhost:11434/api/chat")
    
    request := llamachat.ChatRequest{
        Model: "llama3.1",
        Messages: []llamachat.Message{
            {Role: "user", Content: "Hello!"},
        },
        Options: llamachat.Options{
            Seed: 101,
            Temperature: 0,
        },
    }

    response, err := client.SendChatRequest(request)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Respuesta del asistente:", response)
}
