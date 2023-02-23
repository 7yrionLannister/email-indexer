# email-indexer

## Instructions

### Running the indexer

```bash
cd indexer
go run .
```
A web server is started with `host=localhost` and `port=6060`. Below you can find a summary of the methods it allows

| Method | URL | Description |
| - | - | - |
| GET | http://localhost:6060/process-emails?name=arnold-j | Reads Arnold's emails and saves them in the ZincSearch database |
| GET | http://localhost:6060/get-emails?name=arnold-j | Fetches Arnold's emails from the database and returns a JSON response |
| GET | http://localhost:6060/debug/heap?seconds=15 | Run the heap profiling for 15 seconds |
| GET | http://localhost:6060/debug/cpu?seconds=15 | Run the CPU profiling for 15 seconds |


### Generating and analyzing CPU and memory profiles

Do this to run the heap profiling for 15 seconds and save the results to a file

```bash
curl http://localhost:6060/debug/heap?seconds=15 --output heap.prof
```

Do this to run the CPU profiling for 15 seconds and save the results to a file

```bash
curl http://localhost:6060/debug/cpu?seconds=15 --output cpu.prof
```

This way you can generate an graph of the given profile file
```go
go tool pprof myprofile.prof
web
```

### Running the web GUI to query emails

```bash
cd email-viewer
npm install
npm run dev
```
