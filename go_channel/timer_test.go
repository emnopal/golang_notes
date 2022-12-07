package go_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
delay job
*/

func TestNewTimer(t *testing.T) {
	timer := time.NewTimer(25 * time.Second)
	timeNow := time.Now()
	fmt.Println(timeNow)

	timeChannel := <-timer.C
	fmt.Println(timeChannel)

	timeDelta := timeChannel.Sub(timeNow)
	fmt.Println(timeDelta)
}

// ini sama kaya yg tadi, cuma ini kalau kita cuma butuh si channel nya aja
func TestTimerAfter(t *testing.T) {
	timer := time.After(25 * time.Second)
	fmt.Println(<-timer)
}

// ini kalau kita mau buat function tersebut nunggu sampai beberapa waktu,
// ingat ini function nya bentuk data type nya harus func{} begini tanpa return dan parameter
// kalau mau di function yg pake return dan parameter tinggal pake time.Sleep aja
func TestTimerAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("5 second elapsed")
		group.Done()
	})
	group.Wait()
}

// nice explanation, thanks to: https://stackoverflow.com/a/31936175/10925834
func TimerTicker() chan bool {

	done := make(chan bool, 1)

	go func() {
		// ticker itu kejadiannya berulang
		// ketika ticker udah expired, event akan dikirim ke channel
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		select {
		case tick := <-ticker.C:
			fmt.Println(tick)
		case <-done:
			fmt.Println("Done")
			return
		}
	}()

	return done
}

func TestTimerTicker(t *testing.T) {
	stop := TimerTicker()
	time.Sleep(10 * time.Second)
	close(stop)
}
