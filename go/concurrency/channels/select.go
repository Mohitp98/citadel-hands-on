package channels

import (
	"fmt"
	"time"
)

func One(c1 chan<- string, c2 chan<- string) {
	for {
		c1 <- "from 1"
		time.Sleep(time.Second * 2)
	}
}

func Two(c1 chan<- string, c2 chan<- string) {
	for {
		c2 <- "from 2"
		time.Sleep(time.Second * 3)
	}
}

// receiving channel
func Three(c1 <-chan string, c2 <-chan string) {
	for {
		// selects works as a switch case for channels
		select {

		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)

		// time.After creates a channel and after the given duration will send the current time on it.
		case <-time.After(time.Second):
			fmt.Println("timeout")

			// The default case happens immediately if none of the channels are ready.
			// default:
			// 	fmt.Println("nothing ready")

		}

	}
}
