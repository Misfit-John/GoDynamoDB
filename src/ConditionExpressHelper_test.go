package GoDynamoDB

import (
	"testing"
)

func Test_QueryCondHelperAnd(t *testing.T) {
	helperKey := Eq("id", ":v_id")
	helperRange := Eq("#a", ":v_index")
	finalHelper := And(helperKey, helperRange)
	if finalHelper.str != "id = :v_id AND #a = :v_index" {
		t.Errorf("wrong str:%s", finalHelper.str)
	}
}

func Test_QueryCondHelperOr(t *testing.T) {
	helperKey := NE("id", ":v_id")
	helperRange := LT("#a", ":v_index")
	finalHelper := Or(helperKey, helperRange)
	if finalHelper.str != "id <> :v_id OR #a < :v_index" {
		t.Errorf("wrong str:%s", finalHelper.str)
	}
}

func Test_QueryCondHelperIn(t *testing.T) {
	helperKey := In("id", ":v_id1", ":v_id2", ":v_id3")
	if helperKey.str != "id IN (:v_id1,:v_id2,:v_id3)" {
		t.Errorf("wrong str:%s", helperKey.str)
	}
}

func Test_QueryCondHelperSize(t *testing.T) {
	helperKey := Size("id")
	finalHelper := GT(helperKey.String(), ":v_id")
	if finalHelper.str != "size(id) > :v_id" {
		t.Errorf("wrong str:%s", finalHelper.str)
	}
}
func Test_QueryCondHelperGE(t *testing.T) {
	helperKey := GE("id", ":v_id1")
	if helperKey.str != "id >= :v_id1" {
		t.Errorf("wrong str:%s", helperKey.str)
	}
}
func Test_QueryCondHelperLE(t *testing.T) {
	helperKey := LE("id", ":v_id1")
	if helperKey.str != "id <= :v_id1" {
		t.Errorf("wrong str:%s", helperKey.str)
	}
}
