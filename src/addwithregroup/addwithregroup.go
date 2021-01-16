package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Numbers struct {
	first_ones int
	first_tens int
	second_ones int
	second_tens int
}

var numbers = [...]int{0,1,2,3,4,5,6,7,8,9}

func getInitialPlaceValue() int {
	max := len(numbers)
	min := len(numbers)/2
	rand.Seed(time.Now().UTC().UnixNano())
	return numbers[min + rand.Intn(max-min)]
}

func getNextPlaceValue(input int) int {
        max := len(numbers) - input
        if max == 0 {
                return 0
        }
        min := 0
        rand.Seed(time.Now().UTC().UnixNano())
        return numbers[min + rand.Intn(max-min)]
}

func checkAddnoRegroupAnswer(num Numbers,ans int) bool {
	if ((num.first_tens + num.second_tens)*10 + num.second_ones + num.first_ones) == ans {
		return true
	} else {
		return false
	}
}

func main() {
	var ans,count int
	var num Numbers
	for {
		fmt.Println("Solve the following problem.")
		fmt.Println("****************************")
		num.first_ones = getInitialPlaceValue()
		num.second_ones = getInitialPlaceValue()
		num.first_tens = getInitialPlaceValue()
		num.second_tens = getNextPlaceValue(num.first_tens)
		fmt.Printf(" %d%d\n+%d%d\n---\n ",num.first_tens,num.first_ones,num.second_tens,num.second_ones)
		fmt.Scanf("%d",&ans)
		if checkAddnoRegroupAnswer(num,ans) {
			fmt.Println("*****Great answer!! Congratulations, lets try another now.*****")
		} else {
                        fmt.Println("Sorry!! Wrong Answer :-( , lets try again.")                       
                        count++
                        for count < 3 {
                                fmt.Print("Response : ")
                                fmt.Scanf("%d",&ans)
                                if checkAddnoRegroupAnswer(num,ans) {
                                        fmt.Println("Great answer!! Congratulations, lets try another now.")            
                                        break
                                } else {
                                        count++
                                        if count == 3 {
                                                fmt.Println("Sorry!! Wrong Answer :-( , you have exhausted your 3 attempts, lets try another question.")
                                                break
                                        } else {
                                                fmt.Println("Sorry!! Wrong Answer :-( , lets try again.")                       
                                                continue
                                        }
                                }
                        }
		}
	}
}
