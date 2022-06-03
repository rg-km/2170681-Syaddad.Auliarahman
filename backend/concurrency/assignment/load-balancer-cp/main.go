package main

import (
	"fmt"
	"sync"
)

const nWorkers = 3
const nRequesters = 5

type Work struct {
	data string
}

var mu = &sync.Mutex{}

var data = map[string]bool{}

func worker(in <-chan *Work, out chan<- *Work, number int) {
	for w := range in {
		w.data = fmt.Sprintf("worker %d accepted request from %s", number, w.data)
		out <- w
		// TODO: answer here
	}
}

func createRequest(in chan<- *Work, number int) {
	for {
		// TODO: answer here
		in <- &Work{data: fmt.Sprintf("client %d", number)}

	}
}

func receiver(out <-chan *Work) {
	for {
		//pada bagian ini masukkan data dari channel out
		//ke dalam map `data`
		//gunakan mutex untuk mengamankan penulisan ke map secara concurrent
		mu.Lock()
		cleaned := <-out
		data[cleaned.data] = true // TODO: replace this
		<-out
		mu.Unlock()
		// TODO: answer here
	}
}
