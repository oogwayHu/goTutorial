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

func postItem(item post) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshalling item", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Posts"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		fmt.Println("Error putting item into database", err)
		return err
	}

	return nil
}
