package GoDynamoDB

import (
	"fmt"
	"reflect"
	"testing"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func FooReflect(f interface{}) {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Type.Name() == "string" {
			valueField.SetString("hello")
		}

		if typeField.Type.Name() == "int" {
			valueField.SetInt(10)
		}
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n",
			typeField.Name,
			valueField.Interface(),
			tag.Get("tag_name"))
	}
}

func TestRef(t *testing.T) {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	FooReflect(f)
}
