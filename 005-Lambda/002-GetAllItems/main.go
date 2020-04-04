package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type post struct {
	ID   int64  `json:"id"`
	Data string `json:"data"`
}

func show() (*[]post, error) {
	// p, err := getItem("1")
	p, err := getItems()
	if err != nil {
		fmt.Println("Error getting items", err)
		return nil, err
	}

	return p, nil
}

func main() {
	lambda.Start(show)
}
