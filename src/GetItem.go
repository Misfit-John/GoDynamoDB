package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

type GetItemExecutor struct {
	input *dynamodb.GetItemInput
	db    *dynamodb.DynamoDB
	ret   *ReadModel
}

func (db GoDynamoDB) GetGetItemExecutor(i ReadModel) (*GetItemExecutor, error) {
	//actually we need a func called encode key
	key, err := encodeKeyOnly(i, i.GetTableName())
	if err != nil {
		return nil, err
	}
	params := &dynamodb.GetItemInput{
		TableName: aws.String(i.GetTableName()),
		Key:       key,
	}

	return &GetItemExecutor{input: params, db: db.db, ret: &i}, nil
}

func (e *GetItemExecutor) Exec() error {
	resp, err := e.db.GetItem(e.input)

	if err != nil {
		return NewDynError(resp.String())
	}
	decode(resp.Item, e.ret)
	return nil

}

type BathGetItemExecutor struct {
	input *dynamodb.GetItemInput
	db    *dynamodb.DynamoDB
	ret   *ReadModel
}

func (db GoDynamoDB) BathGetItem(is []ReadModel) error {
	return nil

}
