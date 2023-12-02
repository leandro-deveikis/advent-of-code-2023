package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var numbersMap = map[string]string{
	// found the hard way that when a number is replaced,
	// it should not replace the hole number, but the letters can still be used by another one
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

var filename = "Day01/input_complete"

// Possible improvements:
// - Resolve the challenge when reading the file (both of them)
func main() {
	//
	lines := ReadInput(filename)

	//-- Challenge 1
	result1 := sumLines(lines)
	fmt.Printf("Challenge 1 result: %d \n", result1)

	//-- Challenge 2
	newLines := replaceWithNumbers(lines)
	result2 := sumLines(newLines)
	fmt.Printf("Challenge 2 result: %d \n", result2)
}

func sumLines(lines []string) int32 {
	var sum int32 = 0
	for _, l := range lines {
		var first, second int32 = 0, 0
		firstFound := false

		for _, c := range l {
			// 48 is 1 and 57 is 9
			if c >= 48 && c <= 57 {
				second = runeToNumber(c)
				if !firstFound {
					first = runeToNumber(c)
					firstFound = true
				}
			}
		}
		sum += (first * 10) + second
	}
	return sum
}

func replaceWithNumbers(lines []string) []string {
	newLines := make([]string, 0)

	for _, l := range lines {
		for {
			kFound := ""
			minPos := -1
			for k := range numbersMap {
				pos := strings.Index(l, k)
				if pos != -1 && (pos < minPos || minPos == -1) {
					minPos = pos
					kFound = k
				}
			}

			if kFound == "" {
				break
			}
			l = strings.Replace(l, kFound, numbersMap[kFound], 1)
		}
		newLines = append(newLines, l)
	}
	return newLines
}

func runeToNumber(c int32) int32 {
	// simplest way to convert this to the number
	return c - 48
}

func ReadInput(filename string) []string {
	lines := make([]string, 0)
	dat, err := os.Open(filename)
	Check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
