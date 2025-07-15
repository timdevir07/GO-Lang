package main

import "fmt"

func main() {
	const player = "Virat Kohli"
	fmt.Println("Name of the " + player)
	const (
		name = "Virat"
		team = "RCB"
		role = "Batter"
		run  = 10000
	)
	const (
		p_name = "Raju kumar"
		p_team = "RCB"
		p_role = "Allrounder"
		p_run  = 5000
	)
	fmt.Println(name)
	fmt.Println(team)
	fmt.Println(role)
	fmt.Println(run)
	fmt.Println(" ")
	fmt.Println("Other Player Record")
	fmt.Println(" ")
	fmt.Println("Player Name" + p_name)
	fmt.Println(p_team)
	fmt.Println(p_role)
	fmt.Println(p_run)
}
