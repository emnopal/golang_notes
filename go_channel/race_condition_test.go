package go_channel

import (
	"fmt"
	"testing"
	"time"
)

/*
Ketika kita berurusan dgn concurrency, program itu tidak hanya berjalan secara concurrent saja, tapi bisa juga parallel,
karena beberapa thread bisa berjalan secara parallel, hal ini sangat berbahaya ketika kita melakukan manipulasi
data variable yg sama oleh beberapa goroutine secara bersamaan, hal ini bisa menyebabkan masalah yg namanya race
condition

kenapa begitu, karena concurrent itu jalan berdasarkan thread, nah thread bisa dieksekusi oleh compiler dengan
multiple cpu, jadi tiap sekali run bisa return beberapa thread, nah karena tiap thread ini bergantung sama variable
yang sama, makanya hal ini bahaya karena bisa jadi ketimpa sama hasil yg udah di run dan hasilnya tidak sesuai, dan juga
tiap kali run bisa punya hasil yg beda

istilah gampangnya: goroutine thread balapan buat ubah variable, karena dia pake shared variable
*/

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second) // sleep to make concurrent work

	fmt.Println("Counter:", x) // the result must be 100_000; but actually less or more than 100_000 and not equal 100_000
}

// solusinya pake sync.Mutex/sync.RMWutex atau bisa juga pake sync.Atomic
