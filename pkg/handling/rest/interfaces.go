package rest

type Recorder interface {
	Record(interface{}) error
}
