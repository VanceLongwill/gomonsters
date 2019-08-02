package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// MonsterGame represents the game state
type MonsterGame struct {
	world          *World             // World map to navigate
	ActiveMonsters *MonsterCollection // Keep track of monsters which are not dead or trapped in a location
	done           bool               // Is the game finished
	maxIterations  int                // Maximum number of steps before the game finishes
	logger         io.Writer          // Log for output
	rand           *rand.Rand         // Random number generator
}

// Start runs the game until completion
func (g *MonsterGame) Start() {
	for i := 0; i < g.maxIterations; i++ {
		if g.done {
			break
		}
		// @TODO: introduce concurrency with mutexes on stateful struct fields
		g.step()
	}
	g.done = true
}

// MoveMonsterRandomly transports a monster from its previous location (if any) to another random location
func (g *MonsterGame) MoveMonsterRandomly(monster *Monster) error {
	var possibleDestinations []*City

	if monster.Location() != "" {
		destinations, err := g.world.FindPossibleDestinations(monster.Location())
		if err != nil {
			return err
		}
		if len(destinations) == 0 {
			// Trapped monsters are not active
			if err := g.ActiveMonsters.Remove(monster); err != nil {
				return err
			}
			return nil
		}
		possibleDestinations = destinations
	} else {
		// possibleDestinations are all remaining cities if there is no previous location
		possibleDestinations = g.world.GetUndestroyedCities()
		if len(possibleDestinations) == 0 {
			// The game is finished if all cities are destroyed
			g.done = true
			return nil
		}
	}

	// Remove the monster from the source city (if one has been set)
	if monster.Location() != "" {
		g.world.GetCity(monster.Location()).
			RemoveMonster(monster)
	}

	// Random index with which to choose the destination
	i := g.rand.Intn(len(possibleDestinations))
	// Random destination city
	destCity := possibleDestinations[i]

	// Try to add the monster to the destination destCity
	destroyed, err := destCity.AddMonster(monster)

	// Check if the destCity unexpectedly ran out of space
	if err != nil {
		return err
	}

	monster.SetLocation(destCity.Name)

	if destroyed {
		count := 0
		var msg bytes.Buffer
		msg.WriteString(fmt.Sprintf("%s has been destroyed by ", destCity.Name))
		for _, deadMonster := range destCity.Monsters.GetAll() {
			// Dead monsters are not active
			g.ActiveMonsters.Remove(deadMonster)

			// Pretty print the monsters list
			// E.g. monster 0, monster 1 and monster 2 etc
			if count < destCity.Monsters.Length()-1 {
				msg.WriteString(fmt.Sprintf("monster %s", deadMonster.Name()))
				if count != destCity.Monsters.Length()-2 {
					msg.WriteString(", ")
				}
			} else {
				msg.WriteString(fmt.Sprintf(" and monster %s!\n", deadMonster.Name()))
			}
			count++
		}
		// Write the message to the game logger
		g.logger.Write(msg.Bytes())
	}
	return nil
}

// step: runs one iteration of the game
func (g *MonsterGame) step() {
	// Game is done when no active monsters are left
	if g.ActiveMonsters.Length() == 0 {
		g.done = true
		return
	}
	for _, monster := range g.ActiveMonsters.GetAll() {
		if err := g.MoveMonsterRandomly(monster); err != nil {
			panic(err)
		}
	}
}

// NewMonsterGame sets up a monster game, adding a number of monsters to the provided world in randomly selected locations
func NewMonsterGame(w *World, maxIterations int, initialMonsterCount uint, logger io.Writer) *MonsterGame {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	game := &MonsterGame{
		ActiveMonsters: NewMonsterCollection(),
		world:          w,
		maxIterations:  maxIterations,
		logger:         logger,
		rand:           seededRand,
	}

	for monsterID := uint(0); monsterID < initialMonsterCount; monsterID++ {
		if game.done {
			break
		}
		// Create a monster
		m := NewMonster(monsterID)
		// Add it to the active monsters
		game.ActiveMonsters.Add(m)
		// Place it on the map at random
		if err := game.MoveMonsterRandomly(m); err != nil {
			panic(err)
		}
	}

	return game
}
