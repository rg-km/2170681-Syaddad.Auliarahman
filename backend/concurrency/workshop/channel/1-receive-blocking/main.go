package main

import (
	"fmt"
	"time"
)

//karena sifat channel berukuran 0 yang blocking
//dapat digunakan channel untuk menunggu goroutine
func main() {
	c := make(chan int)
	go func() {
		fmt.Println("sending to channel c")
		time.Sleep(time.Second * 10)

		c <- 1 // blocking hingga ada yang menerima data
	}()

	select {
	case msg := <-c:
		fmt.Print(msg)
	case msg1 := <-c:
		fmt.Println(msg1)
	}
	// msg := <-c //blocking
	// karena terjadi blocking, goroutine #2 akan dijalankan
	// ketika goroutine #2 mengirim ke channel c, maka main goroutine jalan kembali
	fmt.Println("main stop")
}
