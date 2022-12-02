package main

import (
	"bufio"
	"strings"
	"testing"
)

const testInput = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestTopOne(t *testing.T) {
	s := run(bufio.NewScanner(strings.NewReader(testInput)), 1)
	if s != 24000 {
		t.Error(s)
	}
}

func TestTopThree(t *testing.T) {
	s := run(bufio.NewScanner(strings.NewReader(testInput)), 3)
	if s != 45000 {
		t.Error(s)
	}
}
