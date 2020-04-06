# vGo modules

### Intro with stealing packages

Walking with very small steps, then I fall down saying "SHIT"

- "Oh, you know what actually all I need is this"
(taking the CPU i9 box) and run away

### Usual intro

- My name & 5m Intro + topic intro

### A little bit of history about packaging


### `go build/run/test`

Automatically fetch all modules that the project needs

### `go mod`

- Explain `go mod init`
- Example on `go mod init`
- Explain package naming
- Explain `go.mod` anatomy
- Explain `// indirect`
- Reference GOPATH video
- Explain `go.sum`
- Explain `go.mod`
- Explain `go mod tidy` (Good to run after `migrating` to new modules or `experimenting` around)
- Example using `go mod tidy`
- Explain `go mod vendor`
- Example of using `go mod vendor`
- Explain `go mod download`
- Example of using `go mod download`
- Example of using Modules inside libs and deprecations like using `v2` inside path
- Explain `go mod why`
- Example of `go mod why -m github.com/aws/aws-sdk-go`

### `go get`

- Example of using `go get -u ...`
- Example without `@v1.0.1`
- Example with `@v1.0.1`
- Example with `@COMMIT_HASH` or pseudo version
- Explain the difference between `v0`, `v1` and `v2`
- Example using private packages (Mention docker & private key passing)
- Example using `go get ./...`
- Explain `GO111MODULE` variable
- Explain MODULE AWARE feature
- Explain `$GOPATH/pkg/mod`
- Explain `$GOPROXY`
- Explain `$GOSUMDB`
- `go clean -cache -modcache -i -r` to clear cache

### `go list`

- Explain `go list`
- Example using `go list`
- Example with `go list -m all` (all versions that will be used in the final build)
- Example with `go list -u -m all` (all available SEMVER versions if any)
- Example with `go list -u -m all | grep "go.uber.org/multierr" | awk '{print $2}'`
or `go list -m -versions rsc.io/sampler`

### Migrating from `dep`

- Mention that some problems may arise if not enough rights for private packages

### Semantic import versioning

"If an old package and a new package have the same import path,
 the new package must be backwards compatible with the old package"

### Dependency diamond problem

### `go fix`

### Outro

- Mention resources section for all `code used`
- Get your packages & leave

### Key Points to talk about

- Initialize a project to use Go modules
- `go get` variations
- Migrate to modules
- Add a dependency
- Upgrade a dependency

### Resources

- [Using Go Modules](https://blog.golang.org/using-go-modules)
- [Migrating to Go Modules](https://blog.golang.org/migrating-to-go-modules)
- [Publishing Go Modules](https://blog.golang.org/publishing-go-modules)
- [Go Modules v2 & Beyond](https://blog.golang.org/v2-go-modules)
- [Go Modules - GitHub](https://github.com/golang/go/wiki/Modules)
- [Migrating to Go Modules - Supported formats](https://go.googlesource.com/go/+/362625209b6cd2bc059b6b0a67712ddebab312d9/src/cmd/go/internal/modconv/modconv.go#9)
- [vgo import](https://research.swtch.com/vgo-import)
- [Major versions in Go](https://medium.com/@elliotchance/major-versions-in-go-modules-explained-ec7c1df3888b)
- [Go module proxy](https://arslan.io/2019/08/02/why-you-should-use-a-go-module-proxy/)
- [Go module mirror launch](https://blog.golang.org/module-mirror-launch)
- [Go module proxy protocol](https://golang.org/cmd/go/#hdr-Module_proxy_protocol)
