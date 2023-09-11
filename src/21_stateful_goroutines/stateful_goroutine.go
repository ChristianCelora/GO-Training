package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	value int
	resp chan bool
}

/**
* Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result. 
* This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
*/
func main() {
	var nReadOps uint64
	var nWriteOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	/**
	* Here is the goroutine that owns the state. 
	* This goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive. 
	* A response is executed by first performing the requested operation and then sending a value on the response channel resp to indicate success (and the desired value in the case of reads).
	*/
	go func () {
		var state = make(map[int]int)
		for {
			select {
			case read := <- reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	/**
	* This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel. 
	* Each read requires constructing a readOp, sending it over the reads channel, and then receiving the result over the provided resp channel.
	*/
	for r := 0; r < 100; r++ {
		go func(){
			for {
				read := readOp{ 
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&nReadOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	/**
	* We start 10 writes as well, using a similar approach.
	*/
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					value: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&nWriteOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	time.Sleep(time.Second)
    readOpsFinal := atomic.LoadUint64(&nReadOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&nWriteOps)
    fmt.Println("writeOps:", writeOpsFinal)
	fmt.Println("END")
}