package waitgroups

import (
	"fmt"
	"sync"
)

func Foo(wg *sync.WaitGroup) {
	fmt.Println("Inside foo()")
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	// informing to wg once the task is done!
	wg.Done()
}

func Bar() {
	fmt.Println("Inside bar()")
	for i := 10; i < 20; i++ {
		fmt.Println("bar: ", i)
	}
	// time.Sleep(100 * time.Millisecond)
}
