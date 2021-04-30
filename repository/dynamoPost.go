package repository

import (
	"fmt"
	"log"

	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func PostDBitem(dbInsert models.DBparser) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	fmt.Printf("valor de spaceShip: %#v\n\n", dbInsert)

	av, err := dynamodbattribute.MarshalMap(dbInsert)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	fmt.Printf("valor de av: %+v\n\n", av)

	// Create item in table
	tableName := "mercadoLivre"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Println("Successfully added '", tableName)

}
