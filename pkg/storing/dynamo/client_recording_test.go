package dynamo

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Record is "wide" as far as what it accepts as input. This is intentional and idiomatic.
// We will nest 2 structs together to simulate what the method may recieve in use.
type testFoo struct {
	Spam string `json:"spam"`
}

type testBar struct {
	Foo  *testFoo
	Eggs int
}

type recordingTestSuite struct {
	suite.Suite
	client *RecordClient // Client is the repository facade
}

func (s *recordingTestSuite) SetupSuite() {
	// NOTE if multiple tests existed in this suite we'd likely use SetupTest
	s.client = NewRecordClient(&dynamicRecorderMock{})
}

// TestRecord should only test that Record eventually passes the Dynamo Facade's PutItem method the correct input
func (s *recordingTestSuite) TestRecord() {
	// "spy" on the mocked db facade
	// NOTE the .Return on the mock.Call allows us to create tests for specific return cases (like error handling)
	// though we are not using the return here as we only care about the input (or an err)
	// NOTE we still have to return the proper type...
	s.client.db.(*dynamicRecorderMock).On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, nil)

	// record marshals the input into a map, we'll create our own to compare...
	d := &testBar{Eggs: 12, Foo: &testFoo{Spam: "spam"}}

	// call the actual client Record method, it will, of course, hit our mock vs actual dynamo
	assert.Nil(s.T(), s.client.Record(d))

	// we can reference what was actually passed thru to the mocked db from Record via:
	args := s.client.db.(*dynamicRecorderMock).Calls[0].Arguments

	// hand create what that arg should have been...
	m, _ := dynamodbattribute.MarshalMap(d)
	i := &dynamodb.PutItemInput{
		TableName: aws.String(EVENTSTORE),
		Item:      m,
	}

	args.Assert(s.T(), i)
}

func TestSuite(t *testing.T) {
	suite.Run(t, &recordingTestSuite{})
}
