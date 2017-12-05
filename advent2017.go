package main

import "io/ioutil"
import "fmt"
import "strings"
import "strconv"
import "os"

func main() {

	milisecond := os.Args[1]

	switch milisecond {
	case "1":
		captcha()
	case "2":
		checksum()
	case "3":
		var arg int
		if len(os.Args) == 3 {
			arg, _ = strconv.Atoi(os.Args[2])
		} else {
			arg = 361527
		}
		spiral(arg)
	case "4":
		passphrase()
	default:
		fmt.Println("happy holidays")
	}
}

func captcha() {
	stream, _ := ioutil.ReadFile("input1.txt")
	buffer := string(stream)
	sum := 5
	placeHolder := 0
	for _, i := range strings.Split(buffer, "") {
		j, _ := strconv.Atoi(i)
		if placeHolder == j {
			sum += j
		}
		fmt.Println(i, placeHolder, sum)
		placeHolder = j
	}
	fmt.Println(sum)
}

func checksum() {
	stream, _ := ioutil.ReadFile("input2.txt")
	input := string(stream)
	rowSplit := strings.Split(input, "\n")
	var max int
	var min int
	var sum uint
	sum = 0
	for _, row := range rowSplit {
		max = 0
		min = 99999
		for _, value := range strings.Fields(row) {
			number, _ := strconv.Atoi(value)
			if number > max {
				max = number
			}
			if number < min {
				min = number
			}
		}
		sum += uint((max - min))
		fmt.Println(max, min, sum)
	}

}

func spiral(node int) {
	layer := 0
	step := 8
	var i int
	var temp int
	var temp2 int
	temp = 0
	for i = 0; layer < node; i++ {
		layer += (step * i)
		temp2 = temp
		temp = node - layer
	}
	distance := i - 1
	fmt.Println(distance, (temp2-1)%distance)
	fmt.Println(distance + (temp2-1)%distance)
}

func passphrase() {
	stream, _ := ioutil.ReadFile("input4.txt")
	input := string(stream)
	lines := strings.Split(input, "\n")
	valid := 0
	var dupe bool
	for _, phrase := range lines {
		dupe = false
		words := make(map[string]bool)
		for _, word := range strings.Fields(phrase) {
			if !words[word] {
				words[word] = true
			} else {
				dupe = true
			}
		}
		if dupe == false {
			valid++
		} else {
			fmt.Println(phrase)
		}
	}
	fmt.Println(valid, len(lines))
}
