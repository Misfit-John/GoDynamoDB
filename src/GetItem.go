package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

func (db GoDynamoDB) GetItem(i ReadModel) error {
	//actually we need a func called encode key
	key, err := encode(i)
	if err != nil {
		return err
	}
	params := &dynamodb.GetItemInput{
		TableName:      aws.String(i.GetTableName()),
		ConsistentRead: aws.Bool(i.IsConsistentRead()),
		Key:            key,
	}
	resp, err := db.db.GetItem(params)

	if err != nil {
		return NewDynError(resp.String())
	}

	return nil
}
