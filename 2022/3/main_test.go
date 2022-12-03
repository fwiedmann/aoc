package main

import (
	"testing"
)

func TestIngestCompartments(t *testing.T) {

	tt := []struct {
		name  string
		left  []rune
		right []rune
		want  rune
	}{
		{
			name:  "",
			left:  []rune("vJrwpWtwJgWr"),
			right: []rune("hcsFMMfFFhFp"),
			want:  rune('p'),
		},
	}

	for _, te := range tt {
		t.Run(te.name, func(t *testing.T) {
			result := ingestCompartments([][]rune{te.left, te.right})
			if result[0] != te.want {
				t.Errorf("invalid result")
			}
		})
	}

}

func TestCalculatePointsForRucksack(t *testing.T) {

	tt := []struct {
		name  string
		input []rune
		want  int32
	}{
		{
			name:  "single p",
			input: []rune{rune('p')},
			want:  16,
		},
		{
			name: "sample calc from demo",
			input: []rune{
				rune('p'),
				rune('L'),
				rune('P'),
				rune('v'),
				rune('t'),
				rune('s'),
			},
			want: 157,
		},
	}

	for _, te := range tt {
		t.Run(te.name, func(t *testing.T) {
			result := calculatePointsForRucksack(te.input)
			if result != te.want {
				t.Errorf("invalid result")
			}
		})
	}

}

func TestSplitIntoCompartments(t *testing.T) {

	result := splitIntoCompartments([]rune("vJrwpWtwJgWrhcsFMMfFFhFp"))

	if string(result[0]) != "vJrwpWtwJgWr" {
		t.Errorf("invalid result for l")
	}

	if string(result[1]) != "hcsFMMfFFhFp" {
		t.Errorf("invalid result for r")
	}
}

func TestXxx(t *testing.T) {

}
