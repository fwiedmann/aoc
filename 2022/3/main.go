// main day 3 <3
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fwiedmann/aoc/2022/3/alphabet"
)

// rucksacks
// with each two compartments
// 1 Item type 1 compartment
// type by letters

// rucksack:
// 1 line is one rucksack
// 1 compartment == the half of one line
//

// Priorities
// a-z: 1-26
// A-Z: 27-52

type duplicatedFoodTracker map[rune]int32

func ingestCompartments(left, right []rune) []rune {
	tracker := make(duplicatedFoodTracker, 0)

	for _, food := range left {
		tracker[food] = 1
	}

	for _, food := range right {
		if count, ok := tracker[food]; ok {
			tracker[food] = count + 1
		}
	}
	result := make([]rune, 0)

	for food, count := range tracker {
		if count > 1 {
			result = append(result, food)
		}
	}

	return result
}

func calculatePointsForRucksack(duplicatedFood []rune) int32 {
	var points int32
	for _, food := range duplicatedFood {
		point, ok := alphabet.Alphabet[string(food)]
		if !ok {
			log.Fatalf("letter %s not found", string(food))
		}
		points += int32(point)
	}
	return points
}

func splitIntoCompartments(rucksack []rune) ([]rune, []rune) {
	if len(rucksack)%2 != 0 {
		log.Fatalf("total food count is not a even number")
	}
	return rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
}
func run() int32 {

	reader := readFile()
	defer reader.Close()

	var points int32
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		points += calculatePointsForRucksack(ingestCompartments(splitIntoCompartments([]rune(scanner.Text()))))
	}
	return points
}

func readFile() io.ReadCloser {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func main() {
	fmt.Printf("Points %d\n", run())
}
