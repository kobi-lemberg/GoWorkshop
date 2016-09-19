package main

import (
	"fmt"
	"go/types"
)

func main() {
	var i interface{}
	i= 7
	fmt.Printf("%v - %T\n",i,i)

	i= "hello"
	fmt.Printf("%v - %T\n",i,i)

	s:= i.(string)
	fmt.Printf("%v - %T\n",s,s)

	//This will panic
/*	s:= i.(int)
	fmt.Printf("%v - %T\n",s,s)*/

	j , ok := i.(int)
	if ok {
		fmt.Printf("int = %d",j)
	} else {
		fmt.Printf("Not int!")
	}

	switch i.(type) {
	case int:
		fmt.Println("int!")
	case string:
		fmt.Println("string!")
	default:
		fmt.Printf("Unknown type - %T\n",s)
	}
}

