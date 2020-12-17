package dynamoreader

type Reader struct {
	db         DynamicReader // the repository facade in front of dynamo
	eventStore string        // dynamo table name
}

func NewReader(db DynamicReader) *Reader {
	return &Reader{
		db:         db,
		eventStore: EVENTSTORE,
	}
}
