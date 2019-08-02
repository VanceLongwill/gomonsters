package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

// @TODO: add more tests
func TestMoveMonster(t *testing.T) {
	file, _ := os.Open("assets/world_map_small.txt")
	defer file.Close()
	r := NewCSVReader(file)
	inputChannel := r.ReadAll()
	world, _ := BuildWorldFromRecords(inputChannel)

	var b bytes.Buffer
	output := bufio.NewWriter(&b)
	game := NewMonsterGame(world, 100, 20, output)
	game.Start()
	output.Flush()
	t.Logf(b.String())
}
