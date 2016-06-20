package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

type QueryExecutor struct {
	input     *dynamodb.QueryInput
	db        *dynamodb.DynamoDB
	prototype *ReadModel
	ret       []ReadModel
}

func (db GoDynamoDB) GetQueryExecutor(i ReadModel) (*QueryExecutor, error) {
	ret := &QueryExecutor{
		db:        db.db,
		prototype: &i,
	}
	input := &dynamodb.QueryInput{
		TableName: aws.String(i.GetTableName()),
	}

	ret.input = input
	return ret, nil
}

func (q *QueryExecutor) withKeyCondition() *QueryExecutor {
	return q
}
