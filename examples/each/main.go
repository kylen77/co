package main

import (
	"fmt"
	"time"

	"e.coding.net/kunz/startkit-x/go-co"
	"e.coding.net/kunz/startkit-x/go-co/coutil"
)

func main() {
	items := []interface{}{1, 2, 3, 4, 5}

	// new Task
	t := coutil.Each(items, func(item interface{}, index int, items []interface{}) *co.Task {
		return co.Async(func() interface{} {
			s := item.(int)
			time.Sleep(time.Second * time.Duration(s))
			fmt.Printf("[%d] = %d\n", index, item)
			return s
		})
	})

	fmt.Println(co.Await(t))
}
