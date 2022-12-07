package go_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.WaitGroup
/*
WaitGroup adalah fitur buat nunggu proses goroutine sampai selesai dilakukan
sebenernya bisa pakai sleep, tapi time.Sleep ini bad practice dan sangat tidak direkomendasikan
karena sleep ini harus pakai parameter waktunya, dan kita gatau kapan aplikasi itu selesai di run
jadi bakalan ribet klo pake sleep
*/

func RunAsyncTestFunc(group *sync.WaitGroup) { // this function will run for 5 second
	defer group.Done() // close WaitGroup; ketika ada group.Wait tapi gaada group.Done maka akan deadlock

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(5 * time.Second)
}

// best practice
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncTestFunc(group)
	}

	group.Wait()        // wait until goroutine above finish; we don't need to specify time for the function running
	fmt.Println("Done") // this will be run after go RunAsyncTestFunc(group) finished
}

// bad practice
func TestWaitGroupTime(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncTestFunc(group)
	}

	time.Sleep(5 * time.Second) // we have known above goroutine is running for 5 seconds, so if we don't know how we measure it?
	fmt.Println("Done")         // this will be run after go RunAsyncTestFunc(group) finished
}

// don't you dare to use this
func TestWithoutWait(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncTestFunc(group)
	}

	fmt.Println("Done") // this will print randomly among running goroutine
	// so above line is not wait until some function is finished
}
