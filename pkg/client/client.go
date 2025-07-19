package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "https://api.spacetraders.io/v2"

type Client struct {
	http *http.Client
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		http: http.DefaultClient,
		Token: token,
	}
}

func (c *Client) Get(path string, dest any) error {
	req, err := http.NewRequest(http.MethodGet, BaseURL+path, nil)
	if err != nil {
		return err
	}

	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("GET %s Response: %s\n", path, string(b))

	return json.Unmarshal(b, dest)
}

func (c *Client) Post(path string, body any, dest any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, BaseURL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("POST %s Response: %s\n", path, string(b))

	return json.Unmarshal(b, dest)
}
