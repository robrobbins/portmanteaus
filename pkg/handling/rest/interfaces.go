package rest

import "github.com/robrobbins/portmanteaus/pkg/packing"

type PackingRecorder interface {
	Record(string, *packing.Port) error
}
