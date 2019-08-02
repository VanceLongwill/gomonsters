package main

import "fmt"

// Road represents an unidirectional edge in the world graph
type Road struct {
	Direction   string
	Source      CityName
	Destination CityName
}

// NewRoad returns a Road type edge
func NewRoad(dir string, src, dest CityName) *Road {
	return &Road{Direction: dir, Destination: dest, Source: src}
}

// World is the game's map represented by a directed graph
type World struct {
	Cities map[CityName]*City   // Nodes in the graph
	Roads  map[CityName][]*Road // Edges in the graph
}

// NewWorld Create world map
func NewWorld() *World {
	return &World{
		Cities: make(map[CityName]*City),
		Roads:  make(map[CityName][]*Road),
	}
}

// AddCity Add city to world map, return true if added (CityName's are unique).
func (w *World) AddCity(city *City) bool {
	if _, ok := w.Cities[city.Name]; ok {
		return false
	}
	w.Cities[city.Name] = city
	return true
}

// AddRoad Add an edge to the world map
// Note this method does NOT validate that source & destination cities of the road exist
func (w *World) AddRoad(road *Road) {
	// Assumption: roads don't need to be bidirectional because the
	//    - input data takes care of this
	//    - it's not a requirement in this world
	// i.e. we're using a directed graph
	w.Roads[road.Source] = append(w.Roads[road.Source], road)
}

// GetUndestroyedCities returns a list of cities which haven't been destroyed
func (w *World) GetUndestroyedCities() []*City {
	var undestroyed []*City
	for _, city := range w.Cities {
		if !city.Destroyed {
			undestroyed = append(undestroyed, city)
		}
	}
	return undestroyed
}

// FindPossibleDestinations returns a list of possible destinations from a given city
func (w *World) FindPossibleDestinations(cityName CityName) ([]*City, error) {
	var possibleDestinations []*City
	for _, road := range w.GetRoads(cityName) {
		if _, ok := w.Cities[road.Destination]; !ok {
			return nil, fmt.Errorf("Error finding destination city %s: doesn't exist", road.Destination)
		}
		if !w.Cities[road.Destination].Destroyed {
			possibleDestinations = append(possibleDestinations, w.Cities[road.Destination])
		}
	}
	return possibleDestinations, nil
}

// GetCity returns a pointer to a City by its name
func (w *World) GetCity(cityName CityName) *City {
	return w.Cities[cityName]
}

// GetRoads returns a list of Roads leading out of a given city
func (w *World) GetRoads(cityName CityName) []*Road {
	return w.Roads[cityName]
}
