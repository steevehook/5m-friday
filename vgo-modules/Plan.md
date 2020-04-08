# Plan

### Usual Intro

- Name, topic & 5m friday intro

### Small into of what we will do

- Create a new project using Go modules
- Explore `go mod` commands
- Explore `go get` commands that work with modules
- Explore semantic import versioning
- Explore the ins & outs Go modules
- Migrate a small project from `dep` pkg manager to Go modules & how easy it is

### Generate `go.mod`

Let's create a small project using Go modules
We need 2 files, `go.mod` and `go.sum`

Let's generate the `go.mod` file:
```bash
go mod init github.com/steevehook/go-modules
```

Let's add 2 dependencies:
```bash
go get github.com/julienschmidt/httprouter
go get go.uber.org/zap
```

### Explain anatomy of `go.mod` - **The ingredients**

- `module` and good `module` naming
- `go` version
- `require` group
- explain `// indirect` (we don't have any code that uses these packages)

### Explain `go.sum` - **The instructions**
- It contains the entire dependency graph (transitive deps)
- It is a snapshot of a reproducible build
- Normally never modify this file, it is generated

### Let's get rid of `// indirect` & add a small example

- `//indirect` will stay there if we have deps that use them, yet our code does not use them
- open `simple-http` example
- run `go mod tidy`a
- add dep `go get github.com/stretchr/testify`
- run `go mod why go.uber.org/zap`
- run `go mod why github.com/stretchr/testify`

### Explain `go mod tidy`

- `Clears` the `go.mod` of unused packages aka `indirect`
- `Adds` referenced packages aka `direct`
- Good to run after `migrating` to new modules or `experimenting` around

### Quick Notes

- I did not develop this example inside `$GOPATH`
- Module name does not have to be $GOPATH like, but it's good if it is

---

### Let's explore the `go get` variations

- `go get` is used to add/update dep & install all deps

#### add/update a dependency

adds/updates to latest (by default, so can be omitted)
```bash
go get github.com/julienschmidt/httprouter@latest
```

adds/updates to 1.2
`v` is required
```bash
go get github.com/julienschmidt/httprouter@v1.2
```

adds/updates to pseudo version
`v` is required
```bash
go get github.com/julienschmidt/httprouter@90ef33e1709f1f049b2ecc4dfab58869d93eff14
```

- Explain the `pseudo commit anatomy`, look up the `go.mod` file

#### fetch all deps

- run `go get ./...` useful for auto completion and IDE tools
- run `go run/test/build`
- `go clean -cache -modcache -i -r` clears download cache
- `go mod vendor` useful if developing inside `$GOPATH`, IMHO not really useful
- pass `-mod=vendor` to respect vendor directory feature
---

### Explain `$GOPATH/pkg/mod`

- `$GOPATH/pkg/mod` downloaded deps source
- `cache/download` deps download cache

### List all versions

```bash
go list -m -versions go.uber.org/zap
```

### Pre-fill the download cache

`go mod download` downloads deps indicated in `go.mod` without using source code

### Let's explain the difference in dep versions

`v0` - breaking changes
`v1` - stable
`v2+` - major bump, requires path change. Effort from authors

### Explain env variables

- `$GOPROXY`
used as an alternative to fetch deps instead of using VCS
defaults to `https://proxy.golang.org,direct`
setting it to `direct` means do not use any proxy, always download from source VCS

- `$GOSUMDB`
used to consult for checksums of modules
defaults to `https://sum.golang.org`
- Have a look inside `$GOPATH/pkg/sumdb/sum.golang.org`

- GOPRIVATE=*.corp.example.com,rsc.io/private

### Outro

- Mention resources section for all `code used`
- Get your packages & leave
