package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sqssvc := sqs.New(session.New())

	params := &sqs.ListQueuesInput{
		QueueNamePrefix: aws.String("prod-"),
	}
	sqs_resp, err := sqssvc.ListQueues(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, url := range sqs_resp.QueueUrls {
		fmt.Println(*url)
	}
}