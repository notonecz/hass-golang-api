package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IMain struct {
	id     string
	token  string
	secure bool
	ip     string
}

func Init(id string, token string, ip string, secure bool) *IMain {
	return &IMain{
		id:     id,
		token:  token,
		ip:     ip,
		secure: secure,
	}
}

func getProtocol(secure bool) string {
	switch secure {
	case true:
		return "https"
	case false:
		return "http"
	}
	return "http"
}

func comGet(auth *IMain, endpoint string) (interface{}, error) {
	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return com(req, auth.token)
}

func comDelete(auth *IMain, endpoint string) (interface{}, error) {
	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return com(req, auth.token)
}

func comPost(auth *IMain, endpoint string, payload string) (interface{}, error) {
	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	return com(req, auth.token)
}

func com(req *http.Request, token string) (interface{}, error) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errClient := client.Do(req)
	if errClient != nil {
		return nil, errClient
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errBody
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}
