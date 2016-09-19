package main

import (
	"flag"
	"fmt"
	"io"
	"log"
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

type wcWriter struct {
	lines int
	words int
	bytes int
	inWord bool
}


func (wc *wcWriter) Write(p []byte) (int, error){
	for _,b := range p{
		wc.bytes++
		switch b {
		case newline:
			wc.lines++
			wc.inWord = false
		case tab, space:
			wc.inWord = false
		default:
			if !wc.inWord {
				wc.inWord = true
				wc.words++
			}
		}
	}
	return len(p),nil


}

func main() {
	flag.Parse()
	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	var file io.Reader
	var fileName string
	if flag.NArg() == 0 {
		file = os.Stdin
		fileName = "<stdin>"
	} else {
		fileName = flag.Arg(0)
		if f, err := os.Open(fileName); err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			file = f
		}
	}

	lines, words, bytes, err := wc(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %d %d %s\n", lines, words, bytes, fileName)
}

func wc(file io.Reader) (int, int, int, error) {
	wcr :=wcWriter{}
	_ ,err := io.Copy(&wcr,file)
	return wcr.lines , wcr.words , wcr.bytes, err
}