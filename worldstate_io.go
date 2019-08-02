package main

// WorldRecord is generic representation of a city and roads which lead out of it useful in managing data from different formats
type WorldRecord struct {
	City  CityName
	Roads []*Road
}

// WorldStateReader is an generic interface for reading in the state of a World
type WorldStateReader interface {
	// ReadAll: Reads CSV data and returns a channel which it feeds records
	// All world map data input sources should be handled using this interface
	ReadAll() (ch chan *WorldRecord)
}

// WorldStateWriter is an generic interface for writing the state of a World
type WorldStateWriter interface {
	// WriteAll: Writes CSV data from a channel of records
	// All world map data output sources should be handled using this interface
	WriteAll(ch <-chan *WorldRecord)
}

// BuildWorldFromRecords generates a World graph from records passed through a channel (in order to not be specific to a particular input method/source)
func BuildWorldFromRecords(records <-chan *WorldRecord) (*World, error) {
	world := NewWorld()
	maxMonstersPerCity := 2
	for {
		if record, ok := <-records; ok {
			city := NewCity(record.City, maxMonstersPerCity)
			world.AddCity(city)
			for _, road := range record.Roads {
				// Check if destination city exists, if not then create it
				if world.GetCity(road.Destination) == nil {
					world.AddCity(NewCity(road.Destination, maxMonstersPerCity))
				}
				world.AddRoad(road)
			}
		} else {
			break
		}
	}
	return world, nil
}

// GetRemainingWorldRecords finds the remaining cities which are reachable and output's their respective records to the channel
func GetRemainingWorldRecords(w *World) (ch chan *WorldRecord) {
	ch = make(chan *WorldRecord)
	go func() {
		defer close(ch)
		for _, city := range w.GetUndestroyedCities() {
			record := &WorldRecord{City: city.Name}
			possibleDestinations := w.GetRoads(city.Name)
			if len(possibleDestinations) == 0 {
				continue
			}
			var roads []*Road
			for _, dest := range possibleDestinations {
				// Assumption: don't include roads leading to destroyed cities
				if !w.GetCity(dest.Destination).Destroyed {
					roads = append(roads, dest)
				}
			}
			if len(roads) == 0 {
				continue
			}
			record.Roads = roads
			ch <- record
		}
	}()
	return
}
