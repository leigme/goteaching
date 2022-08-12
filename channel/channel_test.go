package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_newTask(t *testing.T) {
	sc := make(chan struct{})
	fn := func() {
		fmt.Println("Hello World!")
		time.Sleep(3 * time.Second)
	}

	go newTask(sc, fn)
	time.Sleep(10 * time.Second)
	close(sc)
	fmt.Println("test stop")
}
