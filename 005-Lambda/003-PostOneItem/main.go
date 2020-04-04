package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Post struct {
	ID   int64  `json:"Id"`
	Data string `json:"Data"`
}

func handler(p Post) (string, error) {
	err := postItem(p)
	if err != nil {
		return "Errorrr", err
	}

	return p.Data, nil
}

func main() {
	lambda.Start(handler)
}
