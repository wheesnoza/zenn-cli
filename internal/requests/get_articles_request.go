package requests

import (
	"encoding/json"
    "fmt"
    "net/http"
)

type Article struct {
    Title string `json:"title"`
    URL   string `json:"path"`
}

type Response struct {
	Articles []Article `json:"articles"`
}

func GetArticles(username string) ([]Article, error) {
	url := fmt.Sprintf("https://zenn.dev/api/articles?username=%s", username)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch articles: %w", err)
    }
    defer resp.Body.Close()

    var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("Error decoding response: %w", err)
	}

    return response.Articles, nil
}
