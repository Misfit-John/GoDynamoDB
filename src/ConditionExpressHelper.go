package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "fmt"

type QueryCondExpressHelper struct {
	str        string
	valueMap   map[string]*dynamodb.AttributeValue
	expressMap map[string]string
}

func NewQueryCondExpress() *QueryCondExpressHelper {
	return &QueryCondExpressHelper{
		str:      "",
		valueMap: make(map[string]*dynamodb.AttributeValue),
	}
}

func (q *QueryCondExpressHelper) withExpressMap(nameMap map[string]string) *QueryCondExpressHelper {
	q.expressMap = nameMap
	return q
}

func (q *QueryCondExpressHelper) Eq(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf("v_", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf(q.str, " AND ")
	}
	q.str = fmt.Sprintf(key, " = ", valueExp)
	return q
}

func (q *QueryCondExpressHelper) LT(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf("v_", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf(q.str, " AND ")
	}
	q.str = fmt.Sprintf(key, " < :", valueExp)
	return q
}

func (q *QueryCondExpressHelper) GT(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf("v_", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf(q.str, " AND ")
	}
	q.str = fmt.Sprintf(key, " > :", valueExp)
	return q
}

func (q *QueryCondExpressHelper) GE(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf("v_", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf(q.str, " AND ")
	}
	q.str = fmt.Sprintf(key, " >= :", valueExp)
	return q
}

func (q *QueryCondExpressHelper) LE(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf("v_", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf(q.str, " AND ")
	}
	q.str = fmt.Sprintf(key, " <= :", valueExp)
	return q
}

func (q *QueryCondExpressHelper) BeginWith(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf("v_", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf(q.str, " AND ")
	}
	q.str = fmt.Sprintf("begins_with ( ", key, ",", valueExp, ")")
	return q
}
