package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

type UpdateItemExecutor struct {
	input *dynamodb.UpdateItemInput
	db    *dynamodb.DynamoDB
}

func (db GoDynamoDB) GetUpdateItemExecutor(i WriteModel) (*UpdateItemExecutor, error) {
	ret := &UpdateItemExecutor{
		db: db.db,
	}

	input := &dynamodb.UpdateItemInput{}
	input.TableName = aws.String(i.GetTableName())
	ret.input = input

	return ret, nil
}

func (e *UpdateItemExecutor) SetUpdateExpress(exp UpdateCondExpressHelper) *UpdateItemExecutor {
	if nil == e.input.ExpressionAttributeNames {
		e.input.ExpressionAttributeNames = make(map[string]*string)
	}
	for key, value := range exp.expressMap {
		e.input.ExpressionAttributeNames[key] = value
	}
	e.input.UpdateExpression = aws.String(exp.str)
	return e
}
