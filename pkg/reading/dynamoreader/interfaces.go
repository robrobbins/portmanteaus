package dynamoreader

import "github.com/aws/aws-sdk-go/service/dynamodb"

// DynamicReader interfaces with dynamodb, encapsulating all necessary querying functionality
type DynamicReader interface {
	Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}
