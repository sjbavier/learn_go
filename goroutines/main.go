package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
	for i := 0; i < 10; i++ {

		fmt.Println(text)
		time.Sleep(400 * time.Millisecond)
	}
}
func printMessageChannel(text string, the_channel chan string) {
	for i := 0; i < 10; i++ {

		fmt.Println(text)
		time.Sleep(400 * time.Millisecond)
	}
	the_channel <- "Done!"
}
func main() { // main goroutine, if this finishes before the end of other goroutines within, they all end

	var the_channel chan string
	go printMessageChannel("go channel", the_channel)
	// <-the_channel // alternatively wait without using value
	response := <-the_channel // it will wait for the channel to return
	fmt.Println(response)

	// go printMessage("Goooo") // only 5 will print out in 2000 milliseconds
	// time.Sleep(2000 * time.Millisecond)

}
