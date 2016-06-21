package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import "strings"
import "fmt"

type QueryCondExpressHelper struct {
	str        string
	valueMap   map[string]*dynamodb.AttributeValue
	expressMap map[string]*string
}

const (
	OperandConst  = 1
	OperandTop    = 2
	OperandNested = 3
)

type Operand struct {
	value  interface{}
	opType int
}

func NewQueryCondExpress() *QueryCondExpressHelper {
	return &QueryCondExpressHelper{
		str:      "",
		valueMap: make(map[string]*dynamodb.AttributeValue),
	}
}

func (q *QueryCondExpressHelper) Eq(l, r Operand) *QueryCondExpressHelper {

}
func (q *QueryCondExpressHelper) NE(l, r Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) LT(l, r Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) LE(l, r Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) GT(l, r Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) GE(l, r Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) In(l, r ...Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) Between(l, rl, rr Operand) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) Attribute_exist(path string) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) Attribute_not_exist(path string) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) Attribute_type(path, t string) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) Begins_with(path, prefix string) *QueryCondExpressHelper {

}
func (q *QueryCondExpressHelper) Conatins(path, str string) *QueryCondExpressHelper {

}
func (q *QueryCondExpressHelper) Size(path string) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) And(l, r *QueryCondExpressHelper) *QueryCondExpressHelper {

}

func (q *QueryCondExpressHelper) Or(l, r *QueryCondExpressHelper) *QueryCondExpressHelper {

}
func (q *QueryCondExpressHelper) Not(q *QueryCondExpressHelper) *QueryCondExpressHelper {

}
func (q *QueryCondExpressHelper) Wrap(l, r *QueryCondExpressHelper) *QueryCondExpressHelper {

}
func (q *QueryCondExpressHelper) removeSharp(key string) string {
	return strings.Replace(key, "#", "", -1)
}
