# Go routines

## Intro

- 5m Friday intro
- Intro to concurrency topic
- Throw 5m on the clock

Application level threads not OS threads

## Light

- `min [2KB]` vs `min 512KB` in most cases
- growable segmented stacks
- [Go stack size](https://github.com/golang/go/blob/bbd25d26c0a86660fb3968137f16e74837b7a9c6/src/runtime/stack.go#L72)
- `hundreds of 1000's or millions` vs `1000's or tens of 1000's`
- `max 1 thread/CPU core & many go routines/thread` vs `1:1 thread mapping` (8 CPU Cores = Max of 8 threads managed by the Go Scheduler)

## Fast

- Fast startup time
- Work stealing
- Communication between go routines are faster (due to internal management) = Fast context switching
- Not hardware dependent
- Context switch off on existing created `Thread` vs Context switch oof on a different `Core` (OS thread never moves into waiting state in case of go routines)
- turn IO/Blocking work into CPU-bound work at the OS level

## Simple

- Go routine scheduler (`managed by runtime` vs `managed by OS kernel`) (GRQ & LRQ - run queues)
- No callback hell/spaghetti
- CSP model (communication over channels)
- Just use the `go` keyword

## Safe

- No thread management
- Allow you to avoid mutex locking 


## Facts

- States (Waiting, Runnable, Executing)
- Less is more
- Know the type of work that need to be done concurrently

## References

- Ardan Labs
- JBD
