package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var result int

//fungsi ini digunakan untuk menerima angka lalu mengembalikan hasil pangkat 2 angka tersebut
func squareWorker(workerInput <-chan int, workerOutput chan<- int) {
	for {
		num := <-workerInput      // menerima angka
		workerOutput <- num * num // mengirim hasil
	}
}

func createRequest(workerInput chan<- int, workerOutput <-chan int, wg *sync.WaitGroup) {
	for i := 1; i < 100; i++ {
		// TODO: answer here
		workerInput <- i
		go func(i int) {
			// TODO: answer here
			var res int

			// TODO: answer here
			mu.Lock()
			res = <-workerOutput
			mu.Unlock()
			//tambahkan res ke result. Selain itu gunakan juga sesuatu yang menghindari data race
			// TODO: answer here
			result += res
			wg.Wait()
			fmt.Println(res)
		}(i)
	}
	wg.Done()
}

func main() {
	workerInput := make(chan int)
	workerOutput := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go squareWorker(workerInput, workerOutput)
	go createRequest(workerInput, workerOutput, &wg)
	wg.Wait()
	mu.Lock()
	fmt.Println(result)
	mu.Unlock()
}
