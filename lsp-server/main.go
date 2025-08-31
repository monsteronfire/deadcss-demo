package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const ollamaBaseURL = "http://host.docker.internal:11434/api/"
const modelName = "hf.co/bartowski/Llama-3.2-1B-Instruct-GGUF:latest"

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
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
		Model:  modelName,
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

	functionName := "handleHover"

	prompt := fmt.Sprintf("Write a creative haiku inspired by a programming function called '%s'. Do not mention, reference, or use the function name or any part of it in the haiku. Instead, capture the essence or purpose of what such a function might do in code, using poetic and programming-themed language. Only return the haiku, nothing else. If you use the function name or any part of it, your answer is incorrect.", functionName)

	haiku, err := s.callOllama(prompt)
	if err != nil {
		log.Fatalf("Error calling Ollama API: %v", err)
	}

	log.Println(haiku)
}

func main() {
	log.Println("Haiku LSP starting...")

	server := NewLSPServer()
	server.serve()
}
