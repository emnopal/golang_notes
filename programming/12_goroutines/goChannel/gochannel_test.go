package go_channel

import (
	"fmt"
	"testing"
	"time"
)

// simple channel

func ChannelAsParamsFunc(str string, strChan chan string) {
	time.Sleep(5 * time.Second)
	strChan <- str
}

func TestSimpleChannel(t *testing.T) {
	// create channel in golang is using chan keyword
	message := make(chan string, 1) // sender
	message <- "Test"               // this execution will be blocked if you don't specify any receiver
	msg := <-message                // receiver; if receiver isn't receiving channel "message" the execution will be blocked
	fmt.Println(msg)                // or if there is only receiver but sender is not defined the execution will be blocked too
	close(message)                  // close channel using close function
}

func TestChannel(t *testing.T) {

	// create channel in golang is using chan keyword
	channel := make(chan string, 3) // what is parameter 3 in make function? see buffered_test.go

	// to close channel using defer
	defer close(channel)

	// how to use channel; 1st method
	channel <- "sending to channel 1st usage"
	channelRec := <-channel
	fmt.Println(channelRec)

	// how to use channel; 2nd method
	channel <- "sending to channel 2nd usage"
	var otherRecChannel string
	otherRecChannel = <-channel
	fmt.Println(otherRecChannel)

	// how to use channel; 3rd method
	channel <- "sending to channel 3rd usage"
	fmt.Println(<-channel)

	// (cont) receive channel as function parameter
	// channel by default is pass by reference
	// so in the channel you don't need to use pointer
	// another example of channel as function parameter
	ChannelAsParams := make(chan string) // no need to specify the size since this channel is wrapped in concurrent function
	ParamsStr := "Test ParamsStr from concurrent function"
	go ChannelAsParamsFunc(ParamsStr, ChannelAsParams)
	fmt.Println(<-ChannelAsParams)
	close(ChannelAsParams)

	// channel with concurrency
	ChannelConc := make(chan string)

	// create concurrent anonymous function
	// sending data to channel
	go func() {
		time.Sleep(2 * time.Second)
		ChannelConc <- "Hello from anonymous function" // channel inside concurrent function
		// if you use channel you don't need return value, just send data to channel and receive into another place
	}()

	// receive data from channel
	data := <-ChannelConc
	fmt.Println(data)

	close(ChannelConc)

}

// constrain the function or method with only send or receive the data

// only send
func OnlySend(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Running function which only for sending a channel"
	// fmt.Println(<-channel) // this will error; Invalid operation: <-channel (receive from the send-only type chan<- string)
	// the error is because you receive the channel in send only function
}

// only receive
func OnlyReceive(channel <-chan string) {
	// channel <- "hello" // this will error; Invalid operation: channel <- "This function only for sending a channel" (send to the receive-only type <-chan string)
	// the error is because you send the channel in receive only function
	fmt.Println(<-channel)
	fmt.Println("Running function which only for receive a channel")
}

func TestOnlySendReceiveChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlySend(channel)
	go OnlyReceive(channel)

	time.Sleep(5 * time.Second) // to hold; sleep for making concurrent work
}

// more example of Channel

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestSumChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	defer close(c)

	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c; you also can write receiver of channel like this

	fmt.Println(x, y, x+y)
}

func TestTest(t *testing.T) {
	a, b := 1, 2

	var c int
	var d int
	c, d = 4, 5

	fmt.Println(a, b)
	fmt.Println(c, d)
}
