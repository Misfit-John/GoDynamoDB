package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import "time"

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
	input        *dynamodb.BatchGetItemInput
	db           *dynamodb.DynamoDB
	ret          *[]ReadModel
	modelSiteMap map[string][]int
	curIt        map[string]int
}

func (db GoDynamoDB) GetBatchGetItemExecutor(is []ReadModel) (*BatchGetItemExecutor, error) {
	out := &dynamodb.BatchGetItemInput{
		RequestItems: make(map[string]*dynamodb.KeysAndAttributes),
	}
	ret := &BatchGetItemExecutor{
		db:           db.db,
		ret:          &is,
		modelSiteMap: make(map[string][]int),
		curIt:        make(map[string]int),
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
			ret.modelSiteMap[tableName] = make([]int, 0)
			ret.curIt[tableName] = 0
		}
		out.RequestItems[tableName].Keys = append(out.RequestItems[tableName].Keys, key)
		ret.modelSiteMap[tableName] = append(ret.modelSiteMap[tableName], i)
	}

	ret.input = out

	return ret, nil
}

func (e *BatchGetItemExecutor) Exec() error {
	rsp, err := e.db.BatchGetItem(e.input)

	if err != nil {
		return err
	}

	for key, value := range rsp.Responses {
		curItInTable := e.curIt[key]

		for i := 0; i < len(value); i++ {
			itInRet := e.modelSiteMap[key][curItInTable]
			decode(value[i], &(*e.ret)[itInRet])
			curItInTable++
		}
		e.curIt[key] = curItInTable
	}

	if len(rsp.UnprocessedKeys) != 0 {
		// this is a big feature that needs to test
		e.input.RequestItems = rsp.UnprocessedKeys
		time.Sleep(100 * time.Millisecond)
		return e.Exec()
	}

	return nil
}
