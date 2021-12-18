package main

import (
	"fmt"
	"time"

	"e.coding.net/kunz/startkit-x/go-co"
	"gopkg.in/ffmt.v1"
)

func sleepAsync(index int, ms int64) *co.Task {
	return co.Async(func() interface{} {
		ffmt.Println(fmt.Sprintf("co %d staring.", index))
		time.Sleep(time.Millisecond * time.Duration(ms))
		ffmt.Println(fmt.Sprintf("co %d End.", index))
		return nil
	})
}

func main() {
	tick := time.Tick(time.Second / 100)
	for range tick {
		fmt.Println("before sleep : ", time.Now())
		t1 := sleepAsync(1, 1000)
		t2 := sleepAsync(2, 1000)
		t3 := sleepAsync(3, 1000)
		t4 := sleepAsync(4, 4000)

		co.Await(t1)
		co.Await(t2)
		co.Await(t3)
		co.Await(t4)
		fmt.Println("after sleep : ", time.Now())
		ffmt.Println("co finish. Done.")
	}
}
