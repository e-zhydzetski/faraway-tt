package quoter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func NewQuotableClient() QuotableClient {
	return QuotableClient{}
}

type QuotableClient struct {
}

func (q QuotableClient) Quote(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// TODO make tags configurable
	req, err := http.NewRequest(http.MethodGet, "https://api.quotable.io/random?tags=wisdom", nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	var body struct {
		Content string `json:"content"`
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.Content, nil
}
