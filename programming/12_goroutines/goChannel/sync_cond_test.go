package go_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Cond
/*
Cond -> conditional locking
Cond butuh locker juga kaya locking biasa (bisa pake Mutex atau RWMutex)
tapi bedanya Cond ada method Wait() untuk stating apakah harus ditungguin atau enggak goroutine yg di lock
di Cond juga ada method Signal() -> ngasitau satu goroutine buat gausah nungguin goroutine yg di lock
dan juga ada method Broadcast() -> ngasitau semua goroutine yg ada buat gausah nungguin goroutine yg di lock
*/

// ini bakalan deadlock
// karena:
func TestCondNotProper(t *testing.T) {

	var mutex = &sync.Mutex{}
	var cond = sync.NewCond(mutex)
	var group = sync.WaitGroup{}

	WaitCondition := func(value int) {
		defer group.Done() // close/delete group
		group.Add(1)       // add group

		cond.L.Lock()              // locking
		cond.Wait()                // nungguin apakah setelah sukses locking, dengan kondisi, boleh dilanjut atau enggak
		fmt.Println("Done", value) // kalo boleh lanjut, goroutine yg lain, bakalan execute ini
		cond.L.Unlock()            // ini unlock
	}

	for i := 0; i <= 10; i++ {
		go WaitCondition(i) // <- disini, ada cond.Wait() nah karena cond.Wait() ini berdasarkan condition,
		// jadi kalau gaada match condition dia bakalan nunggu terus, nah jadinya deadlock
	}

	group.Wait()

}

// jadi kita harus tambahin condition, entah dengan signal ataupun broadcast
func TestCondProperSignal(t *testing.T) {

	var mutex = &sync.Mutex{}
	var cond = sync.NewCond(mutex)
	var group = sync.WaitGroup{}

	WaitCondition := func(value int) {
		defer group.Done() // close/delete group
		group.Add(1)       // add group

		cond.L.Lock()              // locking
		cond.Wait()                // nungguin apakah setelah sukses locking, dengan kondisi, boleh dilanjut atau enggak
		fmt.Println("Done", value) // kalo boleh lanjut, goroutine yg lain, bakalan execute ini
		cond.L.Unlock()            // ini unlock
	}

	for i := 0; i <= 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // <- pastikan ada yg kirim signal ke cond.Wait()
		}
	}()

	group.Wait()

	// dia bakalan keluar satu-satu
	// keluar satu2 karena pake signal, jadi func nya disuruh wait satu persatu

}

func TestCondProperBroadcast(t *testing.T) {

	var mutex = &sync.Mutex{}
	var cond = sync.NewCond(mutex)
	var group = sync.WaitGroup{}

	WaitCondition := func(value int) {
		defer group.Done() // close/delete group
		group.Add(1)       // add group

		cond.L.Lock()              // locking
		cond.Wait()                // nungguin apakah setelah sukses locking, dengan kondisi, boleh dilanjut atau enggak
		fmt.Println("Done", value) // kalo boleh lanjut, goroutine yg lain, bakalan execute ini
		cond.L.Unlock()            // ini unlock
	}

	for i := 0; i <= 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast() // <- pastikan ada yg kirim signal ke cond.Wait()
		}
	}()

	group.Wait()

	// ini hanya 1x di wait, selanjutnya di ignore cond.Wait() nya, jadi keluarnya bebarengan dan mirip locking biasa

}
