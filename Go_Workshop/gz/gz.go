package main

import (
	"io"
	"log"
	"os"
	"compress/gzip"
)


func main() {
	in ,err := os.Open("innocents-abroad.txt")
	if err != nil {
		log.Fatalf("cant open - %s", err)
	}
	defer in.Close()

	out ,err := os.Create("innocents-abroad.txt.gz")
	if err != nil {
		log.Fatalf("cant open - %s", err)
	}
	defer out.Close()

	gz := gzip.NewWriter(out)
	defer gz.Close()

	_ , err = io.Copy(gz,in)
	if err != nil {
		log.Fatalf("Can't compress - %s",err)
	}
}

