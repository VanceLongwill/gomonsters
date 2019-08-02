## Monster Game

#### Rules of the game

 - See `description.txt`

#### Installation (requires golang)

- Clone this repo
- Build the project

    `go build`

#### Usage

- Run the generated executable with no args to see usage options

    `./monsters`

- A small (default) and medium map are provided in `/assets`

    e.g.
    `./monsters -n 100 -d assets/world_map_medium.txt`

- Results are directed to stdout as default, though they can also be written to a file

    e.g.
    `./monsters -n 100 -d assets/world_map_medium.txt -o results_file.txt`

#### Testing

- Run tests with 
    
    `go test -v`

#### Stack

- Golang standard library (no external packages used)

#### About my solution

- [x] Easily flexible/extendible/configurable/transformable into a larger game/simulator by use of generic interfaces and separation of concerns
- [x] Golang best practices
- [x] Linted with [Golint](https://github.com/golang/lint)
- [x] Formatted with [Gofmt](https://golang.org/cmd/gofmt)

#### TODO:

- [ ] Expand test coverage
- [ ] Improve error handling
- [ ] Introduce concurrency
- [ ] Create a frontend/visualisation/graphic representation for the world map
