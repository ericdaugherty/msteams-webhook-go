# GoLang Outgoing WebHook Library

This is a small library to simplify the task of writing BOT Outgoing WebHooks for Microsoft Teams, as outlined in their documentation https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/outgoingwebhook

The sample project (samples/helloworld) provide full documentation on using and deploying this library. In short, you need to implement a main function and an OnMessage function as follows:

```
import (
	"github.com/aws/aws-lambda-go/lambda"
	teams "github.com/ericdaugherty/msteams-webhook-go"
)

type webHook struct {
}

func (w webHook) OnMessage(req teams.Request) (teams.Response, error) {
	return teams.BuildResponse("Hello " + req.FromUser.Name), nil
}

func main() {
	lambda.Start(teams.NewHandler(false, "", webHook{}))
}
```

The main function starts the Lambda API and registers your callback. Then you simply process the request and return the response, using the BuildResponse helper method.

This library also supports HMAC authenctication. Simply pass in 'true' and the Base64 encoded String returned by Microsoft Teams when you register your callback URL.