package GoDynamoDB

import (
	"fmt"
	"reflect"
)

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
		fieldIndex:   make(map[string]int),
		fieldNameMap: make(map[string]string),
		key:          make(map[string]KeyPair),
	}
	tNmae := t.Name()

	for i := 0; i < t.NumField(); i++ {
		fieldI := t.Field(i)
		out.fieldIndex[fieldI.Name] = i
		tags := fieldI.Tag
		nameTag := tags.Get("DAlias")
		if "" != nameTag {
			out.fieldNameMap[fieldI.Name] = nameTag
			//should also add the alias' into the index map
			out.fieldIndex[nameTag] = i
		}
		keyTag := tags.Get("DPKey")
		if "" != keyTag {
			if pair, ok := out.key[keyTag]; ok {
				if pair.pkey != "" {
					return nil, NewDynError(fmt.Sprintf("double defined partition key:", fieldI.Name))
				} else {
					pair.pkey = fieldI.Name
				}
			} else {
				out.key[keyTag] = KeyPair{pkey: fieldI.Name}
			}
		}

		rangeTag := tags.Get("DRKey")
		if "" != keyTag {
			if pair, ok := out.key[rangeTag]; ok {
				if pair.rkey != "" {
					return nil, NewDynError(fmt.Sprintf("double defined range key:", fieldI.Name))
				} else {
					pair.pkey = fieldI.Name
				}
			} else {
				out.key[rangeTag] = KeyPair{rkey: fieldI.Name}
			}
		}

	}

	if nil == fieldCache {
		fieldCache = make(map[string]*FieldCache)
	}
	fieldCache[tNmae] = &out

	return &out, nil
}
