package channel

import (
	"time"
)

func newTask(stop chan struct{}, fn func()) {
	for {
		select {
		case <-stop:
			return
		case <-time.After(1 * time.Second):
			fn()
		}
	}
}
