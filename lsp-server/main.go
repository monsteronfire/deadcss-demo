package main

import "log"

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

func main() {
	log.Println("Haiku LSP starting...")
}
