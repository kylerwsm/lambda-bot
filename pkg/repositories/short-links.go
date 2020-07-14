package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kylerwsm/lambda-bot/pkg/entity"
)

// Declare DynamoDB instance which is safe for concurrent use.
var svc = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))

const tableName = "Links"

// GetShortLinks returns all the shortened links in the database.
func GetShortLinks() ([]entity.Link, error) {
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return []entity.Link{}, err
	}
	var links []entity.Link
	for _, item := range result.Items {
		var link entity.Link
		err = dynamodbattribute.UnmarshalMap(item, &link)
		if err != nil {
			return []entity.Link{}, err
		}
		links = append(links, link)
	}
	return links, nil
}
