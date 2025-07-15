package main

import "fmt"

func main() {
	//while loop in go
	/* i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}*/

	//output
	/*
	   PS C:\Users\Raju Kumar\OneDrive\Desktop\Go Language\3_loop> go run loop.go
	   1
	   2
	   3
	   4
	   5
	*/

	//infinite loop in go lan

	/*for {
		fmt.Print("#")
	}*/

	//output

	/*
	   PS C:\Users\Raju Kumar\OneDrive\Desktop\Go Language\3_loop> go run loop.go
	   #####################################################################################################################################################################################################################.......................,, press ctrl+C for exit the the Infinite loop
	*/

	//for loop in Go
	/*
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	*/
	//output
	/*
		PS C:\Users\Raju Kumar\OneDrive\Desktop\Go Language\3_loop> go run loop.go
		0
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
	*/

	//Continue and Break In Go lan
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue //here go skip the 5
		}
		if i == 8 {
			break //here loop will stop at 8
		}
		fmt.Println(i)
	}
	//output
	/*
			PS C:\Users\Raju Kumar\OneDrive\Desktop\Go Language\3_loop> go run loop.go
		1
		2
		3
		4
		6
		7

	*/

}
