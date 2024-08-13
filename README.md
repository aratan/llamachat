
# ollama4go

`ollama4go` es un paquete de Go que facilita la interacción con la API del chat en Ollama.
Cuando hagas tu app tendras que instalar el paquete así:

is a Go package that facilitates interaction with the Ollama chat API. When you make your app you will need to install the package like this:
```go
go get github.com/aratan/ollama4go
```

### Main.go

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
```

# Otra opcion

### config.yaml
```yaml
api:
  url: "http://localhost:11434/api/chat"

chat:
  model: "llama3.1"
  seed: 101
  temperature: 0

message:
  role: "user"
  content: "en python como seria el codigo: Hello!"
```

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"

    "github.com/aratan/ollama4go"
    "gopkg.in/yaml.v2"
)

type Config struct {
    API struct {
        URL string `yaml:"url"`
    } `yaml:"api"`
    Chat struct {
        Model       string  `yaml:"model"`
        Seed        int     `yaml:"seed"`
        Temperature float32 `yaml:"temperature"`
    } `yaml:"chat"`
    Message struct {
        Role    string `yaml:"role"`
        Content string `yaml:"content"`
    } `yaml:"message"`
}

func loadConfig() (*Config, error) {
    var config Config
    data, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}

func main() {
    config, err := loadConfig()
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    client := llamachat.NewChatClient(config.API.URL)

    request := llamachat.ChatRequest{
        Model: config.Chat.Model,
        Messages: []llamachat.Message{
            {Role: config.Message.Role, Content: config.Message.Content},
        },
        Options: llamachat.Options{
            Seed:        config.Chat.Seed,
            Temperature: config.Chat.Temperature,
        },
    }

    response, err := client.SendChatRequest(request)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("*** ``` ", response, "``` ***")
}
```
