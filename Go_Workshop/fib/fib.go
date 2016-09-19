package main //Must heve that name!!


import (
	"flag"
	"os"
	"fmt"
	"strconv"
)

func fib(n int) int {
	a, b := 1,1
	for ; n>1; n--{
		a,b = b,a+b
	}
	return a
}

func main() {
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr, "Usage: %s NUM\n", os.Args[0])
		flag.PrintDefaults();
	}

	var verbose bool
	flag.BoolVar(&verbose,"verbose",false,"emit more noise")
	flag.Parse()
	if flag.NArg()!=1{
		fmt.Fprintf(os.Stderr, "Usage: %s NUM\n , Wron var numbers")
		os.Exit(1)
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil{
		fmt.Fprintf(os.Stderr, "Atoi error")
		os.Exit(1)
	}

	if verbose{
		fmt.Printf("Starting cal");
	}

	fmt.Println("Got %v of type %T\n", n, n)
	fmt.Println("Fib(n):", fib(n))

	for i:=1; i<=n ; i++{
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}


	/*
	fmt.Printf("fib(10) %d",fib(10))

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
	*/





}
