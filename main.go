package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)

	port := "8080"
	host := os.Getenv("CODESPACE_NAME")
	if host != "" {
		// GitHub Codespaces forwarding URL pattern
		fmt.Printf("🌐 Server running at: https://%s-%s.app.github.dev/\n", host, port)
	} else {
		// Fallback for local development
		fmt.Printf("🌐 Server running at: http://localhost:%s\n", port)
	}

	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Kubernetes 👋")
}
