package main

import "fmt"

func main() {
	students := make(map[string]int)

	students["Alice"] = 20
	students["Bob"] = 22
	students["Charlie"] = 19

	fmt.Println("Students and their ages:", students)

	fmt.Println("Alice's age:", students["Alice"])

	students["Bob"] = 23
	fmt.Println("Updated Students and their ages:", students)

	delete(students, "Charlie")
	fmt.Println("After deleting Charlie:", students)

}
