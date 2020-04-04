package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type post struct {
	Id   int64  `json:"id"`
	Data string `json:"data"`
}

func show() (*post, error) {
	p, err := getItem("1")
	if err != nil {
		return nil, err
	}

	return p, nil
}

func main() {
	lambda.Start(show)
}
