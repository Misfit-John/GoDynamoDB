package GoDynamoDB

import "reflect"

type KeyPair struct {
	pkey, rkey string //p for partition, r for range
}

type FieldCache struct {
	name          string
	fieldIndex    map[string]int     //map from name to index
	expressionMap map[string]string  //expressiono map
	key           map[string]KeyPair //key map, key is table/index name, value will be (partition key, range key)
	nilAct        map[string]string  //should be one of ignore, setNull, panic
	fieldNameMap  map[string]string  //key is original field name in struct, value is name in dynamodb.
}

var fieldCache map[string]*FieldCache

func GetCache(t reflect.Type) (*FieldCache, error) {
	if nil == fieldCache {
		fieldCache = make(map[string]*FieldCache)
	}
	field, ok := fieldCache[t.Name()]
	if ok {
		return field, nil
	} else {
		newField, err := initCache(t)
		if err != nil {
			return nil, err
		}
		return newField, nil
	}

}

func initCache(t reflect.Type) (*FieldCache, error) {
	out := FieldCache{
		fieldIndex: make(map[string]int),
	}
	tNmae := t.Name()
	for i := 0; i < t.NumField(); i++ {
		fieldI := t.Field(i)
		out.fieldIndex[fieldI.Name] = i
	}

	if nil == fieldCache {
		fieldCache = make(map[string]*FieldCache)
	}
	fieldCache[tNmae] = &out

	return &out, nil
}
