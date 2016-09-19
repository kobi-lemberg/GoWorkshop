package main //Must heve that name!!


import (
	"fmt"
	"time"
)

func Print(i int){
	time.Sleep(100 * time.Microsecond)
	fmt.Println(i)
}

func main() {
	//fmt.Println("Hello, World!")

	//Thsi is print 5 5 5 5 5 - > ==
	/*for i:=0; i<5 ; i++ {
		go func (){
			fmt.Println(i)
		}()
	}*/

	for i:=0; i<5 ; i++ {
		go func (n int){
			fmt.Println(n)
		}(i)
	}

	ch := make(chan int) //unbuffered
	go func() {
		ch <- 7

	}()

	i := <-ch
	fmt.Printf("Got %d\n",i)


	bch := make(chan int,10)
	bch <- 7

	j := <-bch
	fmt.Printf("Got %d\n",j)





	time.Sleep(500 * time.Microsecond)

}
