package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ConnectlyClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewConnectlyClient(apiKey, baseURL string) *ConnectlyClient {
	return &ConnectlyClient{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *ConnectlyClient) BatchSendCampaign(ctx context.Context, request *BatchSendMessageRequest) (*BatchSendMessageResponse, error) {
	endpoint := c.baseURL + "/v1/businesses/f1980bf7-c7d6-40ec-b665-dbe13620bffa/send/whatsapp_templated_messages"
	// You may need to replace {businessId} with the actual business ID in the endpoint URL.
	fmt.Println("Endpoint URL:", endpoint) // Add this line for debugging

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	fmt.Println("Request:", req)

	// Perform the HTTP request
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var apiResponse BatchSendMessageResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	// Check the status code and handle errors if necessary
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	return &apiResponse, nil
}
