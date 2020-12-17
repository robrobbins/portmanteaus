package dynamorecorder

import "github.com/aws/aws-sdk-go/service/dynamodb"

// DynamicRecorder interfaces with dynamodb, encapsulating all necessary writing functionality
type DynamicRecorder interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
