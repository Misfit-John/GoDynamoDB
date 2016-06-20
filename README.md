# GoDynamoDB
DynamoORM for go, just for fun because I don't know go now.

## supported tag
**DAlias** the alia name of the field which will finally show in db.

**DPKey** mark this field as the partition key of table/index, the value should be the name of the table/index.

**DRKey** mark this field as the range key of table/index, the value should be the name of the table/index.

## Reference
Some of the encoder and decoder codes is copied from github.com/justonia/dynamodb-helpers, great thanks for justonia, that repo saves my time.

## known limit:
- if you wanna decode this field, your field must be exportable, which means the field's name should start with capital letter. This limitation is from golang's export strategy , I have no idea with it.
- if there is a field named "Test", and you set an alias as "test" to it, then you shouldn't use "test" as field name in your struct, otherwise it will get error when doing encode/decode.
- DPKey and DRKey should use the table/index's name
- table name can't be same as index name
- Will sleep 100ms or more if the request for BatchGetItem is too large to return in one time.
- Don't allow duplicate keys in batch get request because the limit of AWS SDK


