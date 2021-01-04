package pgrecorder

import "database/sql"

// Executable implementations expose Exec with the same args as sql.DB.Exec
type PGRecorder interface {
	Exec(string, ...interface{}) (sql.Result, error)
}
