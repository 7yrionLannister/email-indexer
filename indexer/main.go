package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var userid = flag.String("userid", "", "username of the owner of the emails to process`")

func main() {
	flag.Parse()
	cpuProfile()
	rootPath := "maildir"
	// iterate all files and directories in the user folder recursively and perform the given function on each item
	fmt.Println(*userid)
	filepath.Walk(rootPath+"/"+*userid+"/", processEmailFile)

	memoryProfile()
}

func processEmailFile(path string, info os.FileInfo, err error) error {
	file, err := os.ReadFile(path)
	if err != nil { // if there is an error it is probably because "file" is not a file but a directory, so we ignore it
		return nil
	}
	fileContent := string(file)
	m := make(map[string]string)
	lines := strings.Split(fileContent, "\n")
	startOfBody := 0
	lastKey := ""
	// iterate each line of the file
	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		if len(currentLine) > 1 {
			// the line is not empty
			keyVal := strings.Split(currentLine, ":")
			if len(keyVal) == 1 {
				// this line is not a new key-value pair, but an addition to the previos field
				m[lastKey] += currentLine
			} else {
				// this line is a new key-value pair
				lastKey = keyVal[0]
				m[keyVal[0]] = keyVal[1]
			}
			startOfBody += len(currentLine)
		} else {
			// first empty line marks the beginning of the body
			startOfBody += i
			m["body"] = fileContent[startOfBody+1:]
			break
		}
	}
	saveEmailInfoToDatabase(m)
	return nil
}

func saveEmailInfoToDatabase(m map[string]string) {
	data, _ := json.Marshal(m)
	// each user is an index
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+*userid+"/_doc", strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func cpuProfile() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
}

func memoryProfile() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
