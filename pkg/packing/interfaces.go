package packing

type Packable interface {
	Pack(Recorder) error
}

type Recorder interface {
	Record(interface{}) error
}

type PackableFactory func() Packable
