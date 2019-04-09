package main

import "fmt"

func main() {
	var listOne string
	var listTwo string
	fmt.Scanf("%s", &listOne)
	fmt.Scanf("%s", &listTwo)
	fmt.Println("Lista: ", subtract(listOne, listTwo))
}

func subtract(listOne []string, listTwo []string) string {
	return listOne + listTwo
}
