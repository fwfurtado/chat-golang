package stream

type Stream interface {
	Subscribe(id, channel string) <-chan interface{}
	Emit(channel string, data interface{})
	Unsubscribe(id string)
	UnsubscribeAll()
	Start()
}
