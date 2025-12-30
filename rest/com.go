package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IMain struct {
	Id     string
	token  string
	secure bool
	ip     string
}

func Init(id string, token string, ip string, secure bool) *IMain {
	return &IMain{
		Id:     id,
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

func comGet[T any](auth *IMain, endpoint string) (T, error) {
	var result T

	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	return com[T](req, auth.token)
}

func comDelete[T any](auth *IMain, endpoint string) (T, error) {
	var result T

	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return result, err
	}

	return com[T](req, auth.token)
}

func comPost[T any](auth *IMain, endpoint string, payload string) (T, error) {
	var result T
	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return result, err
	}

	return com[T](req, auth.token)
}

func IcomPost[T any](auth *IMain, endpoint string, payload interface{}) (T, error) {
	var resultdef T

	url := fmt.Sprintf("%s://%s/%s", getProtocol(auth.secure), auth.ip, endpoint)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return resultdef, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return resultdef, err
	}

	return com[T](req, auth.token)
}

func com[T any](req *http.Request, token string) (T, error) {
	var result T

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errClient := client.Do(req)
	if errClient != nil {
		return result, errClient
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return result, errBody
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}
