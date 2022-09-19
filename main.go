package main

import (
	"fmt"
)

func main() {

	// Create a buffered channel of size the maximum number of go routines
	// that can run at one instance
	ch := make(chan int, 10)
	// defer close(ch)
	// Accept request and process it
	go process(ch)

	for v := range ch {
		fmt.Println("Read value of ID", v, "from ch")
		//Blocked the channel
		func() {
			fmt.Println("Remove value of ID", v, "from ch")
			<-ch
		}()
	}

}

func process(ch chan int) {
	// send 20 ids to be processed
	for i := 1; i < 20; i++ {
		ch <- i
		fmt.Println("successfully added ID", i, "to ch")
	}
	close(ch)
}
