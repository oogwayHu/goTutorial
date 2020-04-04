package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type post struct {
	ID   int64  `json:"Id"`
	Data string `json:"Data"`
}

func show() error {
	test := post{
		ID:   6,
		Data: "Heelo world",
	}

	err := postItem(test)
	if err != nil {
		fmt.Println("Error posting item", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(show)
}
