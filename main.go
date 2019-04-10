package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	processList(bulkOfTests()[4].start, bulkOfTests()[4].end)
}

func processList(start string, end string) {
	fmt.Println(strToMinutes(start))
	fmt.Println(strToMinutes(end))
}

func bulkOfTests() []TestStruct {
	var tests []TestStruct
	result := append(tests,
		TestStruct{"9:00-10:00", "9:00-9:30"},
		TestStruct{"9:00-10:00", "9:00-10:00"},
		TestStruct{"9:00-9:30", "9:30-15:00"},
		TestStruct{"9:00-9:30, 10:00-10:30", "9:15-10:15"},
		TestStruct{"9:00-11:00, 13:00-15:00", "9:00-9:15, 10:00-10:15, 12:30-16:00"})
	return result
}

func strToMinutes(values string) []TimeRange {
	listRange := strings.Split(values, ",")
	var arrRange []TimeRange
	for _, element := range listRange {
		startAndEnd := strings.Split(element, "-")
		startTime := strings.Split(strings.Trim(startAndEnd[0], " "), ":")
		endTime := strings.Split(strings.Trim(startAndEnd[1], " "), ":")

		firstStartPart, _ := strconv.ParseInt(startTime[0], 10, 32)
		secondStartPart, _ := strconv.ParseInt(startTime[1], 10, 32)
		start := (firstStartPart * 60) + secondStartPart

		firstEndPart, _ := strconv.ParseInt(endTime[0], 10, 32)
		secondEndPart, _ := strconv.ParseInt(endTime[1], 10, 32)
		end := (firstEndPart * 60) + secondEndPart

		tmp := TimeRange{int32(start), int32(end)}
		newArrRange := append(arrRange, tmp)
		arrRange = newArrRange
	}
	return arrRange
}

func minutesToHour(value int32) {

}

//TimeRange is a struct of times
type TimeRange struct {
	start int32
	end   int32
}

//TestStruct is a struct of tests
type TestStruct struct {
	start string
	end   string
}
