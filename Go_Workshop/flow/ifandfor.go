package main //Must heve that name!!

import "fmt"
import "os"

func main() {

	a :=1
	if a>0 {

		fmt.Println("a>0")//Each exported method starts with capital letter
		b:=10
		fmt.Println("b ",b)

	}else {
		fmt.Println("a<0")
		b:=20
		fmt.Println("b ",b)
	}
	b:=10
	fmt.Println("b ",b);

	if a>0 && b==0 {

		fmt.Println("a>0")//Each exported method starts with capital letter
		b:=10
		fmt.Println("b ",b)
	}

	if file,error := os.Open("test.test"); error!=nil{
		fmt.Println("Cant open!!")
	} else {
		fmt.Println("File is opened ",file.Name())
	}

	for i:=0; i<=3; i++ {
		fmt.Println(i)
	}

	i :=0
	for i<3 { //While
		fmt.Println(i)
		i++
	}
/*
	for { //While
		if i<3{
			break
		}
		fmt.Println(i)
		i++
	}*/

	s := "hello"
	//range with single value
	for c:= range s{
		fmt.Println(c)
	}

	//range with double value -> index,item
	//works on runes
	for i,c:= range s{
		fmt.Println(i,c,string(c))
	}


	//For each
	for _,c:= range s{
		fmt.Println(i,c)
	}

	counter :=0
	//i=0
	for i:=0; i<1000;i++ { //
		if i%3 ==0 || i%5 ==0{
			counter+=i
		}
	}
	fmt.Println("counter ",counter)



}
