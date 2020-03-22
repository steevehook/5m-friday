# Go routines

## Intro

- 5m Friday intro
- Intro to 5 simple facts about concurrency
- Throw 5m on the clock

`Note`: Application level threads not OS threads

## Light

- `min [2KB]` vs `min 512KB` in most cases
- growable segmented stacks
- [Go stack size](https://github.com/golang/go/blob/bbd25d26c0a86660fb3968137f16e74837b7a9c6/src/runtime/stack.go#L72)
- `hundreds of 1000's or millions` vs `1000's or tens of 1000's`
- No thin air `max 1 thread/CPU core & many go routines/thread` vs `1:1 thread mapping` (8 CPU Cores = Max of 8 threads managed by the Go Scheduler)

## Speed

- Fast startup time (light & used internally)
- Work stealing (Go scheduler & threads which manage go routines)
- Communication between go routines are faster (due to internal management) = Fast context switching
- Not hardware dependent
- Context switch off on existing created `Thread` vs Context switch oof on a different `Core` (OS thread never moves into waiting state in case of go routines)
- Optional - turn IO/Blocking work into CPU-bound work at the OS level

## Simplicity

- Go routine scheduler (`managed by runtime` vs `managed by OS kernel`) (GRQ & LRQ - run queues)
- No thread ID, or related to thread management
- No callback hell/spaghetti, but CSP model (communication over channels)
- Just use the `go` keyword

## Safety

- No thread management (Go scheduler takes care of all)
- Allows you to avoid semaphores or mutex locking 


## Facts

- States (Waiting, Runnable, Executing)
- Less is more (less threads, more work)
- Optional - Know the type of work that need to be done concurrently

## Add

- `runtime.NumCPU()`
- `runtime.GOMAXPROCS()`

## References

- Ardan Labs
- JBD
