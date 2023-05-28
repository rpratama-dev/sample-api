package service

import "net/http"

type HttpClient struct {
	client *http.Client
}

func APIService() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}
