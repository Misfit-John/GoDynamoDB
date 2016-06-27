package GoDynamoDB

import (
	"testing"
)

func Test_PutItem(t *testing.T) {
	e, err := GetDBInstance().GetPutItemExcutor(&TestStruct{Name: "John", Id: "123"})
	if err != nil {
		t.Error(err.Error())
	}
	execErr := e.Exec()
	if execErr != nil {
		t.Error(execErr.Error())
	}
}

func Test_BatchPutItem(t *testing.T) {
	toPut := make([]WriteModel, 3)
	toPut[0] = &TestStruct{Id: "12355", Name: "John123"}
	toPut[1] = &TestStruct2nd{Id: "1234", Name: "2ndJohn", School: "UESTC"}
	toPut[2] = &TestStruct{Id: "1235", Name: "1235John"}
	e, err := GetDBInstance().GetBatchWriteItemExecutor(toPut, make([]WriteModel, 0))
	if err != nil {
		t.Error(err.Error())
	}

	execErr := e.Exec()
	if nil != execErr {
		t.Error(execErr.Error())
	}

}

func init() {
	InitLocalDBInstance("http://127.0.0.1:8000")
	exe1, _ := GetDBInstance().GetCreateTableExecutor(&TestStruct{})
	exe1.Exec()
	exe2, _ := GetDBInstance().GetCreateTableExecutor(&TestStruct2nd{})
	exe2.Exec()
	exe3, _ := GetDBInstance().GetCreateTableExecutor(&QueryTest{})
	exe3.Exec()
}
