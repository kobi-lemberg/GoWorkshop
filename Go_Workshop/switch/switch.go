package main //Must heve that name!!

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().Unix())
	loons := []string{"bugs","tweety","elmer"}
	winner := loons[rand.Intn(len(loons))]

	switch winner {

	case "bugs":
		fmt.Println("Bugs")
	case "tweety":
		fmt.Println("tweety")
	case "elmer":
		fmt.Println("elmer")
/*	default:
		fmt.Println("")*/

	}

	switch  {
	case winner=="bugs":
		fmt.Println("Bugs")
	case 10>2:
		fmt.Println("10>2")
	}
	/*
	defer  fmt.Println(1)
	defer  fmt.Println(2)
	defer  fmt.Println(3)
	fmt.Println("in main")
	defer  fmt.Println(4)
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
	fmt.Printf("len(hi) %d", len(hi))*/





}
