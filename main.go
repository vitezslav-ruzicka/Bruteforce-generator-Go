package main

import (
	"fmt"
)

func nextLetter (login []rune, count []rune, place []int) {

	x := 1;
	for index:=0 ; index<x; index++ {

		//zarazi pripadny presah indexu
		if x==12 {break
		} else if login[index]==count[61] {
			login[index]=count[0]
			place[index]=0
			x++

		} else {
			login[index]=count[place[index] +1]
			place[index]++
		}
	}
}

func setup (login []rune, count []rune, place []int) {

	for i :=0 ; i<12; i++ {
		place[i]=-1
		login[i]=0
	}
	//nastavi prvni hodnoty
	place[0]=0
	login[0]=count[0]
}

func main() {
	count := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	login := make([]rune, 12, 12)
	place := make([]int, 12, 12)

	setup(login, count, place)

	for x :=0 ; x<500000000 ; x++ {
		nextLetter(login, count, place)
		fmt.Println(string(login),x)
		//time.Sleep(5000000)
	}
}