package GoDynamoDB

import (
	"testing"
)

func Test_UpdateHelper(t *testing.T) {
	setExp := SetValue(SetValueMap{SetName: "Name", SetValue: ":v_name"})
	if setExp.str != "SET Name = :v_name" {
		t.Errorf("wrong str :%s", setExp.str)
	}
}
