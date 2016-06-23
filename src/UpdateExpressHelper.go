package GoDynamoDB

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"strings"
)

type UpdateCondExpressHelper struct {
	str        string
	expressMap map[string]*string
}

//won't gusee place holder for it, just create a query with place holer, I think that's more effective for a Server side program
func NewUpdateExpress() *UpdateCondExpressHelper {
	return &UpdateCondExpressHelper{
		str: "",
	}
}

func (q *UpdateCondExpressHelper) AddExpressMap(originalKey, placeHolder string) *UpdateCondExpressHelper {
	if nil == q.expressMap {
		q.expressMap = make(map[string]*string)
	}
	q.expressMap[placeHolder] = aws.String(originalKey)
	return q
}

func RemoveField(path ...string) *UpdateCondExpressHelper {
	str := fmt.Sprintf("REMOVE %s", strings.Join(path, ","))
	return &UpdateCondExpressHelper{
		str: str,
	}
}

type SetValueMap struct {
	SetName, SetValue string
}

func AddToField(toAdd ...SetValueMap) *UpdateCondExpressHelper {
	nameMapList := make([]string, len(toAdd))
	for i := 0; i < len(toAdd); i++ {
		nameMapList[i] = fmt.Sprintf("%s %s", toAdd[i].SetName, toAdd[i].SetValue)
	}
	str := fmt.Sprint("Add %s", strings.Join(nameMapList, ","))

	return &UpdateCondExpressHelper{
		str: str,
	}
}

func DeleteFromSet(toDelete ...SetValueMap) *UpdateCondExpressHelper {
	nameMapList := make([]string, len(toDelete))
	for i := 0; i < len(toDelete); i++ {
		nameMapList[i] = fmt.Sprintf("%s %s", toDelete[i].SetName, toDelete[i].SetValue)
	}
	str := fmt.Sprint("DELETE %s", strings.Join(nameMapList, ","))

	return &UpdateCondExpressHelper{
		str: str,
	}
}
