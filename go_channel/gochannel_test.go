package go_channel

import (
	"fmt"
	"testing"
)

func TestSimpleChannel(t *testing.T) {

	// create channel in golang is using chan keyword
	message := make(chan string, 1) // sender
	message <- "Test"
	msg := <-message // receiver
	fmt.Println(msg)
	close(message) // close channel using close function

	// best practice way

	// create channel in golang is using chan keyword
	channel := make(chan string, 3) // make(chan <Type>, size -> mandatory) // specify your size as you need to avoid deadlock
	// this channel will be consuming by 2 variable and 1 function

	// to close channel using defer
	defer close(channel)

	// 1st usage; send data to channel, cannot be re-usage, or it will be deadlock
	channel <- "sending to channel 1st usage"

	// 1st usage; receive data from channel
	channelRec := <-channel
	fmt.Println(channelRec)

	// 2nd usage; send data to channel, cannot be re-usage, or it will be deadlock
	channel <- "sending to channel 2nd usage"

	// 2nd usage; or
	var otherRecChannel string
	otherRecChannel = <-channel
	fmt.Println(otherRecChannel)

	// 3rd usage; send data to channel, cannot be re-usage, or it will be deadlock
	channel <- "sending to channel 3rd usage"

	// 3rd usage; channel as function parameter
	fmt.Println(<-channel)
}
