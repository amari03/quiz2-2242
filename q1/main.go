//creating and implementing go interfaces
//bookstore scenario
package main

import "fmt"

type Bookstore interface{
	printItems()
}

type Book struct{
	name string
	price float32
}
type Bookmarks struct{
	name string
	price float32
}

func (b Book) printItems(){
	fmt.Printf("Title: %s Price: %.2f \n", b.name, b.price)
}

func (m Bookmarks) printItems(){
	fmt.Printf("Title: %s Price: %.2f \n", m.name, m.price)
}

func main(){

	item1 := Book{
		name: "One of us is lying",
		price: 15.00,
	}

	item2 := Bookmarks{
		name: "floral pattern",
		price: 1.50,
	}

	item1.printItems()
	item2.printItems()
}