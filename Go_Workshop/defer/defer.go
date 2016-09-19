package main //Must heve that name!!

import "fmt"


func main() {
	defer  fmt.Println(1)
	defer  fmt.Println(2)
	defer  fmt.Println(3)
	fmt.Println("in main")
	defer  fmt.Println(4)

}
