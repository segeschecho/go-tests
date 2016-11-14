package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}

		fmt.Println("Se ejecuta siempre")
	}()
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in F !!", r)
		}

		fmt.Println("Se ejecuta siempre 2")
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		var a, b int

		// se va a hacer una division por cero !
		a = 0
		b = 3

		fmt.Sprintf(string(b/a))
		//panic(fmt.Sprintf("%v", i))
	}
	fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
