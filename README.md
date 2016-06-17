# GoDynamoDB
DynamoORM for go, just for fun because I don't know go now.

## Reference
Some of the encoder and decoder codes is copied from github.com/justonia/dynamodb-helpers, greate thanks for justonia, that repo saves my time.

## known limit:
- if you wanna decode, your field must be exportable(starting with capical letter)
- if there is a field named "Test", and you set an alias as "test" to it, then don't use "test" as field name in your struct, otherwise it will get error when don't decode.
- DPKey and DRKey should use the table/index 's name
- table name can't be same as index name



