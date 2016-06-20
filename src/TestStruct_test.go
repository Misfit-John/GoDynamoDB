package GoDynamoDB

type TestStruct struct {
	Name string
	Id   string `DAlias:"id" DPKey:"Test"`
}

func (*TestStruct) GetTableName() string {
	return "Test"
}

type TestStruct2nd struct {
	Name   string
	School string
	Id     string `DAlias:"id" DPKey:"test2nd"`
}

func (*TestStruct2nd) GetTableName() string {
	return "test2nd"
}
