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
	fmt.Println(run(bufio.NewScanner(r), 3))
}

func run(s *bufio.Scanner, c int) int32 {
	massElfs := initMassElfs(c)
	elfCursor := make(elf, 0)

	for s.Scan() {
		// add elf
		if s.Text() == "" {
			massElfs.add(elfCursor)

			// rest the cursor to count for the next elf
			elfCursor.reset()
			continue
		}

		// add food to current elf
		elfCursor.add(s.Text())
	}

	// when EOF then also add the last built elf
	massElfs.add(elfCursor)
	return massElfs.totalCalories()
}

type elf []int32

func (e *elf) add(food string) {
	v, err := strconv.Atoi(food)
	if err != nil {
		log.Fatal(err)
	}
	*e = append(*e, int32(v))
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

func initMassElfs(count int) MasseElfs {
	return MasseElfs{
		toCount: count,
		elfs:    make([]elf, count),
	}
}

// MasseElfs comment
type MasseElfs struct {
	toCount int
	elfs    []elf
}

func (m *MasseElfs) add(e elf) {
	for i := 0; i < m.toCount; i++ {
		if e.totalCalories() > m.elfs[i].totalCalories() {
			// We have to move all  smaller elfs one position to the right.
			// This creates an empty elf at the index and moves the rest to the right of it
			m.elfs = append(m.elfs[:i+1], m.elfs[i:]...)
			// enter the elf
			m.elfs[i] = e

			// shrink it to the original size
			m.elfs = m.elfs[:m.toCount]
			return
		}
	}
}

func (m *MasseElfs) totalCalories() int32 {
	var c int32
	for i := 0; i < m.toCount; i++ {
		c += m.elfs[i].totalCalories()
	}
	return c
}
