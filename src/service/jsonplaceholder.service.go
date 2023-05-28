package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://jsonplaceholder.typicode.com"

func (s *HttpClient) GetPost() (io.ReadCloser, error) {
	resp, err := s.client.Get(fmt.Sprint(baseURL, "/posts"))
	if err != nil {
		// Handle error if the API request fails
		return nil, err
	}
	return resp.Body, nil
}

func (s *HttpClient) GetPostById(id string) (io.ReadCloser, error) {
	resp, err := s.client.Get(fmt.Sprint(baseURL,"/posts/", id))
	if err != nil {
		// Handle error if the API request fails
		return nil, err
	}
	return resp.Body, nil
}

func (s *HttpClient) StorePost(jsonData []byte) (io.ReadCloser, error) {
	resp, err := s.client.Post(fmt.Sprint(baseURL,"/posts"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// Handle error if the API request fails
		return nil, err
	}
	return resp.Body, nil
}

func (s *HttpClient) DestroyPostById(id string) (io.ReadCloser, error) {
	// Create a DELETE request
	req, err := http.NewRequest("DELETE", fmt.Sprint(baseURL,"/posts/", id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		return resp.Body, nil
	}
	return nil, err
}
