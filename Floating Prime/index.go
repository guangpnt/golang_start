package main

import (
	"fmt"
	"strconv"
)

func main() {
    var s string
    var f float64
	var array [3]string
   for {
        _, err := fmt.Scan(&s)
        f, err = strconv.ParseFloat(s, 8)
        if err != nil {
            fmt.Println("Enter a valid number")
        } else if len(s) > 13 {
			fmt.Println("Number is more than 12 digits")
		} else if f > 0 && f <= 10 {
			var prime bool
            s := fmt.Sprintf("%f", f)
			char0 := string(s[0])
			char1 := string(s[1])
			char2 := string(s[2])
			char3 := string(s[3])
			char4 := string(s[4]) 
			if char1 == "0" {
				array[0] = char0 + char1
			} else {
				array[0] = char0 + char2
			}
			array[1] = array[0] + char3
			array[2] = array[1] + char4
			for i := 0 ; i < 3 ; i++ {
				num, _ := strconv.Atoi(array[i])
				for k := 2 ; k < num - 1 ; k++ {
					if num % k == 0 {
						prime = false
						break
					} else {
						prime = true
					}
				}
				if prime {
					break
				}
			}
			if prime {
				fmt.Println("TRUE")
			} else {
				fmt.Println("FALSE")
			}
        } else if f == 0 {
			break
		} else {
			fmt.Println("Enter a valid number")
		}
    }
}