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
