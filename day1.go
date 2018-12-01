package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	data := getData()
	freq := 0
	mapp := make(map[string]int)
	mapp["0"] = 1

	i := 0
	for i <= len(data) {
		adjust, err := strconv.Atoi(data[i])
		if err != nil {
			panic(err)
		}
		freq = freq + adjust
		mapp[strconv.Itoa(freq)] += 1
		if mapp[strconv.Itoa(freq)] == 2{
			fmt.Println("Winner! " + strconv.Itoa(freq))
			os.Exit(0)
		}
		i++
		if i == len(data){
			i = 0
		}
	}
}

func getData()(lines []string){
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}
