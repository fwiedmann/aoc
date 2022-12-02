// main do stuff
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const elfFile = "massephase.txt"

func main() {
	r := readFile()
	defer r.Close()

	fmt.Println(run(bufio.NewScanner(r)))
}

func run(s *bufio.Scanner) int32 {
	var massElf elf
	currentElf := make(elf, 0)

	for s.Scan() {
		if s.Text() == "" {
			if currentElf.totalCalories() > massElf.totalCalories() {
				massElf = make(elf, len(currentElf))
				_ = copy(massElf, currentElf)
			}
			currentElf.reset()
			continue
		}

		v, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		currentElf.add(int32(v))
	}
	return massElf.totalCalories()
}

type elf []int32

func (e *elf) add(food int32) {
	*e = append(*e, food)
}

func (e *elf) reset() {
	*e = make(elf, 0)
}

func (e *elf) totalCalories() int32 {
	var total int32
	for _, fc := range *e {
		total += fc
	}
	return total
}

func readFile() io.ReadCloser {
	f, err := os.Open(elfFile)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
