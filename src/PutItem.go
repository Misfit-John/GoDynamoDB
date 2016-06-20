package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import "time"

type PutItemExecutor struct {
	input *dynamodb.PutItemInput
	db    *dynamodb.DynamoDB
	// I don't sure if we are going to need this, so just keep it
	//	res   *WriteModel
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

type BatchWriteItemExecutor struct {
	input *dynamodb.BatchWriteItemInput
	db    *dynamodb.DynamoDB
}

func (db GoDynamoDB) GetBatchWriteItemExecutor(toPut, toDelete []WriteModel) (*BatchWriteItemExecutor, error) {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems:                make(map[string][]*dynamodb.WriteRequest),
		ReturnItemCollectionMetrics: aws.String("NONE"),
	}
	ret := &BatchWriteItemExecutor{db: db.db}

	// append put request
	if len(toPut) > 0 {
		for i := 0; i < len(toPut); i++ {
			tableName := toPut[i].GetTableName()
			if _, ok := input.RequestItems[tableName]; !ok {
				input.RequestItems[tableName] = make([]*dynamodb.WriteRequest, 0)
			}
			attMap, err := encode(toPut[i])
			if nil != err {
				return nil, err
			}
			input.RequestItems[tableName] = append(input.RequestItems[tableName], &dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: attMap,
				},
			})
		}
	}

	// append delete request
	if len(toDelete) > 0 {
		for i := 0; i < len(toDelete); i++ {
			tableName := toDelete[i].GetTableName()
			if _, ok := input.RequestItems[tableName]; !ok {
				input.RequestItems[tableName] = make([]*dynamodb.WriteRequest, 0)
			}
			attMap, err := encodeKeyOnly(toDelete[i], tableName)
			if nil != err {
				return nil, err
			}
			input.RequestItems[tableName] = append(input.RequestItems[tableName], &dynamodb.WriteRequest{
				DeleteRequest: &dynamodb.DeleteRequest{
					Key: attMap,
				},
			})
		}
	}

	ret.input = input

	return ret, nil
}

func (e *BatchWriteItemExecutor) Exec() error {
	rsp, err := e.db.BatchWriteItem(e.input)

	if err != nil {
		return err
	}

	if len(rsp.UnprocessedItems) != 0 {
		e.input.RequestItems = rsp.UnprocessedItems
		time.Sleep(100 * time.Millisecond)
		return e.Exec()
	}
	return nil
}
