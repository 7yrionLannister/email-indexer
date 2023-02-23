package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/process-emails", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		// iterate all files and directories in the user folder recursively and perform the given function on each item
		filepath.Walk("maildir/"+name+"/", processEmailFile)
		w.Write([]byte("Finished processing emails!"))
	})
	r.HandleFunc("/debug/heap", func(w http.ResponseWriter, r *http.Request) {
		pprof.Handler("heap").ServeHTTP(w, r)
	})
	r.HandleFunc("/debug/cpu", func(w http.ResponseWriter, r *http.Request) {
		pprof.Profile(w, r)
	})
	log.Println(http.ListenAndServe("localhost:6060", r))
}

func processEmailFile(path string, info os.FileInfo, err error) error {
	log.Println("inicio processEmailFile para " + path)
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
	log.Println("fin processEmailFile para " + path)
	return nil
}

func saveEmailInfoToDatabase(m map[string]string) {
	data, _ := json.Marshal(m)
	// each user is an index
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+"*userid"+"/_doc", strings.NewReader(string(data)))
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
