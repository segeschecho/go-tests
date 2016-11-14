package main

type Fruit struct {
	name string
	amount int
}


type Stock struct {
	elements []Fruit
}

func (s Stock) sum() int{
	var total int = 0

	for _, elem := range s.elements{
		total = total + elem.amount
	}

	return total
}




func main(){
	f1 := Fruit{name: "pera", amount:5}
	f2 := Fruit{name: "banana", amount: 2}
	f3 := Fruit{name: "Kiwi", amount: 1}

	array := []Fruit{f1,f2,f3}

	stock := Stock{elements: array}

	println(stock.sum())
}