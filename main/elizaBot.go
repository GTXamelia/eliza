package main

import (
	"fmt"
	"net/http"
	//".."
)

func elizaInterface(w http.ResponseWriter, r *http.Request) {
	
	userInput := r.URL.Query().Get("user-input")



	output := userInput

	fmt.Fprintf(w, output)
	fmt.Println(output)
}

func main() {
	dir := http.Dir("../static")
	fileServer := http.FileServer(dir)

	http.Handle("/", fileServer)
	http.HandleFunc("/chat", elizaInterface)

	http.ListenAndServe(":8080", nil)
}
