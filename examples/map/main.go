package main

import (
	"fmt"
	"time"

	"e.coding.net/kunz/startkit-x/go-co"
	"e.coding.net/kunz/startkit-x/go-co/coutil"
)

func sleep(sec int) *co.Task {
	return co.Async(func() interface{} {
		time.Sleep(time.Second * time.Duration(sec))
		return nil
	})
}

func main() {
	items := []interface{}{1, 2, 3, 4}
	fmt.Println("before work : ", time.Now())

	// with concurrency 2
	t := coutil.Map(items, func(item interface{}, index int, arr []interface{}) *co.Task {
		return co.Async(func() interface{} {
			co.Await(sleep(item.(int)))
			return item.(int) * item.(int)
		})
	}, 2)

	res, _ := co.Await(t)
	fmt.Println("after work : ", time.Now())

	fmt.Println(res)
}
