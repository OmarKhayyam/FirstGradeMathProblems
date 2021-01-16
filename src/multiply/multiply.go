package main

import (
	"fmt"
	"math/rand"
)

type numpair struct {
	numgroup int
	numsize int
}

var problems = make([]numpair,100)

func initialize(probs []numpair) {
	sz := len(probs)
	count := 0
	for outer := 0;outer < 10;outer++ {
		for inner := 0;inner < 10;inner++ {
			probs[count].numgroup = outer + 1
			probs[count].numsize = inner + 1
			count++
			if count == sz {
				break
			}
		}
		if count == sz {
			break
		}
	}
}

func checkMultiplyAnswer(sz int,grps int,answer int) bool {
	if (sz * grps) == answer {
		return true
	} else {
		return false
	}
}

func main() {
	var ans int
	max := 10
	things := [...]string{"Television","Computer","Radio","Orange","Rose","Axe","Mobile Phone","Dress","Ball","Statue"}
	owner := [...]string{"Jane","Richard","Lila","Susan","George","Ansh","Dhriya","Aditi","Chinmay","Eric"}
	initialize(problems)
	neworder := rand.Perm(100)
	var arraymember int
	var count int	
	for i := 0;i < 100;i++ {
		count = 0
		arraymember = rand.Intn(max)
		fmt.Print(owner[arraymember]," has ",problems[neworder[i]].numsize," ",things[arraymember],"'s in ",problems[neworder[i]].numgroup," groups, how many ",things[arraymember]," does ",owner[arraymember]," have altogether.","\n","Response : ")
		fmt.Scanf("%d",&ans)
		if checkMultiplyAnswer(problems[neworder[i]].numsize,problems[neworder[i]].numgroup,ans) {
			fmt.Println("*****Great answer!! Congratulations, lets try another now.*****")
		} else {
			fmt.Println("Sorry!! Wrong Answer :-( , lets try again.")
			count++
			for count < 3 {
				fmt.Print("Response : ")
				fmt.Scanf("%d",&ans)
				if checkMultiplyAnswer(problems[neworder[i]].numsize,problems[neworder[i]].numgroup,ans) {
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
