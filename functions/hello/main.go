package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler to serve static files for the base path.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := "Hello!"
	return events.APIGatewayProxyResponse{StatusCode: http.StatusFound, Body: body}, nil
}

func main() {
	lambda.Start(Handler)
}
