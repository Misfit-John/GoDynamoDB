package GoDynamoDB

import (
	"testing"
)

func Test_UpdateItem(t *testing.T) {
	setExp := SetValue(SetValueMap{SetName: "FName", SetValue: ":v_name"})
	exector, getErr := GetDBInstance().GetUpdateItemExecutor(&TestStruct{Id: "999"})
	if nil != getErr {
		t.Error("get fail")
	}
	exector.SetUpdateExpress(setExp).AddValue(":v_name", "hello")
	err := exector.Exex()
	if nil != err {
		t.Error(err.Error())
	}

}
