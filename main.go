package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for _, element := range bulkOfTests() {
		fmt.Println(processList(element.start, element.end))
	}
}

func processList(start string, end string) []TimeStruct {
	var result []TimeStruct
	var lastTime int32
	var lastItem int32
	for _, startItem := range strToMinutes(start) {
		var startValue int32
		var endValue int32
		for _, endItem := range strToMinutes(end) {
			if startItem.start == endItem.start && startItem.end == endItem.end {
				startValue = 0
				endValue = 0
			} else {
				if endItem.start > startItem.start && endItem.start < startItem.end {
					if startItem.start > startValue || startValue == 0 {
						if startItem.start > lastTime {
							startValue = startItem.start
						}
					}
					if endItem.start < endValue || endValue == 0 {
						endValue = endItem.start
					}
				} else {
					if endItem.end < startItem.end {
						if endItem.end > lastItem {
							startValue = endItem.end
						}
						endValue = startItem.end
					} else {
						if endItem.start < startItem.start && endItem.end > startItem.end {
							endValue = lastTime
						} else {
							if startValue == 0 {
								startValue = startItem.start
							}
							if endValue == 0 {
								endValue = startItem.end
							}
						}
					}
				}
			}
		}
		lastTime = startItem.end
		lastItem = endValue
		Tmp := append(result, TimeStruct{minutesToHour(startValue), minutesToHour(endValue)})
		result = Tmp
	}
	return result
}

func bulkOfTests() []TimeStruct {
	var tests []TimeStruct
	result := append(tests,
		TimeStruct{"9:00-10:00", "9:00-9:30"},
		TimeStruct{"9:00-10:00", "9:00-10:00"},
		TimeStruct{"9:00-9:30", "9:30-15:00"},
		TimeStruct{"9:00-9:30, 10:00-10:30", "9:15-10:15"},
		TimeStruct{"9:00-11:00, 13:00-15:00", "9:00-9:15, 10:00-10:15, 12:30-16:00"})
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

func minutesToHour(value int32) string {
	hours, minutes := divmod(value, 60)
	return strconv.Itoa(int(hours)) + ":" + padNumberWithZero(minutes)
}

func padNumberWithZero(value int32) string {
	return fmt.Sprintf("%02d", value)
}

func divmod(numerator, denominator int32) (quotient, remainder int32) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

//TimeRange is a struct of times
type TimeRange struct {
	start int32
	end   int32
}

//TimeStruct is a struct of tests
type TimeStruct struct {
	start string
	end   string
}
