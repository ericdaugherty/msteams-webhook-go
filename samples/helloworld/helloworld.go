package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	teams "github.com/ericdaugherty/msteams-webhook-go"
)

type webHook struct {
}

func (w webHook) OnMessage(req teams.Request) (teams.Response, error) {
	return teams.BuildResponse("Hi " + req.FromUser.Name), nil
}

func main() {
	lambda.Start(teams.NewHandler(false, "", webHook{}))
}
