package go_channel

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Map
/*
ini mirip sama Map biasa di golang tapi ini lebih aman buat concurrent programming, karena aman dari race condition
method nya pun mirip, ada Store(key, value), Load(key), Delete(key), Range(func(key,value))
*/

func TestSyncMap(t *testing.T) {
	var data sync.Map
	var group sync.WaitGroup

	addToMap := func(data *sync.Map, value int, group *sync.WaitGroup) {
		defer group.Done()
		group.Add(1)
		data.Store(fmt.Sprint("Key of ", value), value)
	}

	for i := 0; i <= 100; i++ {
		go addToMap(&data, i, &group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
