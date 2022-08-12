package channel

import "fmt"

func newTask(stop chan struct{}, fn func()) {
	for {
		select {
		case <-stop:
			fmt.Println("func stop...")
			return
		default:
			fn()
		}
	}
}
