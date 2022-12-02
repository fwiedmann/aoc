package main

import "testing"

func TestCalcMove(t *testing.T) {

	tt := []struct {
		name   string
		elf    string
		mine   string
		points int32
	}{
		{
			name:   "WIN",
			elf:    "A",
			mine:   "Y",
			points: 8,
		},
		{
			name:   "LOSE",
			elf:    "B",
			mine:   "X",
			points: 1,
		},
		{
			name:   "DRAW",
			elf:    "C",
			mine:   "Z",
			points: 6,
		},
		{
			name:   "LOSE",
			elf:    "A",
			mine:   "Z",
			points: 3,
		},
	}

	for _, te := range tt {
		t.Run(te.name, func(t *testing.T) {
			result := calcMove(te.elf, te.mine)

			if result != te.points {
				t.Errorf("invalid calculation got %d, but want %d", result, te.points)
			}
		})
	}
}

func TestScoreCalc(t *testing.T) {
	input := []GameMove{
		{
			elf:       "A",
			encrypted: "Y",
		},
		{
			elf:       "B",
			encrypted: "X",
		},
		{
			elf:       "C",
			encrypted: "Z",
		},
		{
			elf:       "A",
			encrypted: "X",
		},
	}

	result := calculateScoreForFirstQuest(input)
	if result != 19 {
		t.Errorf("invalid calculation got %d, but want %d", result, 19)
	}
}
