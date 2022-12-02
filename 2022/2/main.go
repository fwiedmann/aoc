// main day 2
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// RuleForSecondQuest defines what the encryptedMove should do
type RuleForSecondQuest struct {
	is         MovePoint
	shouldLose bool
	shouldWin  bool
	shouldDraw bool
}

var (
	secondQuestRules = map[MovePoint]RuleForSecondQuest{
		ROCK: {
			is:         ROCK,
			shouldLose: true,
		},
		SCISSOR: {
			is:        SCISSOR,
			shouldWin: true,
		},
		PAPER: {
			is:         PAPER,
			shouldDraw: true,
		},
	}
)

// GameRule defines the moving rules for a moving point
type GameRule struct {
	is           MovePoint
	winsAgainst  MovePoint
	losesAgainst MovePoint
}

var (
	gameRules = map[MovePoint]GameRule{
		ROCK: {
			is:           ROCK,
			winsAgainst:  SCISSOR,
			losesAgainst: PAPER,
		},
		PAPER: {
			is:           PAPER,
			winsAgainst:  ROCK,
			losesAgainst: SCISSOR,
		},
		SCISSOR: {is: SCISSOR,
			winsAgainst:  PAPER,
			losesAgainst: ROCK,
		},
	}
)

// find the move point which needs to be played in order to win against the given move
func findWinMoveAgainstGivenMove(m MovePoint) MovePoint {
	for _, eu := range gameRules {
		if eu.winsAgainst == m {
			return eu.is
		}
	}
	log.Fatalf("could not find win move for point %d", m)
	return -1
}

// find the move point which needs to be played in order to lose against the given move
func findLoseMoveAgainstGivenMove(m MovePoint) MovePoint {
	for _, eu := range gameRules {
		if eu.losesAgainst == m {
			return eu.is
		}
	}
	log.Fatalf("could not find lose move for point %d", m)
	return -1
}

// MovePoint defines the given Moves and their points when chosen
type MovePoint int32

func (m MovePoint) caluclateWithPoints(movePoints int32) int32 {
	return int32(m) + movePoints
}

const (
	ROCK MovePoint = iota + 1
	PAPER
	SCISSOR
)

// Point defines the available receivable points for win, lose or draw
type Point int32

const (
	LOSE int32 = 0
	DRAW int32 = 3
	WIN  int32 = 6
)

// The mapping to the union type MovePoint
var (
	elfMoves       = map[string]MovePoint{"A": ROCK, "B": PAPER, "C": SCISSOR}
	encryptedMoves = map[string]MovePoint{"X": ROCK, "Y": PAPER, "Z": SCISSOR}
)

// lookup the original encrypted move point key
func getEncryptedMoveForMove(m MovePoint) string {
	for k, v := range encryptedMoves {
		if v == m {
			return k
		}
	}
	log.Fatalf("could not find the encrypted key for the given move point %d", m)
	return ""
}

func calcMove(elf, encrypted string) int32 {
	// lookup the union type / deserialize
	elfMove, ok := elfMoves[elf]
	if !ok {
		log.Fatal("unknown move")
	}

	// lookup the union type / deserialize
	encryptedMove, ok := encryptedMoves[encrypted]
	if !ok {
		log.Fatal("unknown move")
	}

	if elfMove == encryptedMove {
		return encryptedMove.caluclateWithPoints(DRAW)
	}

	rule, ok := gameRules[elfMove]
	if !ok {
		log.Fatal("could not find game rule")
	}

	if rule.winsAgainst == encryptedMove {
		return encryptedMove.caluclateWithPoints(LOSE)
	}

	return encryptedMove.caluclateWithPoints(WIN)
}

type GameMove struct {
	elf       string
	encrypted string
}

func calculateScoreForFirstQuest(moves []GameMove) int32 {
	var score int32
	for _, m := range moves {
		score += calcMove(m.elf, m.encrypted)
	}
	return score
}

func calculateScoreForSecondQuest(moves []GameMove) int32 {
	var score int32
	for _, m := range moves {
		// lookup the union type / deserialize
		encryptedMove, ok := encryptedMoves[m.encrypted]
		if !ok {
			log.Fatal("unknown move")
		}

		// lookup the union type / deserialize
		elfMove, ok := elfMoves[m.elf]
		if !ok {
			log.Fatal("unknown move")
		}

		rule, ok := secondQuestRules[encryptedMove]
		if !ok {
			log.Fatal("unknown rule for second quest for given move")
		}

		switch {
		case rule.shouldDraw:
			score += calcMove(m.elf, getEncryptedMoveForMove(elfMove))
		case rule.shouldLose:
			// get the move which will trigger a lose for the elf
			score += calcMove(m.elf, getEncryptedMoveForMove(findLoseMoveAgainstGivenMove(elfMove)))
		case rule.shouldWin:
			// get the move which will trigger a win for the elf
			score += calcMove(m.elf, getEncryptedMoveForMove(findWinMoveAgainstGivenMove(elfMove)))
		}
	}
	return score
}

func readFile() io.ReadCloser {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// run a quest
func run(calc func(moves []GameMove) int32) {
	reader := readFile()
	defer reader.Close()

	var endScore int32
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		splited := strings.Split(scanner.Text(), " ")
		if len(splited) != 2 {
			log.Fatal("invalid split")
		}

		endScore += int32(calc([]GameMove{{elf: splited[0], encrypted: splited[1]}}))
	}

	fmt.Printf("Score %d\n", endScore)
}

func main() {
	run(calculateScoreForFirstQuest)
	run(calculateScoreForSecondQuest)
}
