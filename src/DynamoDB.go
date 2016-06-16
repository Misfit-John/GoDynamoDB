package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/aws"
import "fmt"

type GoDynamoDB struct {
	db *dynamodb.DynamoDB
}

var dbInstance GoDynamoDB

//we will use a singleton so you should only call this for one time
func InitLocalDBInstance(endpoint string) error {
	if nil == dbInstance.db {
		dbInstance.db = dynamodb.New(session.New(
			&aws.Config{
				Region:   aws.String("us-east-1"),
				Endpoint: aws.String(endpoint)}))
	} else {
		return NewDynError("should not init the db again")
	}
	return nil
}

func InitServerDBInstance(region, accessKey, secretKey string) error {
	return NewDynError(fmt.Sprintf("not implement yet", region, accessKey, secretKey))
}

func GetDBInstance() GoDynamoDB {
	return dbInstance
}
