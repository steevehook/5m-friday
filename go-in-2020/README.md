# Go in 2020

Current Go version: 1.14

## Pros

- C like (statically typed)
- Productive (easy prototyping)
- Great community
- Robust & easy to learn (Fixed, limited feature set)
- Great STD lib
- Fast
- Great concurrent model
- Garbage collected
- Great package management
- Built in testing & profiling
- Cross platform 

## Cons

- No generics
- Clunky error management
- Nil interface values (nil pointer dereference)
- Discourages ORMs
- No FP
- No OOP
- Big binary size
- Not made for ultra high performance due to GC

## Use cases

- CLI tools
- System tools: Docker, CockroachDB, InfluxDB
- Server side services: PUB/SUB server and clients, caching mechanisms, integration layers
- API(s): REST, gRPC, GraphQL
- DevOps tools
- Cloud & networked tools
- Website building: Hugo (Not my preference)

### C like

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

### Prototyping

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Community

8.2% According to Stack Overflow
21M developers (estimated in 2016)

1.7M out of 21M

### Limited feature set

Only 25 keywords

### Fast

[Benchmark](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go.html)

### Package management

- Independent & decentralized
- Private packaging
- Custom paths
- Built in modules with semver support

### Cross platform

- build tags
- suffix file names
- `$GOOS` & `$GOARCH`

```bash
env GOOS=windows GOARCH=amd64 go build
```

### Great STD lib

`net`
`net/http`
`os`
`crypto`
`database`
`encoding`
`errors`
`bufio`
`bytes`
`strings`
`strconv`
`text`

### Concurrent model

- [CSP](https://levelup.gitconnected.com/communicating-sequential-processes-csp-for-go-developer-in-a-nutshell-866795eb879d)
- Go routines
- Channels

### Job rate

How easy it is to find a job
Well, depends on location


# 5 reasons to love Go in 2020

- Productivity
- Simplicity
- Reliability
- Robustness
- Maturity

**Note**: Current Go version: 1.14

## Productivity

- C like & statically typed
- Easy to learn, read & write
- Capable & powerful STD lib

## Simplicity

- The language philosophy
- Limited feature set & keywords (no fancy features like generics)
- Simple & limited data structures (no classes, just funcs and types)

## Reliability

- Blazing fast
- Garbage collected
- Cross platform

## Robustness

- Super stable with limited updates
- 100% compatible & no breaking changes
- Cross platform ($GOOS $GOARCH)

## Maturity

- Baked & maintained by Google
- Powerful STD lib
- Built in pkg manager, testing framework, concurrency model
- Many good libs & tools
- Great docs, examples & resources
- Growing market depending on location (everyone is catching up, from PHP, JS, and even Java and C++)
- About 2M gophers

## Temp

- Blazing fast specs with magic
- Cross platform with magic
