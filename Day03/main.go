package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var filename = "Day03/input_test"

func main() {
	//read input
	input := readInput()

	// -- Challenge 1
	sum1 := 0
	// build adjacent matrix -
	jmax := len(input[0])
	adjMatrix := makeAdjMatrix(len(input), jmax, input)

	for _, str := range adjMatrix {
		fmt.Printf("%s\n", str)
	}

	fmt.Printf("Challlenge 1: %d \n", sum1)

	// -- Challenge 2
	sum2 := 0

	fmt.Printf("Challlenge 2: %d \n", sum2)
}

func makeAdjMatrix(x int, jmax int, input []string) []string {
	result := make([]string, x)
	for i := 0; i < x; i++ {
		result[i] = strings.Repeat(".", jmax)
	}

	for i, str := range input {
		for j, c := range str {
			if !unicode.IsDigit(c) && '.' != c {
				// then is a special character, so we complete the adjacent with "#"
				if i > 0 {
					replaceAdjacent(result, i-1, j, jmax)
				}
				replaceAdjacent(result, i, j, jmax)
				if i < len(result) {
					replaceAdjacent(result, i+1, j, jmax)
				}
			}
		}
	}
	return result
}

func replaceAdjacent(adjMatrix []string, i, j, jmax int) {
	r := []rune(adjMatrix[i])
	if j > 0 {
		r[j-1] = '#'
	}
	r[j] = '#'
	if j < jmax-1 {
		r[j+1] = '#'
	}
	adjMatrix[i] = string(r)
}

func readInput() []string {
	input := make([]string, 0)
	dat, err := os.Open(filename)
	Check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func Check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
