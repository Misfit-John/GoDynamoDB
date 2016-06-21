package GoDynamoDB

import (
	"testing"
)

func Test_QueryCondHelper(t *testing.T) {
	helper := NewQueryCondExpress().Eq("id", "123").Eq("index", 11)
	if helper.str != "id = :v_id AND index = :v_index" {
		t.Errorf("wrong str:%s", helper.str)
	}

	if v, ok := helper.valueMap[":v_id"]; ok == false {
		t.Error("no v_id")
	} else {
		if *v.S != "123" {
			t.Errorf("wrong v_id:", *v.S)
		}
	}

	if v, ok := helper.valueMap[":v_index"]; ok == false {
		t.Error("no v_index")
	} else {
		if *v.N != "11" {
			t.Errorf("wrong v_index:", *v.N)
		}
	}

}
