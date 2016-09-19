package main //Must heve that name!!

import "fmt"


func main() {
	var i int
	fmt.Println(i)
	j :=17
	fmt.Printf("%d %T\n",j,j)

	f :=0.314
	fmt.Printf("%f %T\n", f, f)

	a := float64(j) + f
	fmt.Printf("%f %T\n", a, a)

	s := "This is a string"
	fmt.Printf("%s %T\n", s, s)

	x := 1
	y := "1"
	fmt.Printf("x=%d y=%q\n", x, y)
	fmt.Printf("x=%v y=%v\n", x, y)
	hi := "שלום"
	fmt.Printf("len(hi) %d", len(hi))





}
