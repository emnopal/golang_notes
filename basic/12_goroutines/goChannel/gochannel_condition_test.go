package go_channel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// ranged channel
/*
Kadang ada kasus, channel dikirim terus menerus oleh pengirim, dan tidak diketahui/tidak jelas kapan channel tsb
akan berhenti, nah untuk kasus ini pakein aja range channel dengan menggunakan perulangan sehingga channel tersebut
dapat diterima terus menerus (gak blocking), buat tutup range channel ini bisa pake function close
*/

func TestSimpleRangedChannel(t *testing.T) {
	channel := make(chan string)

	// creating up to 100 channel
	go func() {
		for i := 0; i <= 100; i++ {
			channel <- "Sending " + strconv.Itoa(i) + " st channel"
		}
		close(channel) // if result exceed 100 will close the channel
	}()

	for data := range channel {
		fmt.Println(data)
	}

}

// select channel
/*
ada kasus dimana kita buat beberapa channel dan menjalankan beberapa goroutine
lalu kita ingin mendapatkan semua data yg ada di semua channel tsb
nah ini bisa pakai select channel, untuk select channel ini, akan dipilih channel yg paling cepat dahulu
jika datang bersamaan maka akan di random (intinya select channel ini pilih nya random)
*/

func TestSimpleSelectChannel(t *testing.T) {
	counter := 0

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go func() {
		for i := 0; i <= 100; i++ {
			channel1 <- "Sending " + strconv.Itoa(i*1) + " channel"
			channel2 <- "Sending " + strconv.Itoa(i*2) + " channel"
			channel3 <- "Sending " + strconv.Itoa(i*3) + " channel"
		}
		close(channel1)
		close(channel2)
		close(channel3)
	}()

	for {
		select {
		case data := <-channel1:
			fmt.Println(counter, "(channel1)", data)
			counter++
		case data := <-channel2:
			fmt.Println(counter, "(channel2)", data)
			counter++
		case data := <-channel3:
			fmt.Println(counter, "(channel3)", data)
			counter++
		}
		if counter == 300 {
			break
		}
	}

}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func TestAnotherSelectChannel(t *testing.T) {

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}

// default channel
/*
ada kasus gimana kalau kita select ke channel yg gaada datanya? kalo misal gak kita handling bakalan error deadlock
nah cara handling nya pake default channel
*/

func ResponseChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello"
}

// error => gaada yg handle ketika function nya lagi sleep, error nya deadlock
func TestDefaultSelectChannelError(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go ResponseChannel(channel1)
	go ResponseChannel(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel 2", data)
			counter++
			if counter == 2 {
				break
			}
		}
	}
} // running: deadlock error

// success
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go ResponseChannel(channel1)
	go ResponseChannel(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel 2", data)
			counter++
		default: // untuk menghindari error ketika channel kita sleep, maka secara default selama function sleep, akan print bagian default ini
			fmt.Println("Wait...")
		}
		if counter == 2 {
			break
		}
	}
} // running: 2 sec

func TestDefaultSelectChannelOtherCase(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go ResponseChannel(channel1)
	go ResponseChannel(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel 2", data)
			counter++
		default: // untuk menghindari error ketika channel kita sleep, maka secara default selama function sleep, akan print bagian default ini
			continue // kalo gamau print apapun
		}
		if counter == 2 {
			break
		}
	}
} // running: 2 sec

// channel timeout
/*
Channel timeout, channel dalam waktu tertentu akan di close jika tidak ada aktivitas baik sending channel ataupun receive channe;
*/

func TestSimpleChannelTimeout(t *testing.T) {
	channel := make(chan int)

	// create infinite loop
	go func() {
		for i := 0; true; i++ {
			channel <- i
			time.Sleep(10 * time.Second) // but before it received by any receiver channel, it will be slept for 10sec
		}
	}()

	// receiving channel
	go func() {
	loop: // <- this called label; used for loop/break/continue
		for {
			select {
			case data := <-channel:
				fmt.Print(`receive data "`, data, `"`, "\n") // this, never be printed out
			case <-time.After(time.Second * 5):
				fmt.Println("timeout. no activities under 5 seconds") // this will break if no activities for 5sec
				break loop
			}
		}
	}()
}
