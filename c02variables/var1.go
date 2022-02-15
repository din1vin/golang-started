package main

import "fmt"

func main() {
	/**
	way1 var + type without initValue
	Golang type init zero value
	*/
	var number int64
	fmt.Println("number = ", number)

	/**
	way2 var name = $value
	Golang predicated type on value
	*/
	var name = "name"
	fmt.Println("name = ", name)
	fmt.Printf("type of name : %T \n", name)

	var name1 string = "name"
	fmt.Println("name1 = ", name1)
}
