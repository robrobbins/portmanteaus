package packing

type Packable interface {
	Pack(PortRecorder) error
}

type PortRecorder interface {
	Record(string, *Port) error
}

type PackableFactory func() Packable
