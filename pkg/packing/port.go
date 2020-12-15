package packing

type Port struct {
	Maker  string `json:"maker"`
	First  string `json:"first"`
	Second string `json:"second"`
	Result string `json:"result"`
	Algo   string `json:"algo"` // "rando" || "carrollian"
}

func (p *Port) Pack(r Recorder) error {
	// until carrollian algo is implemented, all ports are rando generated
	// TODO implement carrollian algo
	if err := randoPack(p); err != nil {
		// TODO logging
		return err
	}

	// with the port made, record this event
	if err := r.Record(NewEvent(p)); err != nil {
		return err
	}
	return nil
}

func NewPort() *Port {
	return &Port{
		Algo: RANDO,
	}
}
