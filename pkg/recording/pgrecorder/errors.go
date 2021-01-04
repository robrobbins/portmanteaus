package pgrecorder

import "errors"

var errEventNotInserted = errors.New("event not written to events table")
var errTypeAssertionToByteArray = errors.New("type assertion to []byte failed")
