package dynamorecorder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Recorder struct {
	db         DynamicRecorder
	eventStore string
}

// Record persists the input to the specified eventstore.
// NOTE if the Recorder possesed a large number of methods, they could be put in separate files
// NOTE that the return is discarded.
func (r *Recorder) Record(e interface{}) error {
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

func NewRecorder(db DynamicRecorder) *Recorder {
	return &Recorder{
		db:         db,
		eventStore: EVENTSTORE,
	}
}
