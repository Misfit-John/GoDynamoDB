package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import (
	"math"
	"reflect"
	"strconv"
)

func encodeToQueryAtt(i interface{}) (*dynamodb.AttributeValue, error) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	// query atti should not return in these types
	case reflect.Struct, reflect.Ptr, reflect.Slice, reflect.Map, reflect.Array:
		return nil, NewDynError("unknow type")
	}

	return encodeToAtt(v)
}

func encodeToAtt(v reflect.Value) (*dynamodb.AttributeValue, error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			b := true
			return &dynamodb.AttributeValue{NULL: aws.Bool(b)}, nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		b := v.Bool()
		return &dynamodb.AttributeValue{BOOL: aws.Bool(b)}, nil
	case reflect.String:
		s := v.String()
		if len(s) == 0 {
			b := true
			return &dynamodb.AttributeValue{NULL: aws.Bool(b)}, nil
		} else {
			return &dynamodb.AttributeValue{S: aws.String(s)}, nil
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := strconv.FormatInt(v.Int(), 10)
		return &dynamodb.AttributeValue{N: aws.String(n)}, nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := strconv.FormatUint(v.Uint(), 10)
		return &dynamodb.AttributeValue{N: aws.String(n)}, nil

	case reflect.Float32, reflect.Float64:
		f := v.Float()
		if math.IsInf(f, 0) || math.IsNaN(f) {
			return nil, NewDynError("aws.dynamodb.convertToNumericString: NaN and infinite floats not supported")
		}
		fs := strconv.FormatFloat(f, 'g', -1, v.Type().Bits())
		return &dynamodb.AttributeValue{N: aws.String(fs)}, nil

	case reflect.Struct:
		if v.IsNil() {
			b := true
			return &dynamodb.AttributeValue{NULL: aws.Bool(b)}, nil
		} else {
			m, err := encodeStruct(v)
			if err == nil {
				return &dynamodb.AttributeValue{M: m}, nil
			} else {
				return nil, err
			}
		}
	case reflect.Map:
		if v.IsNil() {
			b := true
			return &dynamodb.AttributeValue{NULL: aws.Bool(b)}, nil
		}

		if v.Type().Key().Kind() != reflect.String {
			return nil, NewDynError("can't transform from a map who is not using string as key")
		}

		containerOut := map[string]*dynamodb.AttributeValue{}
		for _, key := range v.MapKeys() {
			v2, err := encodeToAtt(v.MapIndex(key))
			if err != nil {
				return nil, err
			}
			if v2 != nil {
				containerOut[key.String()] = v2
			} else {
				b := true
				containerOut[key.String()] = &dynamodb.AttributeValue{NULL: &b}
			}
		}
		return &dynamodb.AttributeValue{M: containerOut}, nil
	case reflect.Slice:
		// empty lists are not supported in dynamo, kinda sucks we can't
		// differentiate nil slices from empty slices...
		if v.IsNil() || v.Len() == 0 {
			b := true
			return &dynamodb.AttributeValue{NULL: aws.Bool(b)}, nil
		}

		// Special-case, byte blob, binary can't be nil...
		if v.Type().Elem().Kind() == reflect.Uint8 {
			if v.Len() == 0 {
				b := true
				return &dynamodb.AttributeValue{NULL: aws.Bool(b)}, nil
			} else {
				return &dynamodb.AttributeValue{B: v.Bytes()}, nil
			}
		}

		fallthrough

	case reflect.Array:
		arrayLength := v.Len()
		containerOut := make([]*dynamodb.AttributeValue, arrayLength)
		for i := 0; i < arrayLength; i++ {
			v2, err := encodeToAtt(v.Index(i))
			if err != nil {
				return nil, err
			}
			if v2 != nil {
				containerOut[i] = v2
			} else {
				b := true
				containerOut[i] = &dynamodb.AttributeValue{NULL: aws.Bool(b)}
			}
		}
		return &dynamodb.AttributeValue{L: containerOut}, nil
	default:
		return nil, NewDynError("unknow datatype, please contect author")
	}
	return nil, NewDynError("unknow error")
}

func encodeStruct(v reflect.Value) (map[string]*dynamodb.AttributeValue, error) {
	out := map[string]*dynamodb.AttributeValue{}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	cache, err := getCache(v.Type())
	if err != nil {
		return nil, err
	}
	for k, index := range cache.fieldIndex {
		att, err := encodeToAtt(v.Field(index))
		if err == nil {
			if mapName, ok := cache.fieldNameMap[k]; ok {
				//use DAlias name if this tag exist
				out[mapName] = att
			} else {
				//use original name if no tag exist
				out[k] = att
			}
		} else {
			return nil, err
		}
	}
	return out, nil
}

func encodeKeyOnly(i interface{}, name string) (map[string]*dynamodb.AttributeValue, error) {
	v := reflect.ValueOf(i)
	//copy from encodeStruct
	out := map[string]*dynamodb.AttributeValue{}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	cache, err := getCache(v.Type())
	if err != nil {
		return nil, err
	}

	keyPair, ok := cache.key[name]
	if !ok {
		return nil, NewDynError("no key pair is defined for table/index name")
	}

	if index, ok := cache.fieldIndex[keyPair.pkey]; ok {
		att, err := encodeToAtt(v.Field(index))
		if err == nil {
			if mapName, ok := cache.fieldNameMap[keyPair.pkey]; ok {
				//use DAlias name if this tag exist
				out[mapName] = att
			} else {
				//use original name if no tag exist
				out[keyPair.pkey] = att
			}
		} else {
			return nil, err
		}
	}

	if index, ok := cache.fieldIndex[keyPair.rkey]; ok {
		att, err := encodeToAtt(v.Field(index))
		if err == nil {
			if mapName, ok := cache.fieldNameMap[keyPair.rkey]; ok {
				//use DAlias name if this tag exist
				out[mapName] = att
			} else {
				//use original name if no tag exist
				out[keyPair.rkey] = att
			}
		} else {
			return nil, err
		}
	}
	return out, nil
}

func encode(i interface{}) (map[string]*dynamodb.AttributeValue, error) {
	v := reflect.ValueOf(i)
	return encodeStruct(v)
}
