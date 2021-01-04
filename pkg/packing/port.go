package packing

type Port struct {
	Maker  string
	First  string
	Second string
	Result string
	Algo   string // "rando" || "carrollian"
}

func (p *Port) Pack(r PortRecorder) error {
	// until carrollian algo is implemented, all ports are rando generated
	// TODO implement carrollian algo
	if err := randoPack(p); err != nil {
		// TODO logging
		return err
	}

	// the center passes domain_event, and the data being serialized
	if err := r.Record(PORT_CREATED, p); err != nil {
		return err
	}
	return nil
}

func NewPort() *Port {
	return &Port{
		Algo: RANDO,
	}
}
