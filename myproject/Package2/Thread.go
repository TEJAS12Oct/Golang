package Package2

import (
	"fmt"
	"time"
)

func task1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Task 1 completed"
}

func task2(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "Task 2 completed"
}

func RunTasks() []string {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	results := make([]string, 0, 2)

	for len(results) < 2 {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
			results = append(results, msg1)
			ch1 = nil
		case msg2 := <-ch2:
			fmt.Println(msg2)
			results = append(results, msg2)
			ch2 = nil
		}
	}

	return results
}
