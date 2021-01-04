package pgrecorder

import (
	"fmt"
	"time"

	"github.com/robrobbins/portmanteaus/pkg/packing"
)

type PackingRecorder struct {
	db PGRecorder
}

func (r *PackingRecorder) Record(e string, p *packing.Port) error {
	// prepare...
	t := time.Now().Unix()
	port := NewPort(p)
	// exec...
	a, err := r.db.Exec(fmt.Sprintf("INSERT INTO %s VALUES($1, $2, $3, $4)", EVENTSTORE), e, CURRENT_EVENT_VERSION, t, port)

	if err != nil {
		return err
	}

	n, err := a.RowsAffected()

	if err != nil {
		return err
	}

	if int(n) < 1 {
		return errEventNotInserted
	}

	return nil
}

func NewPackingRecorder(db PGRecorder) *PackingRecorder {
	return &PackingRecorder{
		db: db,
	}
}
