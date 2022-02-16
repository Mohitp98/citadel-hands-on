package main

import (
	"fmt"
	"sync"
	// "runtime"
	// "github.com/Mohitp98/citadel-hands-on/go/concurrency/waitgroups"
	// "github.com/Mohitp98/citadel-hands-on/go/concurrency/channels"
)

var wg sync.WaitGroup

// main funtion is itself a go routine, hence getting count 1
func main() {

	// {
	// 	fmt.Println("OS: ", runtime.GOOS)
	// 	fmt.Println("ARCH: ", runtime.GOARCH)
	// 	fmt.Println("CPUs: ", runtime.NumCPU())
	// 	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	// 	go foo()
	// }

	// ===================================== WaitGroups =====================================
	/*
		//	In following case program will shut down even if any go routine's
		//	execution is in progress! as main func exits
		//	ie. there is no synchronization between go routines (asynchronus)

		 {
		 	fmt.Println("OS: ", runtime.GOOS)
		 	fmt.Println("ARCH: ", runtime.GOARCH)
		 	fmt.Println("CPUs: ", runtime.NumCPU())
		 	fmt.Println("Goroutines: ", runtime.NumGoroutine())

		 	go waitgroups.Foo()

		 	fmt.Println("CPUs: ", runtime.NumCPU())
		 	fmt.Println("Goroutines: ", runtime.NumGoroutine())
		 }

	*/

	/*
		// to solve the above problem waitGroup is introduced
		{
			fmt.Println("OS: ", runtime.GOOS)
			fmt.Println("ARCH: ", runtime.GOARCH)
			fmt.Println("CPUs: ", runtime.NumCPU())
			fmt.Println("Goroutines: ", runtime.NumGoroutine())

			// we've mentioned the count of goroutines to wait for
			wg.Add(1)
			go waitgroups.Foo(&wg)
			waitgroups.Bar()

			fmt.Println("CPUs: ", runtime.NumCPU())
			fmt.Println("Goroutines: ", runtime.NumGoroutine())
			wg.Wait()
		}
	*/

	// ===================================== CHANNELS =====================================

	/*
		// channels-1
		{
			var c chan string = make(chan string)

			// providing channel `c` for Pinger and Printer
			go channels.Pinger(c)
			go channels.Ponger(c)
			go channels.Printer(c)

			var input string
			fmt.Scanln(&input) // to exit
		}
	*/

	/*
		// Select:
		// select picks the first channel that is ready and receives from it (or sends to it).
		// Case Channels Ready: then it randomly picks which one to receive from.
		// Case Channels Not Ready: the statement blocks until one becomes available.
		{
			var c1 chan string = make(chan string)
			var c2 chan string = make(chan string)

			go channels.One(c1, c2) // “from 1” every 2 seconds
			go channels.Two(c1, c2) // “from 2” every 3 seconds
			go channels.Three(c1, c2)

			var input string
			fmt.Scanln(&input) // to exit
		}
	*/

	{
		// range over channel
		c := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c) // stops writing to channel but can receive from it
		}()

		// pulls values from channel until the channel is close/empty
		for n := range c {
			fmt.Println(n)
		}
	}

}
