package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"os"
)

const fabricSize = 1100 // Change as you get bigger samples.

func main() {

	puzzle1()
	puzzle2()
	os.Exit(0)

}

func puzzle1(){
	claims := parseData()
	fabric := populateGrid(claims) 
	//printGrid(fabric, fabricSize, fabricSize) // Uncomment for visual representation of grid
	fmt.Println(countBad(fabric, fabricSize, fabricSize))
}

func puzzle2(){
	claims := parseData()
	fabric := populateGrid(claims)
	//printGrid(fabric, fabricSize, fabricSize) // Uncomment for visual representation of grid
	// If i come across any number greater than 1, the claim is diseased.
	for z, claim := range claims{
		w := claim.w_offset
		for w < (claim.w_offset + claim.w){
			h := claim.h_offset
			for h < (claim.h_offset + claim.h){
				if fabric[w][h] > 1{
					claims[z].diseased = true
				}
				h++
			}
			w++
		}
		z++
	}
	// There should be one healthy claim.
	for _, claim := range claims{
		if claim.diseased == false{
			claim.printClaim()
		}
	}
}


type claim struct {
	id int
	w_offset int
	h_offset int
	w int
	h int
	diseased bool
}

func (c claim) printClaim(){
	fmt.Print(strconv.Itoa(c.id) + ",")
	fmt.Print(strconv.Itoa(c.w_offset) + ",")
	fmt.Print(strconv.Itoa(c.h_offset) + ",")
	fmt.Print(strconv.Itoa(c.w) + ",")
	fmt.Print(strconv.FormatBool(c.diseased) + ",")
	fmt.Print(strconv.Itoa(c.h) + "\n")
}

func printGrid(grid [fabricSize][fabricSize]int, w int, h int){
	x, y := 0, 0
	fmt.Print("\n")
	for y < h{
		for x < w{
			fmt.Print(strconv.Itoa(grid[x][y]))
			x++
		}
		fmt.Print("\n")
		x = 0
		y++
	}
}

func populateGrid(c []claim)([fabricSize][fabricSize]int){
	fabric := [fabricSize][fabricSize]int{}
	for _, claim := range c{
		w := claim.w_offset
		for w < (claim.w_offset + claim.w){
			h := claim.h_offset
			for h < (claim.h_offset + claim.h){
				fabric[w][h] += 1
				h++
			}
			w++
		}
	}
	return fabric
}


func countBad(grid [fabricSize][fabricSize]int, w int, h int)(int) {
	x, y, count := 0, 0, 0
	for y < h{
		for x < w{
			if grid[x][y] >= 2{
				count += 1
			}
			x++
		}
		x = 0
		y++
	}
	return count
}

func getData()(lines []string){
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func parseData()(claims []claim){
	data := getData()
	i := 0

	for i < len(data){
		clean := cleanData(data[i])
		newClaim := claim{clean[0], clean[1], clean[2], clean[3], clean[4], false}
		claims = append(claims, newClaim)
		i++
	}
	return claims
}

func cleanData(dirty string)(clean []int){
	dirty = strings.Replace(dirty," ","", -1)
	dirty = strings.Replace(dirty,"#","", -1)
	dirty = strings.Replace(dirty,"@",",", -1)
	dirty = strings.Replace(dirty,":",",", -1)
	dirty = strings.Replace(dirty,"x",",", -1)
	cleanStr := strings.Split(dirty,",")
	i := 0
	for i < len(cleanStr){
		add, err := strconv.Atoi(cleanStr[i])
		if err != nil{
			panic(err)
		}
		clean = append(clean, add)
		i++
	}
	return clean
}
