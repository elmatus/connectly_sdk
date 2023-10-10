package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	// Initialize Connectly client
	apiKey := "<YOUR_API_KEY>"
	baseURL := "https://cde176f9-7913-4af7-b352-75e26f94fbe3.mock.pstmn.io"
	client := NewConnectlyClient(apiKey, baseURL)

	// Path to the CSV file
	csvFilePath := "connectly_campaign.csv"

	// Parse CSV file
	request, err := parseCSV(csvFilePath)
	if err != nil {
		fmt.Printf("Error parsing CSV: %v\n", err)
		return
	}

	// Send batch campaign
	response, err := client.BatchSendCampaign(ctx, request)
	if err != nil {
		fmt.Printf("Error sending campaign: %v\n", err)
		return
	}

	fmt.Printf("Campaign sent successfully. Response: %+v\n", response)
}

func parseCSV(csvFilePath string) (*BatchSendMessageRequest, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		req := &BatchSendMessageRequest{
			Number:       line[1], // Assuming the phone number is in the second column
			TemplateName: line[2], // Assuming the template name is in the third column
			Language:     "en",    // Set the language accordingly
		}
		return req, nil
	}

	return nil, fmt.Errorf("no data found in CSV")
}
