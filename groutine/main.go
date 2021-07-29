package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 0; i <5 ; i++ {
		go test(i,c)
	}
	for i := 0; i < 5; i++ {
		r := <- c
		fmt.Println(r)
	}
}


func test(id int, c chan int){
	fmt.Println("go",id)
	c <- id
}