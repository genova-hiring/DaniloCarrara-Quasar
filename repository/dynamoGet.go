package repository

import (
	"fmt"

	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetDBitem(satellite models.DBparser) models.DBparser {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "mercadoLivre"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(satellite.Name),
			},
			"Distance": {
				N: aws.String(fmt.Sprintf("%f", satellite.Distance)),
			},
		},
	})
	if err != nil {
		fmt.Errorf("Got error calling GetItem: %s", err)
	}

	item := models.DBparser{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return item

}
