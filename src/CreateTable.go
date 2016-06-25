package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

import (
	"fmt"
	"reflect"
	"strings"
)

type CreateTableExecutor struct {
	input *dynamodb.CreateTableInput
	db    *dynamodb.DynamoDB
}

func (d GoDynamoDB) GetCreateTableExecutor(i CreateCollectionModel) (*CreateTableExecutor, error) {
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
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
			throughPut, exist := i.GetPrevision()[key]
			if !exist {
				return nil, NewDynError("no through put for table is provided")
			}
			input.ProvisionedThroughput = &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(throughPut.read),
				WriteCapacityUnits: aws.Int64(throughPut.write),
			}

		} else {
			//this is an index key
			if "" == pKeyName && "" != rKeyName {
				//this is a local index
				localIndex := &dynamodb.LocalSecondaryIndex{
					IndexName: aws.String(key),
				}
				tablePartKey, _ := cache.key[tableName]
				localIndex.KeySchema = make([]*dynamodb.KeySchemaElement, 1)

				keyAlis, aliasExist := cache.fieldNameMap[tablePartKey.pkey]
				if !aliasExist {
					keyAlis = tablePartKey.pkey
				}

				hashKeySchema := &dynamodb.KeySchemaElement{
					AttributeName: aws.String(keyAlis),
					KeyType:       aws.String(dynamodb.KeyTypeHash),
				}
				localIndex.KeySchema[0] = hashKeySchema

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

				iprojection, ok := i.GetProjection()[key]
				if !ok {
					return nil, NewDynError("no projection defination for index")
				}
				projection := &dynamodb.Projection{
					ProjectionType: aws.String(iprojection.projectType),
				}
				if ProjectDefined == iprojection.projectType {
					fieldList := strings.Split(iprojection.projectFields, ",")
					projection.NonKeyAttributes = make([]*string, len(fieldList))
					for i, value := range fieldList {
						projection.NonKeyAttributes[i] = aws.String(value)
					}
				}
				localIndex.Projection = projection
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
				throughPut, exist := i.GetPrevision()[key]
				if !exist {
					return nil, NewDynError("no through put for local index is provided")
				}
				globalIndex.ProvisionedThroughput = &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(throughPut.read),
					WriteCapacityUnits: aws.Int64(throughPut.write),
				}
				if nil == input.GlobalSecondaryIndexes {
					input.GlobalSecondaryIndexes = make([]*dynamodb.GlobalSecondaryIndex, 0)
				}
				iprojection, ok := i.GetProjection()[key]
				if !ok {
					return nil, NewDynError("no projection defination for index")
				}
				projection := &dynamodb.Projection{
					ProjectionType: aws.String(iprojection.projectType),
				}
				if ProjectDefined == iprojection.projectType {
					fieldList := strings.Split(iprojection.projectFields, ",")
					projection.NonKeyAttributes = make([]*string, len(fieldList))
					for i, value := range fieldList {
						projection.NonKeyAttributes[i] = aws.String(value)
					}
				}
				globalIndex.Projection = projection
				input.GlobalSecondaryIndexes = append(input.GlobalSecondaryIndexes, globalIndex)
			}
		}
	}
	return ret, nil
}

func (e *CreateTableExecutor) insertAtt(t reflect.Type, index int) error {
	fieldDef := t.Field(index)
	name := fieldDef.Name
	aliaTag := fieldDef.Tag.Get("DAlias")
	if "" != aliaTag {
		name = aliaTag
	}

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

func (db GoDynamoDB) DeleteTable(i ReadModel) error {
	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(i.GetTableName()),
	}
	_, err := db.db.DeleteTable(input)
	return err
}
