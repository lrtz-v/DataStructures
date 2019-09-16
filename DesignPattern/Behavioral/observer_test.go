package behavioraldesignpattern

import (
	"testing"
	"time"
)

func TestObserser(t *testing.T) {
	notifier := &eventNotifier{
		observers: make(map[Observer]struct{}),
	}

	notifier.Register(&eventObserver{id: 1})
	notifier.Register(&eventObserver{id: 2})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			notifier.Notify(Event{Data: t.UnixNano()})
		}
	}

}
