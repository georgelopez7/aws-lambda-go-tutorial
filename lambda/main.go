package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// --- PARSE REQUEST BODY ---
	var body RequestBody
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       `{"message":"Invalid request body"}`,
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	// --- HANDLER LOGIC ---
	response := ResponseBody{
		Message: fmt.Sprintf("Hello, %s", body.Name),
	}

	// --- GENERATE & MARSHALL RESPONSE BODY ---
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       `{"message":"Error generating response"}`,
			StatusCode: 500,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	// --- RETURN RESPONSE BODY ---
	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(handler)
}
