# Connectly Go SDK
Welcome to the Connectly Go SDK repository! This SDK allows you to interact with the Connectly API to send WhatsApp templated messages using Go programming language. It wraps the Connectly API schema and provides a convenient way to integrate Connectly services into your applications.

## Usage
### Prerequisites
Before you start using the SDK, make sure you have the following:

Connectly API Key: Obtain your API key from Connectly to authenticate your requests.

### Example
Here's an example of how you can use the SDK to send a batch campaign from a CSV file:

```
package main

import (
	"context"
	"fmt"
	"github.com/connectly-sdk"
)

func main() {
	ctx := context.Background()

	// Initialize Connectly client
	apiKey := "<YOUR_API_KEY>"
	baseURL := "https://cde176f9-7913-4af7-b352-75e26f94fbe3.mock.pstmn.io"
	client := connectly.NewConnectlyClient(apiKey, baseURL)

	// Path to the CSV file
	csvFilePath := "sample_connectly_campaign.csv"

	// Send batch campaign
	response, err := client.BatchSendCampaign(ctx, csvFilePath)
	if err != nil {
		fmt.Printf("Error sending campaign: %v\n", err)
		return
	}

	// Process response as needed
	fmt.Printf("Campaign sent successfully. Response: %+v\n", response)
}

```

Make sure to replace <YOUR_API_KEY> with your actual Connectly API key.

## Documentation
For detailed information on the Connectly API and its capabilities, please refer to the official [Connectly API Documentation](https://docs.connectly.ai/message-api/message-service-api/messageservice/send-template-message)

## Issues
If you encounter any issues with the SDK or have suggestions for improvement, please feel free to open an issue in this repository.