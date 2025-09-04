package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const ollamaBaseURL = "http://host.docker.internal:11434/api/"
const modelName = "hf.co/bartowski/Llama-3.2-1B-Instruct-GGUF:latest"

type Message struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

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

func (s *LSPServer) readMessage(reader *bufio.Reader) (*Message, error) {
	// Read "Header Part"
	contentLength := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break // End of headers
		}

		if strings.HasPrefix(line, "Content-Length:") {
			fmt.Sscanf(line, "Content-Length: %d", &contentLength)
		}
	}

	// Read "Content Part"
	content := make([]byte, contentLength)
	_, err := io.ReadFull(reader, content)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = json.Unmarshal(content, &msg)
	return &msg, err
}

func (s *LSPServer) writeMessage(write io.Writer, msg *Message) error {
	return nil
}

func (s *LSPServer) handleMessage(msg *Message) *Message {
	switch msg.Method {
	case "initialize":
		return &Message{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Result: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"hoverProvider":      true,
					"codeActionProvider": true,
				},
			},
		}

	case "textDocument/didOpen":
		var params struct {
			TextDocument struct {
				URI  string `json:"uri"`
				Text string `json:"text"`
			} `json:"textDocument"`
		}
		json.Unmarshal(msg.Params.([]byte), &params)
		s.documents[params.TextDocument.URI] = params.TextDocument.Text
		return nil
	}

	return nil
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
		log.Printf("Could not connect to Ollama service at %s: %v", ollamaBaseURL+"generate", err)
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

func (s *LSPServer) serve__() {
	log.Println("LSP Server is now serving...")

	functionName := "handleHover"

	prompt := fmt.Sprintf("Write a creative haiku inspired by a programming function called '%s'. Do not mention, reference, or use the function name or any part of it in the haiku. Instead, capture the essence or purpose of what such a function might do in code, using poetic and programming-themed language. Only return the haiku, nothing else. If you use the function name or any part of it, your answer is incorrect.", functionName)

	haiku, err := s.callOllama(prompt)
	if err != nil {
		log.Fatalf("Error calling Ollama API: %v", err)
	}

	log.Println(haiku)
}

func (s *LSPServer) serve() {
	// receives messages from the Language Client
	reader := bufio.NewReader(os.Stdin)
	// sends messages to the Language Client
	writer := os.Stdout

	for {
		msg, err := s.readMessage(reader)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		response := s.handleMessage(msg)
		if response != nil {
			err = s.writeMessage(writer, response)
			if err != nil {
				log.Printf("Error writing response: %v", err)
				break
			}
		}
	}
}

func main() {
	log.SetOutput(os.Stderr)

	server := NewLSPServer()
	log.Println("Haiku LSP starting...")

	server.serve()
}
