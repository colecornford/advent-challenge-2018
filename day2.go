package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	puzzle1()
	puzzle2()
	os.Exit(0)
}

func getData()(lines []string){
	content, err := ioutil.ReadFile("puzzleinput.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func puzzle1(){
	testData := getData()
	mymap := make(map[string]int)
	doubleBoxes := 0
	tripleBoxes := 0
	i := 0
	for i < len(testData){
		for _, char := range testData[i]{
			mymap[string(char)] += 1
		}
		doubleBoxes += countBoxes(mymap, 2)
		tripleBoxes += countBoxes(mymap, 3)
		mymap = make(map[string]int)
		i++
	}
	fmt.Println(strconv.Itoa(doubleBoxes * tripleBoxes))
}

func puzzle2(){
	testData := getData()
	i, j, x := 0, 0, 0
	lotsaBoxes := make([]boxCompare, 1)
	for i < len(testData){
		for j < len(testData){
			lotsaBoxes = append(lotsaBoxes, addBox(testData[i], testData[j]))
			j++
		}
		j = 0
	i++
	}
	for x < len(lotsaBoxes){
		if lotsaBoxes[x].distance == 1{
			fmt.Println(condemnHeretics(lotsaBoxes[x].b1, lotsaBoxes[x].b2))
			break
		}
		x++
	}

}

func addBox(str1, str2 string) boxCompare{
	s1 := []rune(str1)
	s2 := []rune(str2)

	dist := 0
	a := 0
	for a < len(str2){
		if s1[a] != s2[a]{
			dist += 1
		}
		a++
	}
	return boxCompare{str1, str2, dist}
}

func condemnHeretics(str1, str2 string) string{
	s1 := []rune(str1)
	s2 := []rune(str2)

	a := 0
	for a < len(str2){
		if s1[a] != s2[a]{
			s1[a] = 0
		}
		a++
	}
	return string(s1)
}

func countBoxes(m map[string]int, whichBox int) int{
	boxes := 0
	for _, v := range m {
		if v == whichBox{
			boxes += 1
			break
		}
	}
	return boxes
}

type boxCompare struct {
	b1 string
	b2 string
	distance int
}
