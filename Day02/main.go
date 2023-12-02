package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id   int
	sets []set
}

type set struct {
	red, green, blue int
}

var filename = "Day02/input_complete"
var regexGame = regexp.MustCompile(`Game (\d+):(.*)`)
var regexHand = regexp.MustCompile(`(\d+) (blue|red|green)`)

// Possible improvements:
// - Resolve the challenge when reading the file (both of them)

func main() {
	// -- Challenge 1
	sumValidIds := 0
	games := make([]game, 0)

	//read input
	dat, err := os.Open(filename)
	Check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		l := scanner.Text()
		games = append(games, readLine(l))
	}
	// now validate and sumValidIds the valid
	for _, g := range games {
		if isValidGame(g) {
			sumValidIds += g.id
		}
	}
	fmt.Printf("Challlenge 1: %d \n", sumValidIds)

	// -- Challenge 2
	sumPower := 0
	for _, g := range games {
		sumPower += calculatePower(g)
	}

	fmt.Printf("Challlenge 2: %d \n", sumPower)
}

func readLine(l string) game {
	g := game{
		sets: make([]set, 0),
	}

	match := regexGame.FindStringSubmatch(l)

	if len(match) > 1 {
		gameNumber := match[1]
		id, err := strconv.Atoi(gameNumber)
		Check(err)
		g.id = id
		fmt.Println("Game number: ", gameNumber)
		setsStr := match[2]
		fmt.Println("Sets: ", setsStr)

		// parse the sets
		for _, setStr := range strings.Split(setsStr, ";") {
			s := set{}
			for _, hand := range strings.Split(setStr, ",") {
				matchHand := regexHand.FindStringSubmatch(hand)
				q, err := strconv.Atoi(matchHand[1])
				Check(err)
				switch matchHand[2] {
				case "green":
					s.green += q
				case "red":
					s.red += q
				case "blue":
					s.blue += q
				default:
					panic("value not defined")
				}
			}
			g.sets = append(g.sets, s)
		}
	} else {
		fmt.Println("No match found")
	}
	return g
}

func isValidGame(g game) bool {
	for _, s := range g.sets {
		if !isValidSet(s) {
			return false
		}
	}
	return true
}

func isValidSet(s set) bool {
	return s.red <= 12 && s.green <= 13 && s.blue <= 14
}

func calculatePower(g game) int {
	minGreen, minBlue, minRed := g.sets[0].green, g.sets[0].blue, g.sets[0].red
	for _, s := range g.sets {
		if s.green > 0 && s.green > minGreen {
			minGreen = s.green
		}

		if s.red > 0 && s.red > minRed {
			minRed = s.red
		}

		if s.blue > 0 && s.blue > minBlue {
			minBlue = s.blue
		}
	}
	return minGreen * minBlue * minRed
}

func Check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
