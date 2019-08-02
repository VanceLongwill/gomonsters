package main

import "errors"

// CityName is a string used to identify a city
type CityName string

// City is an stateful location in the world map
type City struct {
	// Name of the city (assumed to be unique)
	Name CityName
	// Monsters currently present in the city
	Monsters *MonsterCollection
	// Whether the city has been destroyed or not
	Destroyed bool
	// The number of monsters which causes the city to be destroyed and unreachable, -1 for never destroyed
	maxMonsters int
}

var (
	// ErrCityFull is returned when an operation requiring free space is performed on a city with no free space
	ErrCityFull = errors.New("City has already reached capacity")
	// ErrCityDestroyed is return when attempting to destroy a city which has already been destroyed
	ErrCityDestroyed = errors.New("City is already destroyed")
)

// AddMonster adds a monster to the city's holdings and causes the city to be destoyed if it has subsequently reached capacity
func (c *City) AddMonster(monster *Monster) (bool, error) {
	if c.maxMonsters == -1 {
		c.Monsters.Add(monster)
		return c.Destroyed, nil
	}

	if c.Monsters.Length() > c.maxMonsters-1 {
		if !c.Destroyed {
			c.Destroy()
		}
		return c.Destroyed, ErrCityFull
	}

	c.Monsters.Add(monster)

	// If the maximum number of monsters has been reached, destroy the city
	if c.Monsters.Length() == c.maxMonsters {
		if err := c.Destroy(); err != nil {
			return c.Destroyed, err
		}
	}
	return c.Destroyed, nil
}

// RemoveMonster removes a monster from the city's holdings
func (c *City) RemoveMonster(monster *Monster) {
	c.Monsters.Remove(monster)
}

// Destroy marks the city as destroyed
func (c *City) Destroy() error {
	if c.Destroyed {
		return ErrCityDestroyed
	}
	c.Destroyed = true
	return nil
}

// NewCity creates a new city instance
func NewCity(cityName CityName, maxMonsters int) *City {
	return &City{Name: cityName, maxMonsters: maxMonsters, Monsters: NewMonsterCollection()}
}
