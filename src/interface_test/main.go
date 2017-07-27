package main

import (
	"regexp"
	"fmt"
)

func main() {
	/*fmt.Println("interface test")

	f, _ := os.Open("/tmp/amt3.log")

	var rd io.Reader
	rd = f

	buf := make([]byte, 1024)
	n, _ := rd.Read(buf)

	fmt.Println("readlen## ", n)*/

	mobile:="13198375508"

	var mobilePattern = regexp.MustCompile("^((\\+86)|(86))?(1(([35][0-9])|[8][0-9]|[7][0123456789]|[4][579]))\\d{8}$")
	if !mobilePattern.Match([]byte(mobile)) {
		fmt.Println("not match")
	}else{
		fmt.Println("match...")
	}

}
