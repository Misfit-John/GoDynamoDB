package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"

import (
	"testing"
)

func Test_decode(t *testing.T) {
	var attMap map[string]*dynamodb.AttributeValue
	var err error
	attMap, err = encode(&TestStruct{Name: "John", Id: "123"})
	if err != nil {
		t.Error("can not trans to AttributeValue")
	}

	if v, ok := attMap["Name"]; ok == false {
		t.Error("no Name att")
	} else {
		if *v.S != "John" {
			t.Errorf("wrong string:", v.S)
		}
	}

	if v, ok := attMap["Id"]; ok == false {
		t.Error("no Id att")
	} else {
		if *v.S != "123" {
			t.Errorf("wrong Id:", v.S)
		}
	}

	outStruct := TestStruct{Name: "test", Id: "ddd"}
	decodeErr := decode(attMap, &outStruct)
	if nil != decodeErr {
		t.Error("can't decode")
	}

	if outStruct.Name != "John" {
		t.Error("wrong Name")
	}

	if outStruct.Id != "123" {
		t.Error("wrong Id")
	}
}
