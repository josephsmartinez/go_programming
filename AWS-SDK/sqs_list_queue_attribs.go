package main

import (
	"fmt"
	"strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sqssvc := sqs.New(session.New())
	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String("https://sqs.<region>.amazonaws.com/<acc_number>/<queue_name>"),
		AttributeNames: []*string{
			aws.String("ApproximateNumberOfMessages"),
			aws.String("ApproximateNumberOfMessagesDelayed"),
			aws.String("ApproximateNumberOfMessagesNotVisible"),
		},
	}
	resp, _ := sqssvc.GetQueueAttributes(params)
	for attrib, _ := range resp.Attributes {
		prop := resp.Attributes[attrib]
		i, _ := strconv.Atoi(*prop)
		fmt.Println(attrib, i)
	}
}