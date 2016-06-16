package GoDynamoDB

import (
	"testing"
)

func Test_GetItem(t *testing.T) {
	err := GetDBInstance().GetItem(&TestStruct{Id: "123"})
	if err != nil {
		t.Error("can not put item")
	}
}
