package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

var usernames []MyEvent

func init() {

	username1 := MyEvent{Name: "Mike"}
	username2 := MyEvent{Name: "Kimmy"}
	username3 := MyEvent{Name: "Danny"}
	username4 := MyEvent{Name: "Oliver"}

	usernames = append(usernames, username1)
	usernames = append(usernames, username2)
	usernames = append(usernames, username3)
	usernames = append(usernames, username4)

	fmt.Println(usernames)
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	finduser := name.Name
	for _, user := range usernames {
		if user.Name == finduser {

			return user.Name, nil
		}
	}
	errmsg := "user " + finduser + " not found!"
	return errmsg, nil
}

func main() {
	lambda.Start(HandleRequest)
}
