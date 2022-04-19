package main

import "sync"

func testWG(output chan<- []bool) {
	// TODO: answer here
	var wg sync.WaitGroup

	done := make([]bool, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		// TODO: answer here
		go func(i int) {
			wg.Done()
			// TODO: answer here
			done[i] = true
		}(i)
	}
	wg.Wait()
	// TODO: answer here
	output <- done
}
