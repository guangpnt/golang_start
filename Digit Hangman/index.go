package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	var num int
	var nums []int
	var output = []string{"_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_"}
	var outputs []string
	var incorrectNums []string
	score := 0
	var prevScore int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) != 23 {
			fmt.Println("Try again")
		} else {
			var pass bool
			for i := 1 ; i < 23 ; i = i+2 {
				if s[i] != 32 || s[i-1] < 48 || s[i-1] > 57 || s[22] < 48 || s[22] > 57 {
					pass = false 
					break
				} else {
					pass = true
				} 
			}
			if pass {
				for i :=0 ; i < 5; i++ {
					fmt.Scanln(&num)
					nums = append(nums, num)
				}
				for i := 0 ; i < 5 ; i++ {
					prevScore = score
					for j := 0 ; j < 23 ; j = j+2 {
						ati, _ := strconv.Atoi(string(s[j]))
						if ati == nums[i] {
							output[j/2] = strconv.Itoa(nums[i])
							score++
						}
						if j == 22 {			
							if prevScore == score {
								incorrectNums =  append(incorrectNums,strconv.Itoa(nums[i]))
								s2 := strings.Join(incorrectNums," ")
								s1 := strings.Join(output," ")
								s1 = s1 + " " + s2
								outputs = append(outputs, s1)
							} else {
								s2 := strings.Join(incorrectNums," ")
								s1 := strings.Join(output," ")
								s1 = s1 + " " + s2
								outputs = append(outputs, s1)
							}
						}
					}
				}
				fmt.Println("_ _ _ _ _ _ _ _ _ _ _ _")
				for i := 0 ; i < 5 ; i++ {
					fmt.Println(outputs[i])
				}
				fmt.Println(score)
				break
			} else {
				fmt.Println("Try again")
			}
		}
	}
}