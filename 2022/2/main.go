package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type rule struct {
	is          MovePoint
	loseAginst  MovePoint
	winsAgainst MovePoint
}

var (
	rules []rule = []rule{
		{
			is:          ROCK,
			winsAgainst: SCISSOR,
		},
		{
			is:          PAPER,
			winsAgainst: ROCK,
		},
		{
			is:          SCISSOR,
			winsAgainst: PAPER,
		},
	}
)

func getRule(m MovePoint) rule {
	for _, eu := range rules {
		if eu.is == m {
			return eu
		}
	}
	return rule{}
}

type MovePoint int32

func (m MovePoint) caluclateWithPoints(movePoints int32) int32 {
	return int32(m) + movePoints
}

const (
	ROCK MovePoint = iota + 1
	PAPER
	SCISSOR
)

type Point int32

const (
	LOSE int32 = 0
	DRAW int32 = 3
	WIN  int32 = 6
)

var (
	elfMoves       = map[string]MovePoint{"A": ROCK, "B": PAPER, "C": SCISSOR}
	encryptedMoves = map[string]MovePoint{"X": ROCK, "Y": PAPER, "Z": SCISSOR}
)

func calcMove(elf, encrypted string) int32 {
	elfMove, ok := elfMoves[elf]
	if !ok {
		log.Fatal("unknown move")
	}

	encryptedMove, ok := encryptedMoves[encrypted]
	if !ok {
		log.Fatal("unknown move")
	}

	if elfMove == encryptedMove {
		return encryptedMove.caluclateWithPoints(DRAW)
	}

	if getRule(elfMove).winsAgainst == encryptedMove {
		return encryptedMove.caluclateWithPoints(LOSE)
	}

	return encryptedMove.caluclateWithPoints(WIN)
}

type gameMove struct {
	elf       string
	encrypted string
}

type score int32

func calcScore(moves []gameMove) score {
	var s score
	for _, m := range moves {
		s += score(calcMove(m.elf, m.encrypted))
	}
	return s
}

func readFile() io.ReadCloser {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func main() {
	reader := readFile()
	defer reader.Close()

	var endScore int32 = 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		splited := strings.Split(scanner.Text(), " ")
		if len(splited) != 2 {
			log.Fatal("invalid split")
		}
		endScore += int32(calcScore([]gameMove{{elf: splited[0], encrypted: splited[1]}}))
	}
	fmt.Printf("Score %d\n", endScore)
}
