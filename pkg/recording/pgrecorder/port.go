package pgrecorder

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/robrobbins/portmanteaus/pkg/packing"
)

// Port is the repository representation of a portmanteau.
// It must implement the sql.driver.valuer in order to be marshal-able
type Port struct {
	Maker  string `json:"maker"`
	First  string `json:"first"`
	Second string `json:"second"`
	Result string `json:"result"`
	Algo   string `json:"algo"` // "rando" || "carrollian"
}

// Value allows the Port to implement driver.Value, thus able to be marshaled by the sql driver
func (p *Port) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// NOTE: this is for posterity and will be moved to a pgreader.
// Scan allows for the unmarshaling of database serializations into structs
// func(p *Port) Scan(v interface{}) error {
// 	b, ok := v.([]byte)
//
// 	if !ok {
// 		return errTypeAssertionToByteArray
// 	}
//
// 	return json.Unmarshal(b, p)
// }

func NewPort(p *packing.Port) *Port {
	return &Port{
		Maker:  p.Maker,
		First:  p.First,
		Second: p.Second,
		Result: p.Result,
		Algo:   p.Algo,
	}
}
