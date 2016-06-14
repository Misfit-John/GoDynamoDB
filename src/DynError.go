package GoDynamoDB

const (
	Unknow = "unknow error"
)

type DynErrors struct {
	errStr string
}

func (e *DynErrors) Error() string {
	return e.errStr
}

func NewDynError(errString string) *DynErrors {
	return &DynErrors{errStr: errString}
}
