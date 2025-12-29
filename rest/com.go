package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type IMain struct {
	id    string
	token string
	ip    string
}

func Init(id string, token string, ip string) *IMain {
	return &IMain{
		id:    id,
		token: token,
		ip:    ip,
	}
}

func comGet(auth *IMain, endpoint string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/%s", auth.ip, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+auth.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errClient := client.Do(req)
	if errClient != nil {
		return nil, errClient
	}
	defer resp.Body.Close()

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errBody
	}

	return body, nil
}

func comPost(auth *IMain, endpoint string, payload string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/%s", auth.ip, endpoint)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+auth.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errClient := client.Do(req)
	if errClient != nil {
		return nil, errClient
	}
	defer resp.Body.Close()

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errBody
	}

	return body, nil
}
