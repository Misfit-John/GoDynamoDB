package GoDynamoDB

import (
	"testing"
)

func Test_GetItem(t *testing.T) {
	toGet := &TestStruct{Id: "123"}
	e, err := GetDBInstance().GetGetItemExecutor(toGet)
	if err != nil {
		t.Error(err.Error())
	}

	execErr := e.Exec()
	if nil != execErr {
		t.Error(execErr.Error())
	}

	if toGet.Name != "John" {
		t.Error("Get fail")
	}
}

func Test_BatchGetItem(t *testing.T) {
	toGet := make([]ReadModel, 3)
	toGet[0] = &TestStruct{Id: "123"}
	toGet[1] = &TestStruct2nd{Id: "1234"}
	toGet[2] = &TestStruct{Id: "1235"}
	e, err := GetDBInstance().GetBatchGetItemExecutor(toGet)
	if err != nil {
		t.Error(err.Error())
	}

	execErr := e.Exec()
	if nil != execErr {
		t.Error(execErr.Error())
	}

	if toGet[0].(*TestStruct).Name != "John" {
		t.Error("Get fail")
	}
}
