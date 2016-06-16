package GoDynamoDB

type WriteModel interface {
	GetTableName() string
}

type ReadModel interface {
	GetTableName() string
	IsConsistentRead() bool
}
