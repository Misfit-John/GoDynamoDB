package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
)

var imapType = reflect.TypeOf(map[string]interface{}{})
var ilistType = reflect.TypeOf([]interface{}{})

func decode(m map[string]*dynamodb.AttributeValue, ret interface{}) error {
	v := reflect.ValueOf(ret)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return decodeStruct(m, v)
}

func decodeAttribute(attr *dynamodb.AttributeValue, v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
			return nil
		} else {
			v.Set(reflect.New(v.Type().Elem()))
			v = v.Elem()
		}
	}

	switch v.Kind() {
	case reflect.Bool:
		if attr.BOOL != nil {
			v.Set(reflect.ValueOf(*attr.BOOL))
		} else if attr.NULL != nil {
			v.Set(reflect.ValueOf(false))
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if attr.N != nil {
			n, err := strconv.ParseInt(*attr.N, 10, 64)
			if err != nil || v.OverflowInt(n) {
				return NewDynError(fmt.Sprintf("overflow number %s for type %s", *attr.N, v.Type().String()))
			}
			v.SetInt(n)
		} else if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if attr.N != nil {
			n, err := strconv.ParseUint(*attr.N, 10, 64)
			if err != nil || v.OverflowUint(n) {
				return NewDynError(fmt.Sprintf("overflow number %s for type %s", *attr.N, v.Type().String()))
			}
			v.SetUint(n)
		} else if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
		}

	case reflect.Float32, reflect.Float64:
		if attr.N != nil {
			n, err := strconv.ParseFloat(*attr.N, v.Type().Bits())
			if err != nil || v.OverflowFloat(n) {
				return NewDynError(fmt.Sprintf("overflow number %s for type %s", *attr.N, v.Type().String()))
			}
			v.SetFloat(n)
		} else if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
		}

	case reflect.String:
		if attr.S != nil {
			v.SetString(*attr.S)
		} else if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
		}

	case reflect.Struct:
		if attr.M != nil {
			if err := decodeStruct(attr.M, v); err != nil {
				return err
			}
		} else if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
		}

	case reflect.Map:
		if attr.M != nil {
			// map must have string kind
			t := v.Type()
			if t.Key().Kind() != reflect.String {
				return NewDynError(fmt.Sprintf("cannot decode a map with a non-string key: %s", t.Key().String()))
			}
			if v.IsNil() {
				v.Set(reflect.MakeMap(t))
			}
			if err := decodeMap(attr.M, v); err != nil {
				return err
			}
		} else if attr.NULL != nil {
			v.Set(reflect.Zero(v.Type()))
		}

	case reflect.Slice:
		// []byte handling
		if v.Type().Elem().Kind() == reflect.Uint8 {
			switch {
			case attr.B != nil:
				v.Set(reflect.ValueOf(attr.B[0:len(attr.B)]))

			case attr.S != nil:
				d, err := base64.StdEncoding.DecodeString(*attr.S)
				if err != nil {
					return NewDynError(fmt.Sprintf("cannot base64 decode string: %s", err.Error()))
				}
				v.Set(reflect.ValueOf(d))

			case attr.NULL != nil:
				v.Set(reflect.Zero(v.Type()))

			default:
				// nothing to do, silently ignore failed coercion
			}
			return nil
		}

		fallthrough

	case reflect.Array:
		return decodeArray(attr, v)

	case reflect.Interface:
		if v.NumMethod() != 0 {
			// TODO: Might be worth adding a custom demarshalling hook
			return NewDynError(fmt.Sprintf("cannot decode into non-empty interface type: %s", v.Type().String()))
		}

		switch {
		case attr.B != nil:
			v.Set(reflect.ValueOf(attr.B[0:len(attr.B)]))
			break

		case attr.BOOL != nil:
			v.Set(reflect.ValueOf(*attr.BOOL))

		case attr.S != nil:
			v.Set(reflect.ValueOf(*attr.S))

		case attr.N != nil:
			var n float64
			var err error
			n, err = strconv.ParseFloat(*attr.N, 64)
			if err != nil {
				return NewDynError(fmt.Sprintf("error parsing number %s into type float64", *attr.N))
			}
			v.Set(reflect.ValueOf(n))

		case attr.NULL != nil:
			v.Set(reflect.Zero(v.Type()))
			break

		case attr.M != nil:
			m := reflect.MakeMap(imapType)
			if err := decodeMap(attr.M, m); err != nil {
				return err
			}
			v.Set(m)

		case attr.L != nil:
			l := reflect.New(ilistType)
			if err := decodeArray(attr, l.Elem()); err != nil {
				return err
			}
			v.Set(l.Elem())

		default:
			return NewDynError("unknown error decoding interface value")
		}

	case reflect.Ptr:
		// should have had indirection taken care of above
		return NewDynError(fmt.Sprintf("could not handle multiple layers of pointers"))
	default:
		return NewDynError(fmt.Sprintf("aws.dynamodb.EncodeError: unsupported type for field: %#v", v.Type()))
	}
	return nil
}

// private
func decodeStruct(attrs map[string]*dynamodb.AttributeValue, v reflect.Value) error {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fields, err := GetCache(v.Type())
	if err != nil {
		return err
	}
	for k, attr := range attrs {
		if index, ok := fields.fieldIndex[k]; ok {
			fv := v.Field(index)
			if err := decodeAttribute(attr, fv); err != nil {
				return err
			}
		}
	}
	return nil
}

func decodeMap(attr map[string]*dynamodb.AttributeValue, v reflect.Value) error {
	// map must have string kind
	if !v.IsValid() {
		v.Set(reflect.MakeMap(v.Type()))
	}
	elemType := v.Type().Elem()
	for key, subAttr := range attr {
		value := reflect.New(elemType).Elem()
		if err := decodeAttribute(subAttr, value); err != nil {
			return err
		}
		kv := reflect.ValueOf(key).Convert(v.Type().Key())
		v.SetMapIndex(kv, value)
	}
	return nil
}

func decodeArray(attr *dynamodb.AttributeValue, v reflect.Value) error {
	t := v.Type()

	if attr.NULL != nil || attr.L == nil {
		v.Set(reflect.Zero(t))
		return nil
	}

	if t.Kind() == reflect.Slice {
		v.Set(reflect.MakeSlice(v.Type(), len(attr.L), len(attr.L)))
	}

	vlen := v.Len()

	if vlen == 0 {
		return nil
	}

	switch t.Elem().Kind() {
	case reflect.Interface:
		if t.Elem().NumMethod() != 0 {
			// TODO: If custom decoding hooks can be provided, support this
			return NewDynError(fmt.Sprintf("cannot decode into array of non-empty interface types: %s", v.Type().String()))
		}

		i := 0
		alen := len(attr.L)
		for ; i < vlen && i < alen; i++ {
			if err := decodeAttribute(attr.L[i], v.Index(i)); err != nil {
				return err
			}
		}

		// zero out the rest
		for ; i < vlen; i++ {
			v.Index(i).Set(reflect.Zero(t.Elem()))
		}

		return nil

	default:
		i := 0
		alen := len(attr.L)
		for ; i < vlen && i < alen; i++ {
			av := v.Index(i)
			if err := decodeAttribute(attr.L[i], av); err != nil {
				return err
			}
			if !av.IsValid() || !av.Type().AssignableTo(t.Elem()) {
				return NewDynError(fmt.Sprintf("could not assign list value %s to array element type %s", v.Type().String(), t.Elem().String()))
			}
		}

		for ; i < vlen; i++ {
			v.Index(i).Set(reflect.Zero(t.Elem()))
		}
	}
	return nil
}
