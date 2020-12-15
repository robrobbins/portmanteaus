package dynamo

type ReadClient struct {
	db         DynamicReader // the repository facade in front of dynamo
	eventStore string        // dynamo table name
}

func NewReadClient(db DynamicReader) *ReadClient {
	return &ReadClient{
		db:         db,
		eventStore: EVENTSTORE,
	}
}

type RecordClient struct {
	db         DynamicRecorder
	eventStore string
}

func NewRecordClient(db DynamicRecorder) *RecordClient {
	return &RecordClient{
		db:         db,
		eventStore: EVENTSTORE,
	}
}
