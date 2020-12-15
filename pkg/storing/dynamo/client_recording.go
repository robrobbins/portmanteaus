package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Record persists the input to the specified eventstore.
// NOTE that the return is discarded.
func (r *RecordClient) Record(e interface{}) error {
	m, err := dynamodbattribute.MarshalMap(e)

	if err != nil {
		return err
	}

	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(r.eventStore),
		Item:      m,
	})

	if err != nil {
		return err
	}

	return nil
}
