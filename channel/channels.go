package channel

import "fmt"

func FuncChannel() {
	produceConsumeUsingChannel()
}

func producer(ch chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Produced:", i)
		ch <- "Value - " + fmt.Sprint(i) // send value to channel
	}
	close(ch) // important to signal end
}

func consumer(ch <-chan string) {
	for val := range ch {
		fmt.Println("Consumed:", val)
	}
}
func produceConsumeUsingChannel() {
	ch := make(chan string, 2)
	go producer(ch)
	consumer(ch)
}
