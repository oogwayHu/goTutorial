package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))

func getItems() (*[]post, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Posts"),
	}

	result, err := db.Scan(input)
	if err != nil {
		fmt.Println("Error scanning", err)
		return nil, err
	}
	if result.Items == nil {
		fmt.Println("No items in result")
		return nil, nil
	}

	p := []post{}
	// err = dynamodbattribute.UnmarshalMap(result.Items, p)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &p)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
		return nil, err
	}

	for _, v := range p {
		fmt.Println(v.ID)
		fmt.Println(v.Data)
	}

	return &p, nil
}
