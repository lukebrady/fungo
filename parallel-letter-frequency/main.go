package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func findLetterCount(letter rune, word string, count chan int) {
	var lcount int
	for _, l := range word {
		if letter == l {
			lcount++
		}
	}
	count <- lcount
}

func main() {
	if len(os.Args) < 2 {
		panic(nil)
	}
	page, err := http.Get("https://" + os.Args[1])
	if err != nil {
		panic(err)
	}
	defer page.Body.Close()
	test, err := ioutil.ReadAll(page.Body)

	alpha := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
		'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x', 'y', 'z'}
	letterCount := map[rune]int{}
	countChan := make(chan int)
	if len(os.Args) < 2 {
		panic(nil)
	} else {
		for _, l := range alpha {
			go findLetterCount(l, string(test), countChan)
			letterCount[l] = <-countChan
			fmt.Printf("%c : %d\n", l, letterCount[l])
		}
	}
}
