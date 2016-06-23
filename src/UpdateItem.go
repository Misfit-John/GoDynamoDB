package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

type UpdateItemExecutor struct {
	input *dynamodb.UpdateItemInput
	db    *dynamodb.DynamoDB
}

func (db GoDynamoDB) GetUpdateItemExecutor(i WriteModel) (*UpdateItemExecutor, error) {
	key, err := encodeKeyOnly(i, i.GetTableName())
	if err != nil {
		return nil, err
	}
	ret := &UpdateItemExecutor{
		db: db.db,
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(i.GetTableName()),
		Key:       key,
	}
	ret.input = input

	return ret, nil
}

func (e *UpdateItemExecutor) SetUpdateExpress(exp *UpdateCondExpressHelper) *UpdateItemExecutor {
	if len(exp.expressMap) != 0 {
		if nil == e.input.ExpressionAttributeNames {
			e.input.ExpressionAttributeNames = make(map[string]*string)
		}
		for key, value := range exp.expressMap {
			e.input.ExpressionAttributeNames[key] = value
		}
	}
	e.input.UpdateExpression = aws.String(exp.str)
	return e
}

func (e *UpdateItemExecutor) Exex() error {
	_, err := e.db.UpdateItem(e.input)
	if nil != err {
		return err
	}

	return nil
}

func (q *UpdateItemExecutor) AddValue(express string, v interface{}) *UpdateItemExecutor {
	if nil == q.input.ExpressionAttributeValues {
		q.input.ExpressionAttributeValues = make(map[string]*dynamodb.AttributeValue)
	}
	att, err := encodeToQueryAtt(v)
	if err != nil {
		return nil
	}
	q.input.ExpressionAttributeValues[express] = att
	return q
}
