package main

import (
	"fmt"
	"time"
)

//mengembalikan hasil pangkat dua angka 1-10
//dapat dilihat tidak terjadi blocking selama buffer belum penuh
var person = "andi"
var names = []string{"budi", "toni", "adi", "ado", "alif", "yudi"}

func squareWorker(output chan<- int) {
	for i := 1; i < 12; i++ {
		//fmt.Println(i * i)
		output <- i * i
	}
	fmt.Println("selesai mengirim")
}

func main() {
	output := make(chan int, 10)
	//output1 := make(chan string)
	go squareWorker(output)
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10; i++ {
		fmt.Println("main receive output:", <-output)
	}
	//greetAll(person, names, output1)
	// time.Sleep(100 * time.Millisecond)
	// for i := 0; i < 6; i++ {
	// 	greeting := <-output
	// 	result <- greeting
	// }
}
