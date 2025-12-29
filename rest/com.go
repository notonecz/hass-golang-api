package rest

import (
	"bytes"
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

func comGet(auth *IMain, endpoint string, secure bool) ([]byte, error) {

	url := fmt.Sprintf("%s://%s/%s", getProtocol(secure), auth.ip, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return com(req, auth.token)
}

func comPost(auth *IMain, endpoint string, payload string, secure bool) ([]byte, error) {
	url := fmt.Sprintf("%s://%s/%s", getProtocol(secure), auth.ip, endpoint)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	return com(req, auth.token)
}

func com(req *http.Request, token string) ([]byte, error) {
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

	return body, nil
}
