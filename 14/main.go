package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i interface{} = make(chan int)

	fmt.Println(reflect.TypeOf(i))

	fmt.Printf("Type = %T\n", i)

	var i1 interface{} = true

	definitionType(i1)

}

// Пустой интерфейс может содержать значения любого типа
func definitionType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	case chan int:
		fmt.Println("channel of int")
	case chan string:
		fmt.Println("channel of string")
	default:
		fmt.Println("unknown type")
	}
}
