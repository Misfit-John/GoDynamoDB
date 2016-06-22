package GoDynamoDB

import (
	"testing"
)

func Test_Scan(t *testing.T) {
	exec, createErr := GetDBInstance().GetScanExecutor(&QueryTest{})
	if createErr != nil {
		t.Error(createErr.Error())
	}

	next, err := exec.Exec()
	if nil != err || nil != next {
		t.Error(err.Error())
	}

	ret := exec.GetRet()
	for i := 0; i < len(ret); i++ {
		reti := ret[i].(*QueryTest)
		if reti.Id != "123" {
			t.Error("wrong id")
		}
	}
	if len(ret) != 3 {
		t.Error("wrong ret num")
	}
}
