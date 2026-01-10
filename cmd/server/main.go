package main

import (
	"impostor/internal/platform/server"
)

func main() {
	srv := server.NewServer()
	srv.Run(":8080")
}
