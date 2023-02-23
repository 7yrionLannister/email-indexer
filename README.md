# email-indexer

## Instructions

### Running the indexer

The indexer must be executed providing the username of interest. Parameters between brackets are optional and they enable profiling.

```go
go run . -userid 'zufferli-j' [-cpuprofile cpu.prof] [-memprofile mem.prof]
```

### Analyzing the profiles

```go
go tool pprof cpu1.prof
top
```

```bash
pprof -top main.go cpu1.prof
```

```go
go tool pprof cpu1.prof
web
```

```bash
pprof -web main.go cpu1.prof
```