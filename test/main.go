package main

import (
	"fmt"
	"time"
)

func main() {
	JackTasks := make(chan int, 1)

	var taskID int
	fmt.Scan(&taskID)

	JackTasks <- taskID
	timer := time.NewTimer(time.Second * 1)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Printf("Jack has finished task (%d)", <-JackTasks)
			return
		default:
			fmt.Println("Jack is working")
			//time.Sleep(time.Second)
		}
	}
}
