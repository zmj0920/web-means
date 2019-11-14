package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			mut.Unlock()
		}()
		mut.Lock()
		say("hello")
		wg.Done()
	}()
	go say("world")
	wg.Wait()
}
