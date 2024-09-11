package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Article struct {
	Title string `json:"title"`
}

type Response struct {
	Articles []Article `json:"articles"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zenn-cli <username>")
		return
	}

	username := os.Args[1]
	url := fmt.Sprintf("https://zenn.dev/api/articles?username=%s", username)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for _, article := range response.Articles {
		fmt.Printf("%s\n", article.Title)
	}
}
