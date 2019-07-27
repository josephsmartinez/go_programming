package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	s3svc := s3.New(session.New())
	inputparams := &s3.ListObjectsInput{
		Bucket:  aws.String("eferro-bucket"),
		MaxKeys: aws.Int64(10),
	}
	pageNum := 0
	s3svc.ListObjectsPages(inputparams, func(page *s3.ListObjectsOutput, lastPage bool) bool {
		fmt.Println("Page", pageNum)
		pageNum++
		for _, value := range page.Contents {
			fmt.Println(*value.Key)
		}
		fmt.Println("pageNum", pageNum, "lastPage", lastPage)

		// return if we should continue with the next page
		return true
	})
}