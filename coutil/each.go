package coutil

import "e.coding.net/kunz/startkit-x/go-co"

// Each: coutil.Each iterate a array
//
// example:
//  coutil.Each(arr, func(item, index, arr) { ... })
func Each(
	items []interface{},
	fn func(interface{}, int, []interface{}) *co.Task,
) *co.Task {
	return co.Async(func() interface{} {
		ret := make([]interface{}, len(items))
		for index, item := range items {
			var err error
			t := fn(item, index, items)
			ret[index], err = co.Await(t)
			if err != nil {
				panic(err)
			}
		}
		return ret
	})
}
