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

func TestMain(t *testing.T) {
	s := run(bufio.NewScanner(strings.NewReader(testInput)))
	if s != 24000 {
		t.Error(s)
	}
}
