package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Header struct {
	key string
	val string
}

type API_OUTPUT struct {
	load     []byte
	err      error
	exitCode int
}

func instAPICall(url, method string, headers []Header) API_OUTPUT {
	// THIS IS FOR TESTING ONLYY <====== >:(
	transportClient := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transportClient,
	}

	requestAPI, errorRequest := http.NewRequest(method, url, nil)
	if errorRequest != nil {
		//log.Fatalf("Failed to create request: %v", errorRequest)
		return API_OUTPUT{[]byte{}, errorRequest, 0}
	}
	for _, header := range headers {
		requestAPI.Header.Set(header.key, header.val)
	}
	responseAPI, errorResponse := client.Do(requestAPI)
	if errorResponse != nil {
		//log.Fatalf("Failed to send request: %v", errorResponse)
		return API_OUTPUT{[]byte{}, errorResponse, -1}
	}
	responseContent, errorContent := io.ReadAll(responseAPI.Body)
	if errorContent != nil {
		//log.Fatalf("Failed to read request: %v", errorContent)
		return API_OUTPUT{[]byte{}, errorContent, responseAPI.StatusCode}
	}

	defer responseAPI.Body.Close()

	return API_OUTPUT{responseContent, nil, responseAPI.StatusCode}

}
func main() {
	headers := [2]Header{
		{"accept", "application/json"},
		{"sid", "Rwp7xIayacRWCCfL9PNmTQ="},
	}
	url := "https://192.168.1.10:443/api/groups"
	responseAPI := instAPICall(url, http.MethodGet, headers[:])
	if responseAPI.err != nil {
		//log.Fatalf("Failed to read body: %v", responseAPI.err)
	}

	fmt.Printf("Exit code: %d\n", responseAPI.exitCode)

	fmt.Printf("Response: %s\n", string(responseAPI.load))
	fmt.Printf("Error: %v\n", responseAPI.err)

}
