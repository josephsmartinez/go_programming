# AWS Lambda GO Notes

## Compiling and Deployments

1. First create a Lambda function deployment package, a .zip file consisting of your code (a Go executable) and any dependencies.

COMPILE
Setting GOOS to linux ensures that the compiled executable is compatible with the Go runtime, even if you compile it in a non-Linux environment.
> GOOS=linux go build main.go

DEPLOY
Create a deployment package by packaging the executable in a ZIP file, and use the AWS CLI to create a function. The handler parameter must match the name of the executable containing your handler.
> zip function.zip main
> aws lambda create-function --function-name my-function --runtime go1.x  --zip-file fileb://function.zip --handler main --role arn:aws:iam::123456789012:role/execution_role

