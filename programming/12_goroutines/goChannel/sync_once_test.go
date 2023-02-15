package go_channel

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Once
/*
memastikan bahwa function harus dijalankan hanya sekali saja dan tidak akan pernah di eksekusi lagi (di runtime yg sama ya)
jadi mau berapa banyakpun goroutine yg akses hanya goroutine pertama yg bisa execute, yg lain nya di ignore
*/

var counter = 0

func DontYouDareToExecuteThisTwice() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			once.Do(DontYouDareToExecuteThisTwice) // jadi ini cuma bisa di eksekusi sekali,
			// tapi ingat disini type function nya harus func () <- tanpa parameter dan tanpa return
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter) // 1; selain 1 pasti salah
}
