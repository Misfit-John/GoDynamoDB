package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import "fmt"

type QueryCondExpressHelper struct {
	str        string
	valueMap   map[string]*dynamodb.AttributeValue
	expressMap map[string]*string
}

func NewQueryCondExpress() *QueryCondExpressHelper {
	return &QueryCondExpressHelper{
		str:      "",
		valueMap: make(map[string]*dynamodb.AttributeValue),
	}
}

func (q *QueryCondExpressHelper) AddExpressMap(org, exp string) *QueryCondExpressHelper {
	if nil == q.expressMap {
		q.expressMap = make(map[string]*string)
	}
	q.expressMap[exp] = aws.String(org)
	return q
}

func (q *QueryCondExpressHelper) Eq(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf(":v_%s", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}
	q.str = fmt.Sprintf("%s%s = %s", q.str, key, valueExp)

	return q
}

func (q *QueryCondExpressHelper) LT(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf(":v_%s", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}
	q.str = fmt.Sprintf("%s%s < %s", q.str, key, valueExp)
	return q
}

func (q *QueryCondExpressHelper) GT(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}

	valueExp := fmt.Sprintf(":v_%s", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}
	q.str = fmt.Sprintf("%s%s > %s", q.str, key, valueExp)

	return q
}

func (q *QueryCondExpressHelper) GE(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf(":v_%s", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}
	q.str = fmt.Sprintf("%s%s >= %s", q.str, key, valueExp)
	return q
}

func (q *QueryCondExpressHelper) LE(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}
	valueExp := fmt.Sprintf(":v_%s", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}
	q.str = fmt.Sprintf("%s%s <= %s", q.str, key, valueExp)

	return q
}

func (q *QueryCondExpressHelper) BeginWith(key string, value interface{}) *QueryCondExpressHelper {
	att, err := encodeToQueryAtt(value)
	if err != nil {
		return nil
	}

	valueExp := fmt.Sprintf(":v_%s", key)
	q.valueMap[valueExp] = att
	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}
	q.str = fmt.Sprintf("%sbegins_with (%s, %s)", q.str, key, valueExp)

	return q
}

func (q *QueryCondExpressHelper) Between(key string, left, right interface{}) *QueryCondExpressHelper {
	attLeft, errLeft := encodeToQueryAtt(left)
	if errLeft != nil {
		return nil
	}
	valueExpLeft := fmt.Sprintf(":v_l_%s", key)
	q.valueMap[valueExpLeft] = attLeft

	attRight, errRight := encodeToQueryAtt(right)
	if errRight != nil {
		return nil
	}
	valueExpRight := fmt.Sprintf(":v_r_%s", key)
	q.valueMap[valueExpRight] = attRight

	if q.str != "" {
		q.str = fmt.Sprintf("%s AND ", q.str)
	}

	q.str = fmt.Sprintf("%s%s BETWEEN :%s AND :%s", q.str, key, valueExpLeft, valueExpRight)

	return q
}
