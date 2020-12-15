package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

type dynamicReaderMock struct {
	mock.Mock
}

type dynamicRecorderMock struct {
	mock.Mock
}

// PutItem allows dynamicRecorderMock to implement DynamicRecorder
func (r *dynamicRecorderMock) PutItem(i *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := r.Called(i)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}
