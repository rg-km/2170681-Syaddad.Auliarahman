package main

import "fmt"

func sendBlock(output chan bool, tes chan int) {
	c := make(chan bool)
	called := false
	called1 := 0
	go func() {
		fmt.Println("receive from main")
		//memberi called nilai dari channel c
		// TODO: answer here
		fmt.Println(<-c)
	}()

	//mengirim bool value true ke channel c
	// TODO: answer here
	called = true
	called1 = 666
	output <- called
	tes <- called1
	fmt.Println(c) //agar variabel c digunakan
}
