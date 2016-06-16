package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

func (db GoDynamoDB) PutItem(i WriteModel) error {
	attMap, err := encode(i)
	if nil != err {
		return err
	}
	params := &dynamodb.PutItemInput{
		Item:         attMap,
		TableName:    aws.String(i.GetTableName()),
		ReturnValues: aws.String("NONE"),
	}
	resp, err := db.db.PutItem(params)

	if err != nil {
		return NewDynError(resp.String())
	}

	return nil
}
