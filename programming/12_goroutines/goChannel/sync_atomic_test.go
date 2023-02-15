package go_channel

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Atomic -> package yg digunakan untuk menggunakan data primitive secara aman pada process concurrent
sebenernya buat secure selain kita bisa pakai locking di Mutex, RWMutex atau Cond, kita juga bisa modified
data primitive kita dengan package sync/atomic ini
*/

func TestAtomic(t *testing.T) {
	var x uint64 = 0
	var group sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddUint64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println("Counter:", x) // the result must be 100_000; but actually less or more than 100_000 and not equal 100_000
}

// pakai ini juga lebih cepat daripada harus lock unlock
