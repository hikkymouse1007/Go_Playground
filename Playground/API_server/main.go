package main

import (
	"net/http"

	"github.com/microsoft/vscode-remote-try-go/Playground/API_server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
