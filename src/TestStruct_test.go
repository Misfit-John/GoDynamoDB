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

type QueryTest struct {
	Id    string `DAlias:"id" DPKey:"query_test"`
	Index int    `DAlias:"index" DRKey:"query_test"`
}

func (*QueryTest) GetTableName() string {
	return "query_test"
}

type CreateTest struct {
	Id     string `DAlias:"id" DPKey:"create_test"`
	Rang   int    `DAlias:"trange" DRKey:"create_test"`
	LRange int    `DAlias:"LRange" DRKey:"lindex"`
	Pid    string `DPKey:"gindex"`
	Prange string `DRKey:"gindex"`
}

func (*CreateTest) GetTableName() string {
	return "create_test"
}

func (*CreateTest) GetPrevision() map[string]Throughput {
	return map[string]Throughput{
		"create_test": NewThroughput(1, 1),
		"lindex":      NewThroughput(1, 1),
		"gindex":      NewThroughput(1, 1),
	}
}

func (*CreateTest) GetProjection() map[string]ProjectionDefination {
	return map[string]ProjectionDefination{
		"lindex": NewProjectionDefination(ProjectKey, ""),
		"gindex": NewProjectionDefination(ProjectKey, ""),
	}
}
