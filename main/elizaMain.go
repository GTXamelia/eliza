package main

// Imports used
import (
	"fmt"
	"net/http"
	"time"
	".."
)

// Eliza interface that communicates information from the web application to the GO language and then outputs result to the webpage
func elizaInterface(w http.ResponseWriter, r *http.Request) {

	// Gets time
	t := time.Now()
	
	// Takes in user input from webpage makes user input lowercase for easy read
	input := r.URL.Query().Get("input")

	// Get output from eliza using function
	output := eliza.ReplyTo(input)

	// Outputs result to the webpage
	fmt.Fprintf(w, output)

	// Ouputs user input to the console to track user interaction
	// Also outputs bots reaction to console
	// N.B This was useful during development to see how users interacted with the bot
	// helped formulate responses and add more reaction for the bot
	fmt.Printf("---%d:%d:%d %d/%d/%d---\n",t.Hour(), t.Minute(), t.Second(), t.Day(), t.Month(), t.Year())
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
	http.HandleFunc("/elizaPrompt", elizaInterface)
	http.ListenAndServe(":8080", nil)
}
