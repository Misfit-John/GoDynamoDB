package GoDynamoDB

import (
	"testing"
)

func Test_PutItem(t *testing.T) {
	e, err := GetDBInstance().GetPutItemExcutor(&TestStruct{Name: "John", Id: "123"})
	if err != nil {
		t.Error(err.Error())
	}
	execErr := e.exec()
	if execErr != nil {
		t.Error(execErr.Error())
	}
}

func init() {
	InitLocalDBInstance("http://127.0.0.1:8000")
}
