package GoDynamoDB

import "github.com/aws/aws-sdk-go/aws"
import "strings"
import "fmt"

//any query helper should be cached after it's created
type QueryCondExpressHelper struct {
	str        string
	expressMap map[string]*string
}

func (q *QueryCondExpressHelper) AddExpressMap(originalKey, placeHolder string) *QueryCondExpressHelper {
	if nil == q.expressMap {
		q.expressMap = make(map[string]*string)
	}
	q.expressMap[placeHolder] = aws.String(originalKey)
	return q
}

func Eq(l, r string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s = %s", l, r)
	return q
}

func NE(l, r string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s <> %s", l, r)
	return q
}

func LT(l, r string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s < %s", l, r)
	return q
}

func LE(l, r string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s <= %s", l, r)
	return q
}

func GT(l, r string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s > %s", l, r)
	return q
}

func GE(l, r string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s >= %s", l, r)
	return q
}

func In(l string, r ...string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s IN (%s)", l, strings.Join(r, ","))
	return q
}

func Between(l, rl, rr string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s BETWEEN %s AND %s", l, rl, rr)
	return q
}

func Attribute_exist(path string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("attribute_exist(%s)", path)
	return q
}

func Attribute_not_exist(path string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("attribute_not_exist(%s)", path)
	return q
}

func Attribute_type(path, t string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("attribute_type(%s, %s)", path, t)
	return q

}

func Begins_with(path, prefix string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("begins_with(%s, %s)", path, prefix)
	return q

}

func Contains(path, op string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("contains(%s, %s)", path, op)
	return q

}

func Size(path string) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("size(%s)", path)
	return q
}

func And(l, r *QueryCondExpressHelper) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s AND %s", l.str, r.str)
	return q
}

func Or(l, r *QueryCondExpressHelper) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("%s OR %s", l.str, r.str)
	return q
}

func Not(l *QueryCondExpressHelper) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("NOT %s", l.str)
	return q
}

func Wrap(l *QueryCondExpressHelper) *QueryCondExpressHelper {
	q := &QueryCondExpressHelper{}
	q.str = fmt.Sprintf("(%s)", l.str)
	return q
}

func (q *QueryCondExpressHelper) String() string {
	return q.str
}
