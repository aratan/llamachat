
# llamachat

`llamachat` es un paquete de Go que facilita la interacción con la API de chat Llama.

## Uso

```go
package main

import (
    "fmt"
    "github.com/tu_usuario/llamachat"
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
