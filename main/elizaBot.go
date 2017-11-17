package main

import (
	"fmt"
	"net/http"
	"strings"
	".."
)

// Eliza interface that communicates information from the web application to the GO language and then outputs result to the webpage
func elizaInterface(w http.ResponseWriter, r *http.Request) {
	
	// Takes in user input from webpage makes user input lowercase for easy read
	input := strings.ToLower(r.URL.Query().Get("input"))

	// Get output from eliza using function
	output := eliza.ReplyTo(input)

	// Outputs result to the webpage
	fmt.Fprintf(w, output)

	// Ouputs user input to the console to track user interaction
	// Also outputs bots reaction to console
	// N.B This was useful during development to see how users interacted with the bot
	// helped formulate responses and add more reaction for the bot
	fmt.Println("User: "+input)
	fmt.Println("Bot: "+output+"\n")
}

// Main function handles HTTP requests
func main() {
	// Webpage Directory
	dir := http.Dir("../webpage")
	fileServer := http.FileServer(dir)

	// HTTP request handler
	// Directs HTTP requests
	http.Handle("/", fileServer)
	http.HandleFunc("/chat", elizaInterface)
	http.ListenAndServe(":8080", nil)
}
