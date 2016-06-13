package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	name string
	id   string
}

func Test_encode(t *testing.T) {
	var attMap map[string]*dynamodb.AttributeValue
	var err error
	attMap, err = encode(&TestStruct{name: "John", id: "123"})
	if err != nil {
		fmt.Printf("%s\n", attMap)
	} else {
		fmt.Printf("hello\n")
	}
}
