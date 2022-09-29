package quotable

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func NewClient() Client {
	return Client{}
}

const addr = "https://api.quotable.io"

type Client struct {
}

func (c Client) Quote(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// TODO make tags configurable
	req, err := http.NewRequest(http.MethodGet, addr+"/random?tags=wisdom", nil)
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
