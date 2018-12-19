package event

type EventType int

const (
	INVALID EventType = iota
	ADD
	SUB
)

var eventStringMap = [...]string{
	"INVALID",
	"ADD",
	"SUB",
}

func (et EventType) String() string {
	if int(et) < len(eventStringMap) {
		return eventStringMap[et]
	}
	return eventStringMap[0]
}

type Event struct {
	Type   EventType
	Source interface{}
	Value  interface{}
}
