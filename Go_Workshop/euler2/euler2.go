package main

import "fmt"

//return n'th fibo

var(


)

func fibs() chan int{
	ret := make(chan int)
	go func() {
		a:=1
		b:=2
		for {
			a,b = b,a+b
			ret <- a
		}
	}()
	return ret
}

func upto(in chan int, n int) chan int{
	ret := make(chan int)
	go func() {
		for range in {
			num := <- in //Get next num
			if num<n{
				ret <- num
			} else{
				close(ret)
				return
			}
		}
	}()
	return ret
}

func even(in chan int) chan int{
	//close(in)
	ret := make(chan int)
	go func() {
		for range in {
			num := <- in //Get next num
			if num%2==0{
				ret <- num
			}
		}
	}()

	return ret
}

func main() {
	sum := 0
	for n := range upto(even(fibs()),4000000) {
		sum += n
	}
	fmt.Println(sum)

}
