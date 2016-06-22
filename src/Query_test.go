package GoDynamoDB

import (
	"testing"
)

func Test_Query(t *testing.T) {
	helper1 := Eq("id", ":v_id")
	helper2 := GT("#a", ":v_index")
	helper := And(helper1, helper2).AddExpressMap("index", "#a")

	exec, createErr := GetDBInstance().GetQueryExecutor(&QueryTest{})
	if createErr != nil {
		t.Error(createErr.Error())
	}

	if helper.str != "id = :v_id AND #a > :v_index" {
		t.Errorf("wrong condition,", helper.str)
	}

	err := exec.WithKeyCondition(helper).AddValue(":v_id", "123").AddValue(":v_index", 11).Exec()
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
	if len(ret) != 2 {
		t.Error("wrong ret num")
	}
}
