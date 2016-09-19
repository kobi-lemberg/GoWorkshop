package main

import (
	"flag"
	"fmt"
	"io"
	//"log"
	"os"
)

const (
	space   = ' '
	newline = '\n'
	tab     = '\t'
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: wc [FILE]\n")
		flag.PrintDefaults()
	}
}

func main() {
	var ages map[string]int
	fmt.Printf("%v (len=%d) %T\n", ages, len(ages), ages)
	age := ages["daffy"]
	fmt.Println(age)


	//ages["daffy"] = 80 Pannic on assignment to nill map
	ages = make(map[string]int)
	ages["daffy"] = 80
	age = ages["daffy"]
	fmt.Println(age)

	if age, ok := ages["bugs"]; !ok {
		fmt.Println("bugs not found")
	}else {
		fmt.Println("bugs age %d", age)
	}

	colors := map[string]string{
		"bugs": "gray",
		"bugs2": "gray2",
		"bugs3": "gray3"}

	for name := range colors {
		fmt.Println(name)
	}

	for name,colors := range colors {
		fmt.Println("%s %s ",name, colors)
	}

	/*var arr [10]int
	fmt.Printf("%v - %T\n", arr, arr)

	s :=arr[1:3]
	fmt.Printf("%v len=%d cap=%d type=%T\n", s, len(s),cap(s),s)

	arr[2] = 17
	fmt.Println(s)

	var a []int //Slice
	fmt.Println(a,len(a))

	b := make([]int ,3 )
	fmt.Printf("%v len=%d cap=%d type=%T\n", b, len(b),cap(b),b)

	var d []int
	for i:=0; i<5; i++{
		d = append(d,i)
		fmt.Printf("%v len=%d cap=%d type=%T\n", d, len(d),cap(d),d)
	}

	cs := []string{"a","b","c"}
	for i,val := range cs{
		fmt.Printf("%s at %d\n",val,i)
	}*/

	//flag.Parse()
	//if flag.NArg() > 1 {
	//	fmt.Fprintf(os.Stderr, "error: wrong number of arguments")
	//	os.Exit(1)
	//}
	//
	//var file io.Reader
	//var fileName string
	//if flag.NArg() == 0 {
	//	file = os.Stdin
	//	fileName = "<stdin>"
	//} else {
	//	fileName = flag.Arg(0)
	//	if f, err := os.Open(fileName); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		defer f.Close()
	//		file = f
	//	}
	//}
	//
	//lines, words, bytes, err := wc(file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d %d %d %s\n", lines, words, bytes, fileName)
}

func wc(file io.Reader) (int, int, int, error) {
	var lines, words, bytes int
	buf := make([]byte, 1024) // allocate 1k buffer
	inWord := false
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, 0, 0, err
			}
		}
		for i := 0; i < n; i++ {
			bytes++
			switch buf[i] {
			case newline:
				lines++
				inWord = false
			case tab, space:
				inWord = false
			default:
				if !inWord {
					inWord = true
					words++
				}

			}
		}
	}
	return lines, words, bytes, nil

}
