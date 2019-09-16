package behavioraldesignpattern

import "log"

type (
	Event struct {
		Data int64
	}

	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observer)
		Deregister(Observer)
		Notify(Event)
	}

	eventObserver struct {
		id int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	log.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}
