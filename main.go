package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var re []byte = []byte("Exception")

func main() {
	start := time.Now()
	log.Println(start)
	// f := "drivertest.log"
	// Read(f)
	getFileList(path)
	for s, f := range Files {
		// println(f)
		println("=======", s, "=======")
		ReadV1(f)
	}
	// ReadV1(f)
	d := time.Since(start)
	log.Println(d)
}

// Read ...
func Read(f string) {
	dat, err := ioutil.ReadFile(f)
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadV1 ...
func ReadV1(f string) {
	fp, err := os.Open(f)
	if err != nil {
		fmt.Println(f, ":", err)
		return
	}
	defer fp.Close()
	r := bufio.NewReader(fp)
	i := 0
	for {
		i++
		b, _, err := r.ReadLine()
		// log.Println(err)
		if err == nil {
			if bytes.Contains(b, re) {

				fmt.Printf("%s: %d-- %s\n", f, i, string(b))
			}

		} else if io.EOF == err {
			return
		} else {
			log.Println(err)
		}
	}
}
