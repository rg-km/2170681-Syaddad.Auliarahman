package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {

	// TODO: answer here
	c := make(chan int)
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//defer wg.Done()
		go func() {
			//kirim 1 ke channel
			defer wg.Done()
			c <- 1
			// TODO: answer here
		}()
	}

	//mengubah nilai count menggunakan data dari channel
	// TODO: answer here
	go func() {
		for {
			//fmt.Println(<-c)
			count += <-c
		}
	}()
	wg.Wait() // menunggu seluruh goroutine selesai berjalan
	output <- count

}
