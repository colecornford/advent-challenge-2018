package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	fmt.Println(len(react(getData())))
}

func puzzle2() {
	fmt.Println(len(react(killUnit(getData(), getShortestPolymer()))))
}

func react(polymer string) string {
	units := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for {
		oldPolymer := polymer
		for _, x := range units {
			var replacer = strings.NewReplacer(x+strings.ToUpper(x), "", strings.ToUpper(x)+x, "")
			polymer = replacer.Replace(polymer)
		}
		if oldPolymer == polymer {
			return polymer
		}
	}
}

func killUnit(polymer string, unit string) string {
	var replacer = strings.NewReplacer(unit, "", strings.ToUpper(unit), "")
	return replacer.Replace(polymer)
}

func getShortestPolymer() string {
	units := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	scores := make(map[int]string)
	polymer := getData()
	count := 10000000
	bestLetter := ""

	for _, x := range units {
		evisceratedPolymer := killUnit(polymer, x)
		scores[len(react(evisceratedPolymer))] = x
	}
	for k, v := range scores {
		if k < count {
			bestLetter = v
			count = k

		}
	}
	return bestLetter
}

func getData() (lines string) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)
}
