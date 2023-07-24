package main

import "fmt"

func main() {
	name := []string{"bill"}
	name1 := "bill"

	namePointer := &name
	namePointer1 := &name1

	fmt.Println(&(namePointer))
	fmt.Printf("%v"+"%v", &name, &namePointer)
	printPointer(&name)

	fmt.Println(&namePointer1)
	printPointer1(namePointer1)
}

func printPointer(namePointer *[]string) {
	fmt.Println(&namePointer)
}

// package main

// import "fmt"

// func main() {
// 	name := "bill"

// 	namePointer := &name

// 	fmt.Println(&namePointer)
// 	printPointer(namePointer)
// }

func printPointer1(namePointer *string) {
	fmt.Println(namePointer)
}
