package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args, " "))
}

func main() {
	t := time.Now()
	echo1()
	println(time.Since(t))
	t = time.Now()
	echo2()
	println(time.Since(t))
	t = time.Now()
	echo3()
	println(time.Since(t))
}
