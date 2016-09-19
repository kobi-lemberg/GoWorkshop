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