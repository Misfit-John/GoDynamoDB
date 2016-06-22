package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

type ScanExecutor struct {
	input     *dynamodb.ScanInput
	db        *dynamodb.DynamoDB
	prototype *ReadModel
	ret       []ReadModel
	count     int64
}

func (db GoDynamoDB) GetScanExecutor(i ReadModel) (*ScanExecutor, error) {
	ret := &ScanExecutor{
		db:        db.db,
		prototype: &i,
		count:     0,
	}
	input := &dynamodb.ScanInput{
		TableName: aws.String(i.GetTableName()),
	}
	ret.input = input

	return ret, nil
}

func (q *ScanExecutor) Exec() (*ScanExecutor, error) {
	resp, err := q.db.Scan(q.input)
	if nil != err {
		return nil, err
	}

	if nil == q.ret {
		q.ret = make([]ReadModel, 0)
	}

	for i := 0; i < len(resp.Items); i++ {
		rspi := resp.Items[i]
		orgPrototype := *q.prototype
		err := decode(rspi, &orgPrototype)
		if err != nil {
			return nil, err
		}
		q.ret = append(q.ret, orgPrototype)
	}
	q.count = *resp.Count
	if nil != resp.LastEvaluatedKey {
		q.input.ExclusiveStartKey = resp.LastEvaluatedKey
		return q, nil
	} else {
		return nil, nil
	}
}

func (q *ScanExecutor) GetRet() []ReadModel {
	return q.ret
}

func (q *ScanExecutor) GetCount() int64 {
	return q.count
}
