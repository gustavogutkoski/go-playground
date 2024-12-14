package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Words Struct for parsing JSON that is defined by API I was testing
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}

	uri, err := url.ParseRequestURI(args[1])
	if err != nil {
		log.Fatal("URI parse fail", err)
	}

	get, err := http.Get(uri.String())
	if err != nil {
		log.Fatal("Request fail", err)
	}

	defer get.Body.Close()

	body, err := io.ReadAll(get.Body)

	if err != nil {
		log.Fatal("Read fail", err)
	}

	var words Words

	// Body from request was parsed into `words`
	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal("Unmarshal fail", err)
	}

	fmt.Printf("HTTP Status Code: %d\nBody: %s\n", get.StatusCode, body)
}
