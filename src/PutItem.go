package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

type PutItemExecutor struct {
	input *dynamodb.PutItemInput
	db    *dynamodb.DynamoDB
	res   WriteModel
}

func (db GoDynamoDB) GetPutItemExcutor(i WriteModel) (*PutItemExecutor, error) {
	attMap, err := encode(i)
	if nil != err {
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:         attMap,
		TableName:    aws.String(i.GetTableName()),
		ReturnValues: aws.String("NONE"),
	}

	return &PutItemExecutor{input: params, db: db.db}, nil
}

func (e *PutItemExecutor) Exec() error {
	resp, err := e.db.PutItem(e.input)

	if err != nil {
		return NewDynError(resp.String())
	}
	return nil

}
