package main

import "fmt"

// Ejemplos de uso de punteros y su comportamiento.

type data struct{
	val int
	val2 string
}

func zero(xPtr *data) {
	(*xPtr).val = 0
}

func hello(s *string){
	*s = "hello"
}

func one(d *data){
	d.val = 1
}

func main() {
	x := data{val:5, val2: "nada"}

	zero(&x)
	fmt.Println(x.val) // x is 0

	y := new(string)  // Devuelve un puntero a string
	hello(y)
	fmt.Println(*y)

	one(&x)
	fmt.Println(x.val)
}
