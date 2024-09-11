package main

import (
	"fmt"
	"os"
	"zenn-cli/internal/requests"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zenn-cli <username>")
		return
	}

	username := os.Args[1]

	articles, err := requests.GetArticles(username)

	if err != nil {
		fmt.Printf("Error: %w", err)
	}

	for _, article := range articles {
		fmt.Printf("Title:%s\nUrl:%s\n", article.Title, article.URL)
	}
}
