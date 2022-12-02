package go_channel

import (
	"fmt"
	"testing"
	"time"
)

// buffered channel

/*

Note for buffered channel (in Indonesia):

1. Proses transfer data ke dan pada channel by default pake un-buffered artinya tidak di-buffer di memori,
   sehingga tidak akan ada antrian, sehingga kita harus menunggu ketika channel sudah ada yg
   mengambil baru kita bisa masukan data yg lain
2. By default nilai buffer size dari un-buffered adalah 0
3. Buffered adalah proses yg agak berbeda dari un-buffered, jadi itu bisa buat antrian di dalam channel,
   jadi bisa kirim sekaligus ke dalam channel sesuai dengan jumlah buffer size
4. Ketika jumlah data yang dikirim sudah melewati batas buffer,
   maka pengiriman data hanya bisa dilakukan ketika salah satu data yang sudah terkirim adalah
   sudah diambil dari channel di goroutine penerima, sehingga ada slot channel yang kosong.

*/

func TestSimpleBufferedChannel(t *testing.T) {
	// by default, channel only accept 1 data
	// if there is another data imported into channel, we have to wait until 1st data is finished picked up by receiver
	// sometimes, there is case which sender is faster than receiver
	// this will cause application running slow or not properly running (or even will cause deadlock)
	// so for fix this case, use buffered channel to make queue

	// buffer capacity

	// buffer capacity has unlimited queue as long as your memory can handle
	// example: if we set buffer capacity into 5, so channel will handle only 5 data in queue
	// if we send 6th data to queue, we have to wait one data is finished handled to enter the queue

	// how to create buffer
	channel := make(chan string, 3)
	defer close(channel)

	fmt.Println(cap(channel)) // buffer size; 3
	fmt.Println(len(channel)) // how many occupied place in channel; 0 (0 from 3)
}

func TestNoBufferChannel(t *testing.T) {
	// why we use buffer? because the sender is too fast (to handle this we can use sleep, but not in every case)
	// example buffer:

	// with no buffer

	channelNoBuffer := make(chan string) // buffer size: 0
	defer close(channelNoBuffer)

	//channelNoBuffer <- "This channel has 0 buffer"
	//fmt.Println(<-channelNoBuffer) // this will cause deadlock, because receiver not have time enough to catch channel

	// wrap the channel into concurrent function
	go func() {
		channelNoBuffer <- "This channel has 0 buffer; 1st function; 1st"
	}()
	fmt.Println(<-channelNoBuffer)
	//fmt.Println(<-channelNoBuffer) <- will throw error; because the receiver not in concurrent function

	go func() {
		channelNoBuffer <- "This channel has 0 buffer; 2nd function"
	}()
	go func() {
		fmt.Println(<-channelNoBuffer)
		fmt.Println(<-channelNoBuffer) // concurrent will ignore this, so there is no error thrown
	}()

	// this is allowed syntax
	go func() {
		channelNoBuffer <- "This channel has 0 buffer; 3rd function"
	}()
	go fmt.Println(<-channelNoBuffer)
	go fmt.Println(<-channelNoBuffer) // concurrent will ignore this, so there is no error thrown

	time.Sleep(2 * time.Second) // to hold sender for sending channel into receiver, if you comment this, it will throw error

}

func TestBufferedChannel(t *testing.T) {

	bufferSize := 1
	channelWithBuffer := make(chan string, bufferSize)
	defer close(channelWithBuffer)
	// this buffer size has 1 place
	// so what happen if we add more channel into this buffer?
	// but we never use it

	channelWithBuffer <- "This channel has 1 buffer"
	//channelWithBuffer <- "This channel has 1 buffer; new buffer 1" // new channel; this will cause deadlock
	fmt.Println(<-channelWithBuffer)
	//fmt.Println(<-channelWithBuffer) // this also with cause deadlock, because channel has only 1 place to receive, but you receive it 2

	// if we have used it, and assign new buffer, it will be okay
	channelWithBuffer <- "This channel has 1 buffer; new buffer"
	fmt.Println(<-channelWithBuffer) // OKay

	// channel with more than 1 bufferSize; queue system
	channelWithBiggerBuffer := make(chan string, 3)
	defer close(channelWithBiggerBuffer)

	channelWithBiggerBuffer <- "This channel has 3 buffer; 1st queue"
	channelWithBiggerBuffer <- "This channel has 3 buffer; 2nd queue"
	channelWithBiggerBuffer <- "This channel has 3 buffer; 3rd queue"
	//channelWithBiggerBuffer <- "This channel has 3 buffer; 4th queue" // error

	fmt.Println(<-channelWithBiggerBuffer) // OKay
	fmt.Println(<-channelWithBiggerBuffer) // OKay
	fmt.Println(<-channelWithBiggerBuffer) // OKay
	//fmt.Println(<-channelWithBiggerBuffer) // Error

}

func TestConcurrentBufferedChannel(t *testing.T) {
	// buffered channel also works with concurrent function to make queue; similar with async await mechanism
	// example
	// note: if function make(chan string) is assigned without size, so it means it only accept max 0 queue
	bufferSize := 3
	channelBufferFunc := make(chan string, bufferSize) // accept max 2 queue, if there are more than 2 queue, they have to wait until all queue finished
	channelBufferFuncNums := make(chan int, bufferSize)
	// 0 -> max 0 queue // default size
	// 1 -> max 1 queue
	// and so on

	defer close(channelBufferFunc)
	defer close(channelBufferFuncNums)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Sending channel no: ", i)
			channelBufferFunc <- fmt.Sprintf("Channel %d", i)
			channelBufferFuncNums <- i
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Receiving channel no:", <-channelBufferFuncNums, "(Channel Name:", <-channelBufferFunc, ")")
		}
	}()

	// the result => they only can receive up to 4 channel

	time.Sleep(2 * time.Second) // hold sender channel; if you comment this, it will deadlock

}

func TestBlockingBufferChannel(t *testing.T) {

	// this function will be error due to closed channel,
	// because there is no enough receiver since channel only can hold 1 channel (see: sizeBuffer)

	// but note: the error will appear only in sender side
	// if you assign, for example more than 5 blockingChannelAsync to fmt.Println, it will not throw an error

	sizeBuffer := 1 // if you set it to 4 it will be okay, 1 will go to output and other will remain sleep
	blockingChannelAsync := make(chan string, sizeBuffer)

	defer close(blockingChannelAsync)

	go func() {
		// channel send 5 channel
		blockingChannelAsync <- "Testing Blocking Async 1"
		blockingChannelAsync <- "Testing Blocking Async 2"
		blockingChannelAsync <- "Testing Blocking Async 3"
		blockingChannelAsync <- "Testing Blocking Async 4"
		blockingChannelAsync <- "Testing Blocking Async 5"
	}()

	go func() {
		// only receive 1 channel
		fmt.Println(<-blockingChannelAsync) // <- Testing Blocking Async 1

		// uncomment this
		//fmt.Println(<-blockingChannelAsync) // <- Testing Blocking Async 2
		//fmt.Println(<-blockingChannelAsync) // <- Testing Blocking Async 3
		//fmt.Println(<-blockingChannelAsync) // <- Testing Blocking Async 4
		//fmt.Println(<-blockingChannelAsync) // <- Testing Blocking Async 5
		//fmt.Println(<-blockingChannelAsync) // <- ""
		//fmt.Println(<-blockingChannelAsync) // <- ""
		//fmt.Println(<-blockingChannelAsync) // <- ""
	}()

	time.Sleep(2 * time.Second)

}

func TestAnotherBufferChannelExample0(t *testing.T) {
	messages := make(chan int, 3)
	defer close(messages)

	go func() {
		for i := 0; i < 100; i++ {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	// no need to sleep, because sender is not use concurrent
	for i := 0; i < 100; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

}

func TestAnotherBufferChannelExample1(t *testing.T) {
	message := make(chan int, 3)
	defer close(message)

	go func() {
		for i := 0; i < 100; i++ {
			i := <-message
			fmt.Println("receive data", i)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("send data", i)
			message <- i
		}
	}()

	time.Sleep(2 * time.Second) // hold receiver, because sender is using concurrent

}
