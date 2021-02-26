package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var words []string
	var shortnessGroup []string
	var numOfInput int

	fmt.Println("Input :")
	fmt.Scanln(&numOfInput)
	
	scanner := bufio.NewScanner(os.Stdin)

	for i :=0 ; i < numOfInput; i++ {
		scanner.Scan()
		word := scanner.Text()
		words = append(words, word)
	}	

	for i := 0; i < len(words); i++ {
		shortness := convertToShortness(words[i])
		shortnessGroup = append(shortnessGroup, shortness)
	}

	sortShortness := shortnessSortByLength(shortnessGroup)
	sortShortness = shortnessSortByChar(sortShortness)
	fmt.Println("Output :")
	for i := 0; i < len(sortShortness); i++ {
		fmt.Printf("%v\n",sortShortness[i])

	}
}

func convertToShortness(str string) string {
	word := strings.Split(str, " ") 
	var char []string
	for i := 0; i < len(word) ; i++ {
		rune := []rune(word[i])
		if (unicode.IsUpper(rune[0])) {
			char = append(char, string(rune[0]))
		}
	}
	shortness := strings.Join(char,"")	
	return shortness
}

func shortnessSortByLength(array []string) []string{
	for i := 1; i < len(array); i++ {
		numCurrent := len(array[i])
		valCurrent := array[i]
		j := i - 1
			for j >= 0 && len(array[j]) < numCurrent {
				array[j+1] = array[j]
				j--
			}
		array[j+1] = valCurrent;
	}
	return array
}

func shortnessSortByChar(array []string) []string {
	for i := 1; i < len(array); i++ {
		numCurrent := len(array[i]) 
		valCurrent := array[i] 
		j := i - 1 
		k := 0
		for j >= 0 && len(array[j]) == numCurrent {
			if array[j][k] > valCurrent[k] {
				array[j+1] = array[j]
				j--
				k = 0	
			} else if array[j][k] == valCurrent[k] && k < numCurrent {
				k++
			} else {
				break
			}
		}
		array[j+1] = valCurrent;
	}
	return array
}


