package main

import (
	"fmt"
	"net/http"
	"strings"
	".."
)

func elizaInterface(w http.ResponseWriter, r *http.Request) {
	
	input := strings.ToLower(r.URL.Query().Get("input"))

	output := eliza.ReplyTo(input)

	fmt.Fprintf(w, output)
}

func main() {
	dir := http.Dir("../webpage")
	fileServer := http.FileServer(dir)

	http.Handle("/", fileServer)
	http.HandleFunc("/chat", elizaInterface)

	http.ListenAndServe(":8080", nil)
}
