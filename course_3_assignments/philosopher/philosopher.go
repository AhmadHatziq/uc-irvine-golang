package main

import (
	"fmt"
	"sync"
	"time"
)

// Define chopsticks and philosopher structs
type ChopS struct {
	sync.Mutex
}
type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
	name    int
}

// Host method. Host will fill up the request_channel first.
// Then, host will keep checking for anything sent to the finish_channel by the philo.
// Philo will send to the finish_channel when they are done.
// When the host receives something in the finish_channel, he will add a new element into the buffer of the request_channel.
func host(request_channel chan int, finish_channel chan int) {

	// Sends 2 elements into the request channel to fill the buffer
	request_channel <- 1
	request_channel <- 1

	// Whenever the philo has finished eating, the philo will add a number to the finish_channel
	// The host can refill the request channel with 1 element in the buffer
	for {
		<-finish_channel
		request_channel <- 1
	}

}

// Define eating method.
// Philo will first request in the request_channel. Will block if buffer is empty.
// After eating, philo will send a request to the finish_channel.
func (p Philo) eat(wg *sync.WaitGroup, request_channel chan int, finish_channel chan int) {

	// Eat only 3 times
	for i := 0; i < 3; i++ {

		// fmt.Println("Waiting for permission for philo: ", p.name)

		// Asks for permission from host.
		<-request_channel

		// Grab chopsticks ie lock
		p.leftCS.Lock()
		p.rightCS.Lock()

		// Philo is eating
		fmt.Println("starting to eat ", p.name)
		// fmt.Println("Philo num ", p.name, " is eating for time ", i+1)
		time.Sleep(100 * time.Millisecond)

		// Release chopsticks ie unlock
		fmt.Println("finishing eating: ", p.name)
		p.leftCS.Unlock()
		p.rightCS.Unlock()

		// Tells host we are done eating.
		finish_channel <- 1
		// fmt.Println("Philo num: ", p.name, " finished eating for time ", i+1)

	}

	// fmt.Println("Philo ", p.name, " is done eating everything")
	wg.Done()
}

func main() {

	// Create a sync.WaitGroup
	var wg sync.WaitGroup

	// Create 2 channels.
	// request_channel is set to buffer 2. Is for philo to request permission to eat.
	// With buffer 2, only 2 philo can eat at the same time.
	// finish_channel is for philo to declare they have finished eating. Buffer is set to a large number
	request_channel := make(chan int, 2)
	finish_channel := make(chan int, 100)

	// Initialize 5 pairs of chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Initialize 5 Philosophers with their chopsticks
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {

		// Create philo objects. Each philo is numbered 1-5
		philos[i] = &Philo{
			CSticks[i],
			CSticks[(i+1)%5],
			i + 1,
		}
	}

	// Host will run indefinitely. No need to wait. Hence, no adding wg.Add(1)
	go host(request_channel, finish_channel)

	// Let the 5 philosophers eat
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, request_channel, finish_channel)
		wg.Add(1)
	}
	wg.Wait()

}
