/*
Unbuffered Channels:
- Synchronize way of communication between the goroutines
- ie. when pinger(sender) sends the msg to channel,
    it will wait until the printer(receiver) is ready to receive the message.
- also called as `blocking`
- Buffer capacity: default is zero, for an unbuffered or synchronous channel.
*/

package channels

import (
	"fmt"
	"time"
)

// using `chan<-` in signature to restrict func to only sending the msg
func Pinger(c chan<- string) {
	for i := 0; ; i++ {
		// sending msg through channel!
		c <- "ping"
	}
}

// using `<-chan` in signature to restrict func to only sending the msg
func Printer(c <-chan string) {
	for {
		// capturing msg through channel!
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func Ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
