package GoDynamoDB

import (
	"testing"
)

func Test_Query(t *testing.T) {
	helper := NewQueryCondExpress().Eq("id", "123").GT("#a", 11).AddExpressMap("index", "#a")
	exec, createErr := GetDBInstance().GetQueryExecutor(&QueryTest{})
	if createErr != nil {
		t.Error(createErr.Error())
	}
	err := exec.WithKeyCondition(*helper).Exec()
	if nil != err {
		t.Error(err.Error())
	}

	ret := exec.GetRet()
	for i := 0; i < len(ret); i++ {
		reti := ret[i].(*QueryTest)
		if reti.Id != "123" {
			t.Error("wrong id")
		}
	}
}
