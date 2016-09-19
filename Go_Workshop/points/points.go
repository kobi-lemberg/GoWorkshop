package main

import (
	//"encoding/csv"
	"fmt"
	//"io"
	//"log"
	//"strconv"
	//"os"
)

//const (
//	space   = ' '
//	newline = '\n'
//	tab     = '\t'
//)

type Point struct {
	X float64
	Y float64
}

var (
	names = []string{"A","B","C","D","F"}
	cutoffs = []int{90,80,70,60,0}
	grades  map[int]string

)

func init() {
/*	grades = make(map[int]string)
	high := 100
	for i , name := range names {
		for score := high; score >= cutoffs[i]; score-- {
		grades[score] = name
		}
		high = cutoffs[i] - 1
	}*/
}


func NewPoint(x,y float64) *Point {
	return &Point{x,y}
}

func main() {

	var p Point
	fmt.Printf("%v %T\n" ,p ,p )
	fmt.Printf("%+v %T\n" ,p ,p )
	fmt.Printf("%#v %T\n" ,p ,p )

	p.Y = 17
	fmt.Printf("%+v %T\n" ,p ,p )

	p1 := Point{1,2}
	fmt.Printf("%+v %T\n" ,p1 ,p1 )

	p2 := Point{Y:10,X:20}
	fmt.Printf("%+v %T\n" ,p2 ,p2 )

	var pp *Point
	pp = &p2
	fmt.Printf("%#v\n",pp)

	pp2 := new(Point)
	fmt.Printf("%#v\n",pp2)

	pp3 := &Point{1,2}
	fmt.Printf("%#v\n",pp3)

	pp4 := NewPoint(7,8)
	fmt.Printf("%#v\n",pp4)


}


/*func wc(file io.Reader) (int, int, int, error) {
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

}*/
