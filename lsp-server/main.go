package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const ollamaBaseURL = "http://host.docker.internal:11434/api/"

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     string `json:"done"`
}

// Custom type representing the LSP server
type LSPServer struct {
	documents map[string]string
}

// Factory function to create new LSPServer instances
func NewLSPServer() *LSPServer {
	return &LSPServer{
		documents: make(map[string]string),
	}
}

func (s *LSPServer) callOllama(prompt string) (string, error) {
	requestBody := OllamaRequest{
		Model:  "Llama-3.2-1B-Instruct-GGUF",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	response, err := http.Post(
		ollamaBaseURL+"generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var ollamaResp OllamaResponse
	err = json.NewDecoder(response.Body).Decode(&ollamaResp)
	if err != nil {
		return "", err
	}

	return ollamaResp.Response, nil
}

func (s *LSPServer) serve() {
	log.Println("LSP Server is now serving...")
}

func main() {
	log.Println("Haiku LSP starting...")

	server := NewLSPServer()
	server.serve()
}
