package GoDynamoDB

type WriteModel interface {
	GetTableName() string
}

type ReadModel interface {
	GetTableName() string
}

const (
	WriteProvisionedThroughput = 0
	ReadProvisionedThroughput  = 1
	ProjectAll                 = "ALL"
	ProjectKey                 = "KEYS_ONLY"
	ProjectDefined             = "INCLUDE"
)

type ProjectionDefination struct {
	projectType, projectFields string
}

type Throughput struct {
	read, write int64
}

func NewProjectionDefination(projectType, fields string) ProjectionDefination {
	if projectType == ProjectDefined && fields == "" {
		panic("no fields is defined for ProjectDefined")
	}
	return ProjectionDefination{projectType: projectType, projectFields: fields}
}

func NewThroughput(read, write int64) Throughput {
	if read <= 0 {
		read = 1
	}
	if write <= 0 {
		write = 1
	}
	return Throughput{read: read, write: write}
}

type CreateCollectionModel interface {
	GetTableName() string
	//provide only table throughput and global index throughput
	GetPrevision() map[string]Throughput
	GetProjection() map[string]ProjectionDefination
}
