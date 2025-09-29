package event

// Event 通用事件接口
type Event interface {
	GetEventID() int
}

type BaseEvent struct {
	EventID int
}

func (e *BaseEvent) GetEventID() int {
	return e.EventID
}
