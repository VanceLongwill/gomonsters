package main

import (
	"errors"
	"strings"
)

const (
	monsterNameLength = 8 // Assumption: 8 characters is long enough to ensure uniqueness per game
)

var (
	// ErrMonsterNotFound is returned when attempting to access a monster which doesn't exist in the collection
	ErrMonsterNotFound = errors.New("Monster not found")
	// ErrMonsterDuplicateID is returned when attempting to add a monster with an id already present in the collection
	ErrMonsterDuplicateID = errors.New("Monster with this id already exists")
)

// MonsterID is the identifier for monsters
type MonsterID uint

// Monster represents a monster in the game
type Monster struct {
	ID       MonsterID
	name     string
	location CityName
}

// SetLocation changes the monster's location
func (m *Monster) SetLocation(city CityName) {
	m.location = city
}

// Location returns the current location of the monster
func (m *Monster) Location() CityName {
	return m.location
}

// Name return the name of the monster
func (m *Monster) Name() string {
	return m.name
}

// NewMonster creates a monster with a randomly generated name
func NewMonster(id uint) *Monster {
	return &Monster{
		ID:   MonsterID(id),
		name: strings.Title(GenerateIdentifier(monsterNameLength)),
	}
}

// MonsterCollection is a store of monsters
type MonsterCollection struct {
	monsters map[MonsterID]*Monster
}

// Add adds a monster to the store if it is not already present
func (mc *MonsterCollection) Add(monster *Monster) error {
	if _, ok := mc.monsters[monster.ID]; ok {
		return ErrMonsterDuplicateID
	}
	mc.monsters[monster.ID] = monster
	return nil
}

// Remove removes a monster from the store if it is present
func (mc *MonsterCollection) Remove(monster *Monster) error {
	if _, ok := mc.monsters[monster.ID]; !ok {
		return ErrMonsterNotFound
	}
	delete(mc.monsters, monster.ID)
	return nil
}

// Length returns the number of monsters in the store
func (mc *MonsterCollection) Length() int {
	return len(mc.monsters)
}

// IsEmpty checks whether there are are monsters in the store
func (mc *MonsterCollection) IsEmpty() bool {
	return mc.Length() == 0
}

// GetAll returns a map of monsters by id
func (mc *MonsterCollection) GetAll() map[MonsterID]*Monster {
	return mc.monsters
}

// NewMonsterCollection creates a new monster store
func NewMonsterCollection() *MonsterCollection {
	return &MonsterCollection{
		monsters: make(map[MonsterID]*Monster),
	}
}
