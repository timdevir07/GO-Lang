package main

import "fmt"

func main() {
	//only for loop use in Go loop
	//while loop
	i := 1
	for i <= 5 {
		//break
		// if i == 2 {
		// 	continue
		// }
		fmt.Println(i)
		i = i + 1
	}
}
