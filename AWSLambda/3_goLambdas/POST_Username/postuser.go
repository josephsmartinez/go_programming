package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	userName := name.Name
	if len(userName) > 0 {
		success := "User was posted!"
		fmt.Println(userName)
		return success, nil
	}
	failed := "User was NOT posted!"
	return failed, nil
}

func main() {
	lambda.Start(HandleRequest)
}
