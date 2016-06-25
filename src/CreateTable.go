package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

import (
	"fmt"
	"reflect"
)

type CreateTableExecutor struct {
	input *dynamodb.CreateTableInput
	db    *dynamodb.DynamoDB
}

func (d GoDynamoDB) GetCreateTableExecutor(i ReadModel) (*CreateTableExecutor, error) {
	t := reflect.TypeOf(i)
	cache, cacheError := getCache(t)
	if cacheError != nil {
		return nil, cacheError
	}

	tableName := i.GetTableName()
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
	}

	ret := &CreateTableExecutor{
		db:    d.db,
		input: input,
	}

	for key, value := range cache.key {
		rKeyName := value.rkey
		pKeyName := value.pkey

		if val, ok := cache.fieldNameMap[rKeyName]; ok {
			rKeyName = val
		}
		if val, ok := cache.fieldNameMap[pKeyName]; ok {
			pKeyName = val
		}

		if key == tableName {
			//this is table key
			hashKeySchema := &dynamodb.KeySchemaElement{
				AttributeName: aws.String(pKeyName),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			}
			input.KeySchema = make([]*dynamodb.KeySchemaElement, 1)
			input.KeySchema[0] = hashKeySchema
			pkeyIndex, pok := cache.fieldIndex[value.pkey]
			if !pok {
				return nil, NewDynError("no partition key for table")
			}
			ret.insertAtt(t, pkeyIndex)

			if "" != value.rkey {
				rangeKeySchema := &dynamodb.KeySchemaElement{
					AttributeName: aws.String(rKeyName),
					KeyType:       aws.String(dynamodb.KeyTypeRange),
				}
				input.KeySchema = append(input.KeySchema, rangeKeySchema)
				rkeyIndex, rExist := cache.fieldIndex[value.rkey]
				if !rExist {
					return nil, NewDynError("no range key")
				}
				ret.insertAtt(t, rkeyIndex)
			}
		} else {
			//this is an index key
			if "" == pKeyName && "" != rKeyName {
				//this is a local index
				localIndex := &dynamodb.LocalSecondaryIndex{
					IndexName: aws.String(key),
				}
				tablePartKey, _ := cache.key[tableName]
				pkeyIndex, pExist := cache.fieldIndex[tablePartKey.pkey]
				if !pExist {
					return nil, NewDynError("no part key exist")
				}
				localIndex.KeySchema = make([]*dynamodb.KeySchemaElement, 1)

				hashKeySchema := &dynamodb.KeySchemaElement{
					AttributeName: aws.String(pKeyName),
					KeyType:       aws.String(dynamodb.KeyTypeHash),
				}
				localIndex.KeySchema[0] = hashKeySchema
				if !pExist {
					return nil, NewDynError("no partition key")
				}
				ret.insertAtt(t, pkeyIndex)

				rangeKeySchema := &dynamodb.KeySchemaElement{
					AttributeName: aws.String(rKeyName),
					KeyType:       aws.String(dynamodb.KeyTypeRange),
				}
				localIndex.KeySchema = append(localIndex.KeySchema, rangeKeySchema)
				rkeyIndex, rExist := cache.fieldIndex[value.rkey]
				if !rExist {
					return nil, NewDynError("no range key")
				}
				ret.insertAtt(t, rkeyIndex)

				if nil == input.LocalSecondaryIndexes {
					input.LocalSecondaryIndexes = make([]*dynamodb.LocalSecondaryIndex, 0)
				}
				input.LocalSecondaryIndexes = append(input.LocalSecondaryIndexes, localIndex)

			} else {
				//this is a global secondery index
				globalIndex := &dynamodb.GlobalSecondaryIndex{
					IndexName: aws.String(key),
				}
				partKeySchema := &dynamodb.KeySchemaElement{
					AttributeName: aws.String(pKeyName),
					KeyType:       aws.String(dynamodb.KeyTypeHash),
				}
				globalIndex.KeySchema = make([]*dynamodb.KeySchemaElement, 1)
				globalIndex.KeySchema[0] = partKeySchema

				pkeyIndex, pEexist := cache.fieldIndex[value.pkey]
				if !pEexist {
					return nil, NewDynError("no partition key")
				}
				ret.insertAtt(t, pkeyIndex)

				if rKeyName != "" {
					rangeKeySchema := &dynamodb.KeySchemaElement{
						AttributeName: aws.String(rKeyName),
						KeyType:       aws.String(dynamodb.KeyTypeRange),
					}
					globalIndex.KeySchema = append(globalIndex.KeySchema, rangeKeySchema)
					rkeyIndex, rExist := cache.fieldIndex[value.rkey]
					if !rExist {
						return nil, NewDynError("now range key")
					}
					ret.insertAtt(t, rkeyIndex)
				}
				if nil == input.GlobalSecondaryIndexes {
					input.GlobalSecondaryIndexes = make([]*dynamodb.GlobalSecondaryIndex, 0)
				}
				input.GlobalSecondaryIndexes = append(input.GlobalSecondaryIndexes, globalIndex)
			}
		}
	}
	return ret, nil
}

func (e *CreateTableExecutor) insertAtt(t reflect.Type, index int) error {
	fieldDef := t.Field(index)
	name := fieldDef.Name
	field := fieldDef.Type
	for field.Kind() == reflect.Ptr {
		field = field.Elem()
	}
	keyType := ""

	switch field.Kind() {
	case reflect.String:
		keyType = "S"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		keyType = "N"
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		if field.Elem().Kind() == reflect.Uint8 {
			keyType = "B"
		}
	default:
		return NewDynError(fmt.Sprintf("unknow type for field: %s!", name))

	}

	if nil == e.input.AttributeDefinitions {
		e.input.AttributeDefinitions = make([]*dynamodb.AttributeDefinition, 0)
	}
	e.input.AttributeDefinitions = append(e.input.AttributeDefinitions, &dynamodb.AttributeDefinition{
		AttributeName: aws.String(name),
		AttributeType: aws.String(keyType),
	})
	return nil
}

func (e *CreateTableExecutor) Exec() error {
	_, err := e.db.CreateTable(e.input)
	if err != nil {
		return err
	}
	return nil
}
