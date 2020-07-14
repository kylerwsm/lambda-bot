package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kylerwsm/lambda-bot/pkg/services"
	"github.com/kylerwsm/lambda-bot/pkg/util"
)

// Handler is our lambda handler to serve static files for the base path.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := util.SetUpdateFromBody(request.Body)
	if err != nil {
		log.Println(err)
		return util.ResponseBadRequest(), err
	}

	command := util.GetText()
	resp, err := services.CreateResponse(command)
	if err != nil {
		log.Println(err)
		return util.ResponseBadRequest(), err
	}

	chatID := util.GetChatID()
	err = services.SendResponse(chatID, resp)
	if err != nil {
		log.Println(err)
		return util.ResponseBadRequest(), err
	}

	return util.ResponseOK(), nil
}

func main() {
	lambda.Start(Handler)
}
