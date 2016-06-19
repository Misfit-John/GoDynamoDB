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

type BatchGetItemExecutor struct {
	input *dynamodb.BatchGetItemInput
	db    *dynamodb.DynamoDB
	ret   *[]ReadModel
}

func (db GoDynamoDB) GetBatchGetItemExecutor(is []ReadModel) (*BatchGetItemExecutor, error) {
	out := &dynamodb.BatchGetItemInput{
		RequestItems: make(map[string]*dynamodb.KeysAndAttributes),
	}
	isLen := len(is)
	for i := 0; i < isLen; i++ {
		tableName := is[i].GetTableName()
		key, err := encodeKeyOnly(is[i], tableName)
		if err != nil {
			return nil, err
		}

		if _, ok := out.RequestItems[tableName]; !ok {
			out.RequestItems[tableName] = &dynamodb.KeysAndAttributes{}
		}
		out.RequestItems[tableName].Keys = append(out.RequestItems[tableName].Keys, key)
	}

	return &BatchGetItemExecutor{input: out, db: db.db, ret: &is}, nil
}
