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
	toPut[2] = &TestStruct{Id: "12325", Name: "1235John"}
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
	//prepire table and data here, you may get error if you do it many times , but don't worry, I won't throw error on this
	exe1, _ := GetDBInstance().GetCreateTableExecutor(&TestStruct{})
	exe1.Exec()
	insertT1, _ := GetDBInstance().GetPutItemExcutor(&TestStruct{Id: "123", Name: "John"})
	insertT2, _ := GetDBInstance().GetPutItemExcutor(&TestStruct{Id: "1235", Name: "John"})
	insertT2.Exec()
	insertT1.Exec()

	exe2, _ := GetDBInstance().GetCreateTableExecutor(&TestStruct2nd{})
	exe2.Exec()
	insertT2nd1, _ := GetDBInstance().GetPutItemExcutor(&TestStruct2nd{Id: "1235", Name: "2ndJohn", School: "UESTC"})
	insertT2nd1.Exec()

	exe3, _ := GetDBInstance().GetCreateTableExecutor(&QueryTest{})
	exe3.Exec()
	insertQ1, _ := GetDBInstance().GetPutItemExcutor(&QueryTest{Id: "123", Index: 11})
	insertQ2, _ := GetDBInstance().GetPutItemExcutor(&QueryTest{Id: "123", Index: 13})
	insertQ3, _ := GetDBInstance().GetPutItemExcutor(&QueryTest{Id: "123", Index: 15})
	insertQ1.Exec()
	insertQ2.Exec()
	insertQ3.Exec()
}
