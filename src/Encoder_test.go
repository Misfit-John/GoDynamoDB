package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"

import (
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
		t.Error("can not trans to AttributeValue")
	}

	if v, ok := attMap["name"]; ok == false {
		t.Error("no name att")
	} else {
		if *v.S != "John" {
			t.Errorf("wrong string:", v.S)
		}
	}

	if v, ok := attMap["id"]; ok == false {
		t.Error("no id att")
	} else {
		if *v.S != "123" {
			t.Errorf("wrong id:", v.S)
		}
	}
}
