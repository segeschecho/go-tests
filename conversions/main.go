package main

import (
	"fmt"
	"strconv"
)

func main(){
	string_convertions()
	int_convertions()
}


func string_convertions(){
	// estas asignaciones son lo mismo
	str := "Hola!"
	var str1 = "Hola!"
	var str2 string = "Hola!"

	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)

	if str == str1 && str1 == str2{
		fmt.Println("Son todas iguales")
	}


	/*************************/
	// Numeros a string
	num_int := 10
	num_str := strconv.Itoa(num_int)
	fmt.Println("Numero a string: " + num_str)

	num_str = "25"
	num_int, _ = strconv.Atoi(num_str)  // Igual a strconv.ParseInt(s, 10, 0)

	fmt.Println("String a numero: ", num_int)


	/*************************/
	//Booleans
	bool_bool, _ := strconv.ParseBool("true")
	if bool_bool{
		bool_str := strconv.FormatBool(bool_bool)
		fmt.Println(bool_str)
	}


	/*************************/
	//Numeros
	f_float, _ := strconv.ParseFloat("3.1415", 64)  // base 32 o 64

	if f_float == 3.1415{
		f_str := strconv.FormatFloat(f_float, 'f', 4, 64)  // float, formato del float, precision, base.
		fmt.Println("Float a string:", f_str)
	}

	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println("int a string:", i)

	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println("uint a string:", u)


	// Quotes y conversiones de caracteres no ascii a ascii
	s := "Hello, 世界"
	fmt.Println("Original: ", s)

	q := strconv.QuoteToASCII(s)
	fmt.Println("Escapando no ascii: ", q)

	q = strconv.Quote(s)
	fmt.Println("Poniendo Quotes: ", q)

}


func int_convertions(){
	var i int = 42  // tambien puede ser i := 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println("int: ", i)
	fmt.Println("float64: ", f)
	fmt.Println("uint: ", u)
}