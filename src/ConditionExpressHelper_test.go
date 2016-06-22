package GoDynamoDB

import (
	"testing"
)

func Test_QueryCondHelper(t *testing.T) {
	helperKey := Eq("id", ":v_id")
	helperRange := Eq("#a", ":v_index")
	finalHelper := And(helperKey, helperRange)
	if finalHelper.str != "id = :v_id AND #a = :v_index" {
		t.Errorf("wrong str:%s", finalHelper.str)
	}

}
