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

type collectionID int
type duplicatedFoodTracker map[rune][]collectionID

// This is the core function which will:
// - look into each collection of food
// - identifies the food which is present AT LEAST ONCE in all given food collections
// NOTE: This function is capable to process a singe rucksack (quest 1) or a group of rucksack (quest 2)
// Example Quest 1:
// Raw Input: "vJrwpWtwJgWrhcsFMMfFFhFp" -> will be spitted correctly by function splitIntoCompartments
// Input: [][]rune{"vJrwpWtwJgWr","hcsFMMfFFhFp"}
// Return value: []rune{'p'} -> only the letter p is present AT LEAST ONCE
//
// Example Quest 2:
// Raw Input: "vJrwpWtwJgWrhcsFMMfFFhFp
//
//	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
//	PmmdzqPrVvPwwTWBwg" -> will be spitted correctly by function getNextGroup
//
// Input: [][]rune{"vJrwpWtwJgWrhcsFMMfFFhFp","jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}
// Return value: []rune{'r'} -> only the letter r is present AT LEAST ONCE in all three elf rucksacks
func ingestCompartments(foodCollections [][]rune) []rune {
	tracker := make(duplicatedFoodTracker, 0)

	for compartmentId, foodCollection := range foodCollections {

		// (1)
	foodLoop:
		for _, food := range foodCollection {
			// check if the food is present
			trackedCompartmentsForFood, ok := tracker[food]

			// if not present add it to the tracker and move on to the next food (1)
			if !ok {
				tracker[food] = append(tracker[food], collectionID(compartmentId))
				continue
			}

			// when the food was found we go through each collection ID
			for _, tc := range trackedCompartmentsForFood {
				// if the food was already tracked for the current collection we go to the next food (1)
				if int(tc) == compartmentId {
					continue foodLoop
				}
			}
			// the food is not present in the tracker for the current collection, add it to the collection
			tracker[food] = append(tracker[food], collectionID(compartmentId))
		}
	}

	result := make([]rune, 0)
	for food, count := range tracker {
		// only get the food which was present on all three food collections.
		if len(count) == len(foodCollections) {
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

// specific for quest 1
func splitIntoCompartments(rucksack []rune) [][]rune {
	if len(rucksack)%2 != 0 {
		log.Fatalf("total food count is not a even number")
	}
	return [][]rune{rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]}
}

func runQuest1(file string) int32 {

	reader := readFile(file)
	defer reader.Close()

	var points int32
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		points += calculatePointsForRucksack(ingestCompartments(splitIntoCompartments([]rune(scanner.Text()))))
	}
	return points
}

// specific for quest 2
func getNextGroup(scanner *bufio.Scanner) [][]rune {
	var group [][]rune
	for i := 0; i < 3; i++ {
		if ok := scanner.Scan(); !ok {
			return nil
		}
		group = append(group, []rune(scanner.Text()))
	}
	return group
}

func runQuest2(file string) int32 {
	reader := readFile(file)
	defer reader.Close()

	var points int32
	scanner := bufio.NewScanner(reader)

	for {
		group := getNextGroup(scanner)
		if group == nil {
			break
		}
		points += calculatePointsForRucksack(ingestCompartments(group))
	}
	return points
}

// for all
func readFile(file string) io.ReadCloser {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func main() {
	fmt.Printf("Quest 1: Points %d\n", runQuest1("input_test.txt"))
	fmt.Printf("Quest 1: Points %d\n", runQuest1("input.txt"))

	fmt.Printf("Quest 2: Points %d\n", runQuest2("input_test.txt"))
	fmt.Printf("Quest 2: Points %d\n", runQuest2("input.txt"))
}
