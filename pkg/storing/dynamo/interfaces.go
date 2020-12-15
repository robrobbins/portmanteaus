package dynamo

import "github.com/aws/aws-sdk-go/service/dynamodb"

// DynamicReader interfaces with dynamodb, encapsulating all necessary querying functionality
type DynamicReader interface {
	Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

// DynamicRecorder interfaces with dynamodb, encapsulating all necessary writing functionality
type DynamicRecorder interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
