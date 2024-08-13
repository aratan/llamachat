package llamachat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Message representa un mensaje dentro de la conversación.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Options define las opciones para la solicitud de chat.
type Options struct {
	Seed        int     `json:"seed"`
	Temperature float64 `json:"temperature"`
}

// ChatRequest contiene todos los parámetros necesarios para una solicitud de chat.
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Options  Options   `json:"options"`
}

// ChatClient es la estructura principal para interactuar con la API de chat.
type ChatClient struct {
	APIURL string
}

// NewChatClient crea una nueva instancia de ChatClient.
func NewChatClient(apiURL string) *ChatClient {
	return &ChatClient{APIURL: apiURL}
}

// SendChatRequest envía una solicitud de chat y devuelve la respuesta concatenada del asistente.
func (c *ChatClient) SendChatRequest(request ChatRequest) (string, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error al serializar el request: %v", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", c.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creando la solicitud HTTP: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error realizando la solicitud: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error leyendo la respuesta: %v", err)
	}

	jsonObjects := strings.Split(string(body), "\n")
	var fullMessage strings.Builder

	for _, jsonObject := range jsonObjects {
		if jsonObject == "" {
			continue
		}

		var chatResponse struct {
			Message Message `json:"message"`
		}
		if err := json.Unmarshal([]byte(jsonObject), &chatResponse); err != nil {
			return "", fmt.Errorf("error al deserializar la respuesta: %v", err)
		}

		fullMessage.WriteString(chatResponse.Message.Content)
	}

	return fullMessage.String(), nil
}
