package synchronization

import (
	"fmt"
	"sync"
)

func Sync() {
	fmt.Println("Hello World from Synchronization!")

	fmt.Println("--- Sync with Mutex ---")
	syncWithMux()

	fmt.Println("~~~~~~~~~~~ End of Synchronization ~~~~~~~~~~~")
}

var (
	counter int
	mutex   sync.Mutex
)

func increment() {
	mutex.Lock()         // Acquire the lock
	defer mutex.Unlock() // Release the lock when the function exits
	counter++
	fmt.Printf("Incremented counter to: %d\n", counter)
}

func syncWithMux() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
