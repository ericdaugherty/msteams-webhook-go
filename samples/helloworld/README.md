# Microsoft Teams GoLang WebHook Sample

This sample uses the msteams-webhook-go library to build a simple 'Hello World' BOT WebHook for MS Teams.

*Instructions assume OS X and the Amazon CLI*

## Build the sample:
```
GOOS=linux go build -o main
```

## Package:
```
zip deployment.zip main
```

## Update
You can create the lambda function by executing:
```
aws lambda create-function \
--region us-east-1 \
--function-name MSTeamsHelloWorld \
--zip-file fileb://./deployment.zip \
--runtime go1.x \
--role arn:aws:iam::<account-id>:role/<role> \
--handler main
```

aws lambda create-function \
--region us-east-1 \
--function-name MSTeamsHelloWorld \
--zip-file fileb://./deployment.zip \
--runtime go1.x \
--role arn:aws:iam::231107391174:role/lambda_basic_execution \
--handler main

## Configure API Gateway

Open the Lambda console in AWS and select your newly created function. Add an 'API Gateway' trigger, select a name for it, and change the permission to 'Open'.

Once saved it should display the URL to invoke the service.  You can test it with curl:

```
curl -X POST -H "Content-Type: application/json" -d @TestRequest.json https://<id>.execute-api.us-east-1.amazonaws.com/prod/<name>
```

It should return:
```
{"type":"message","message":"Hello Tim Jones"}
```

## Update
You can update the lambda function by re-compailing and packaging and then executing:
```
aws lambda update-function-code \
  --function-name MSTeamsHelloWorld \
  --zip-file fileb://./deployment.zip
```
