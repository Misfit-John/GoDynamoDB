package GoDynamoDB

import (
	"testing"
)

func (*TestStruct) GetTableName() string {
	return "Test"
}

func Test_PutItem(t *testing.T) {
	err := GetDBInstance().PutItem(&TestStruct{Name: "John", Id: "123"})
	if err != nil {
		t.Error("can not put item")
	}
}

func Test_Error(t *testing.T) {
	t.Error("error!")
}

func init() {
	InitLocalDBInstance("http://127.0.0.1:8000")
}
