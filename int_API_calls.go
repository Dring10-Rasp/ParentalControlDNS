package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
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

func instAPICall(url, method string, payload any, headers []Header) API_OUTPUT {
	// THIS IS FOR TESTING ONLYY <====== >:(
	transportClient := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transportClient,
	}
	var parsedJson io.Reader
	if payload != nil {
		jsonBytes, errorPayload := json.Marshal(payload)
		if errorPayload != nil {
			return API_OUTPUT{[]byte{}, errorPayload, -1}
		}

		parsedJson = bytes.NewReader(jsonBytes)
	} else {
		parsedJson = nil
	}
	requestAPI, errorRequest := http.NewRequest(method, url, parsedJson)

	if errorRequest != nil {
		//log.Fatalf("Failed to create request: %v", errorRequest)
		return API_OUTPUT{[]byte{}, errorRequest, -2}
	}
	for _, header := range headers {
		requestAPI.Header.Set(header.key, header.val)
	}
	responseAPI, errorResponse := client.Do(requestAPI)
	if errorResponse != nil {
		//log.Fatalf("Failed to send request: %v", errorResponse)
		return API_OUTPUT{[]byte{}, errorResponse, -3}
	}
	responseContent, errorContent := io.ReadAll(responseAPI.Body)
	if errorContent != nil {
		//log.Fatalf("Failed to read request: %v", errorContent)
		return API_OUTPUT{[]byte{}, errorContent, responseAPI.StatusCode}
	}

	defer responseAPI.Body.Close()

	return API_OUTPUT{responseContent, nil, responseAPI.StatusCode}

}
