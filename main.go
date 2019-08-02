package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	defaultMapDataFn := "assets/world_map_small.txt"

	// Get cli flags
	initialMonsterCount := flag.Uint("n", 0, "specify the number of monsters you want to start with (n > 0)")
	mapDataFn := flag.String("d", defaultMapDataFn, "input file path containing data used to build the game map")
	outputDataFn := flag.String("o", "", "output file path to write the world state after the game, writes to stdout as default")
	flag.Parse()

	// Print usage info if no cli arg is provided
	if *initialMonsterCount == 0 {
		flag.Usage()
		return
	}

	// Open the map data file
	file, err := os.Open(*mapDataFn)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// WorldStateReader interface keeps the process of getting input data generic
	var r WorldStateReader
	// currently only CSV is used
	r = NewCSVReader(file)
	// Read records in the map data file
	inputChannel := r.ReadAll()
	// Build world graph/map based on map data records
	worldOfX, err := BuildWorldFromRecords(inputChannel)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new game instance using map, redirecting output to stdout
	game := NewMonsterGame(worldOfX, 10000, *initialMonsterCount, os.Stdout)

	// Run the game to completion
	game.Start()

	os.Stdout.WriteString("\n")

	// WorldStateWriter interface keeps the process of getting input data generic
	var w WorldStateWriter
	// output writer
	var target io.Writer

	if *outputDataFn != "" {
		// Create the map data output file
		// Contents will be overwritten!
		file, err := os.Create(*outputDataFn)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		target = file
	} else {
		// writing the results to stdout as default
		target = os.Stdout
	}

	// currently only CSV format is used
	w = NewCSVWriter(target)

	// Output what's left of the world
	outputChannel := GetRemainingWorldRecords(worldOfX)
	// Store the records somewhere
	w.WriteAll(outputChannel)
}
