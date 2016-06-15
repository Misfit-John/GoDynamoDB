package GoDynamoDB

import (
	"fmt"
	"testing"
)

func (*TestStruct) GetTableName() string {
	return "Test"
}

func Test_PutItem(t *testing.T) {
	err := PutItem(&TestStruct{name: "John", id: "123"})
	if err != nil {
		fmt.Printf("error:%s\n", err)
	} else {
		fmt.Printf("success\n")
	}
}
