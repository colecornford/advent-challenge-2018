package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"os"
	"time"
)


func main() {
	puzzle1()
	puzzle2()
	os.Exit(0)

}

type Guard struct {
	id int
	timeAsleep int
	daysAsleep int
	minuteSleptMost int
	stamps []Stamp
	minutesAsleep map[int]int
}

type Stamp struct {
	timestamp time.Time
	status string
}

func puzzle1(){
	guards := setupPuzzle()
	printGuard(longestSleeper(guards))
}

func puzzle2(){
	guards := setupPuzzle()
	printGuard(clockNapper(guards))
}

func setupPuzzle()([]Guard){
	input := getData() //  file IO
	stamps := createStamps(input) // turn input into Stamp structs.
	guards := addStampsToGuards(stamps) // Give each Guard all their Stamps
	
	for x, a := range guards{
		guards[x].daysAsleep = getSleepDays(a) // Do maths
		guards[x].timeAsleep = getTotalSleepTime(a)
		guards[x].minutesAsleep = getMinutesAsleep(a)
	}
	return guards
}

func addStampsToGuards(stamps []Stamp)([]Guard){
	// This function assigns all the timestamps a guard has in the IO file to him.
	// It identifies a Guard, Then adds all further timestamps to a "shift".
	// When we encounter a new Guard, we append all the timestamps from a "shift" to the guard
	// loop until all shifts have been accounted for.
	shiftStamps := make([]Stamp,0)
	guard := Guard{}
	guards := make([]Guard,0)
	for _, a := range stamps{
		if changingOfTheGuard(a){
			guard.stamps = append(guard.stamps,shiftStamps...)
			guards = updateSleepDiary(guards, guard)
			guard = Guard{}
			shiftStamps = make([]Stamp,0)
			guard.id = getGuardId(a.status) 
		}else{
			shiftStamps = append(shiftStamps,a)
		}
	}
	return guards
}

func changingOfTheGuard(s Stamp)(bool){
	return !strings.HasPrefix(s.status,"f") && !strings.HasPrefix(s.status,"w") // New Guard
}

func longestSleeper(guards []Guard)(Guard){ // Wincon Puzzle1
	maxScore, guardId:= 0,0
	for x, g := range guards{
		if 	g.timeAsleep > maxScore{
			maxScore = g.timeAsleep
			guardId = x
		}
		x++
	}
	return guards[guardId]
}

func clockNapper(guards []Guard)(Guard){ // Wincon Puzzle2
	maxScore, guardId:= 0,0
	for x, g := range guards{
		count := 0
		for _, v := range g.minutesAsleep{
			if v > count{
				count = v
			}
		}
		if 	count > maxScore{
			maxScore = count
			guardId = x
		}
		x++
	}
	return guards[guardId]
}

func updateSleepDiary(guards []Guard, newGuard Guard)([]Guard){
	for x, _ := range guards{
		if guards[x].id == newGuard.id{
			guards[x].stamps = append(guards[x].stamps, newGuard.stamps...)
			return guards
		}
		x++
	}
	guards = append(guards,newGuard)
	return guards
}

func getGuardId(str string)(int){
	i, err := strconv.Atoi(strings.Replace(strings.Split(str," ")[1],"#","",-1))
	if err != nil {
		panic(err)
	}
	return i
}

func printGuard(g Guard){
	fmt.Printf("Guard id: %d \n", g.id)
	for _, a := range g.stamps{
		t := a.timestamp
		fmt.Printf("%d-%02d-%02d %02d:%02d ",t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
		fmt.Println(a.status)
	}
	fmt.Printf("sleepTime: %d \n", g.timeAsleep)
	fmt.Printf("sleepDays: %d \n\n", g.daysAsleep)
	fmt.Print(g.minutesAsleep)
	count, minute := 0,0
	for k, v := range g.minutesAsleep{
		if v > count{
			minute = k
			count = v
		}
	}
	fmt.Printf("\nMostCommonMinute: %d \n\n",minute)
	fmt.Printf("Final Score: %d", g.id * minute)
}

func getSleepDays(g Guard)(int){
	prior := time.Time{}
	days := 0
	for _, a := range g.stamps{
		if a.timestamp.Month() > prior.Month(){
			days++
		}else if a.timestamp.Day() > prior.Day(){
			days++
		}
		prior = a.timestamp
	}
	return days
}

func getTotalSleepTime(g Guard)(int){
	sleepingTime := 0
	sleeping := false
	wakeUp := time.Time{}
	goSleep := time.Time{}
	for _, a := range g.stamps{
		if sleeping == false{
			goSleep = a.timestamp
			sleeping = true
		} else{
			wakeUp = a.timestamp
			sleepingTime = sleepingTime + wakeUp.Minute() - goSleep.Minute()
			sleeping = false
		}
	}
	return sleepingTime
}

func getMinutesAsleep(g Guard)(map[int]int){
	sleeping := false
	minutesAsleep := make(map[int]int)
	wakeUp := time.Time{}
	goSleep := time.Time{}
	for _, a := range g.stamps{
		if sleeping == false{
			goSleep = a.timestamp
			sleeping = true
		} else{
			wakeUp = a.timestamp
			x := 0
			for x < 60{
				if x >= goSleep.Minute() && x < wakeUp.Minute(){
					minutesAsleep[x] = minutesAsleep[x] + 1
				}
				x++
			}
			sleeping = false
		}
	}
	return minutesAsleep
}

func getData()(lines []string){
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func createStamps(lines []string)([]Stamp){
	stamps := make([]Stamp,0)
	for _, line := range lines{
		stamp := Stamp{}
		stamp.timestamp = timeParsing(strings.Split(string(line), "]")[0])
		stamp.status = strings.Split(string(line), "] ")[1]
		stamps = append(stamps,stamp)
	}

	return bubbleSort(stamps)
}

func bubbleSort(input []Stamp)([]Stamp) {
    n := len(input) + 1
    swapped := true
    for swapped {
        swapped = false
        for i := 1; i < n-1; i++ {
            if input[i-1].timestamp.After(input[i].timestamp) {
                input[i], input[i-1] = input[i-1], input[i]
                swapped = true
            }
        }
    }
    return input
}

func timeParsing(str string)(time.Time){
	str = strings.Replace(str,"[","", -1)
	str = strings.Replace(str,"]","", -1)
	str = strings.Replace(str," ",",", -1)
	str = strings.Replace(str,":",",", -1)
	str = strings.Replace(str,"-",",", -1)
	strs := strings.Split(string(str), ",")

	times := make([]int,0)
	for _, str := range strs{
		fk, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		times = append(times, fk)
	}
	return time.Date(times[0],time.Month(times[1]),times[2],times[3],times[4],0,0, time.UTC)
}
