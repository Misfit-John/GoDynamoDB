package GoDynamoDB

import (
	"testing"
)

func Test_CreateTable(t *testing.T) {
	exe, err := GetDBInstance().GetCreateTableExecutor(&CreateTest{Id: "123"})
	if nil != err {
		t.Error(err.Error())
	}
	Eerr := exe.Exec()
	if nil != Eerr {
		t.Logf(exe.input.GoString())
		t.Error(Eerr.Error())
	}
}
