package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	string1 := []byte("hello\nworld\n")
	err := ioutil.WriteFile("result.txt", string1, 0644)
	check(err)

}
