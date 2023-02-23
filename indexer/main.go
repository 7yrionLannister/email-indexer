package main

import (
	"encoding/json"
	"io/ioutil"
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
	r := configureChiRouter()
	log.Println(http.ListenAndServe("localhost:6060", r))
}

func configureChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/process-emails", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		// iterate all files and directories in the user folder recursively and perform the given function on each item
		filepath.Walk("maildir/"+name+"/", processEmailFile)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("Finished processing emails!"))
	})
	r.Get("/get-emails", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // allow the Vue application to acces this method
		w.Write([]byte(fetchEmails(name)))
	})
	r.HandleFunc("/debug/heap", func(w http.ResponseWriter, r *http.Request) {
		pprof.Handler("heap").ServeHTTP(w, r)
	})
	r.HandleFunc("/debug/cpu", func(w http.ResponseWriter, r *http.Request) {
		pprof.Profile(w, r)
	})
	return r
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
			firstIndexOf := strings.Index(currentLine, ":")
			if firstIndexOf < 0 {
				// this line is not a new key-value pair, but an addition to the previos field
				m[lastKey] += currentLine
			} else {
				// this line is a new key-value pair
				lastKey = currentLine[0:firstIndexOf]
				m[lastKey] = currentLine[firstIndexOf+1:]
			}
			startOfBody += len(currentLine)
		} else {
			// first empty line marks the beginning of the body
			startOfBody += i
			m["body"] = fileContent[startOfBody+1:]
			break
		}
	}
	saveEmailInfoToDatabase(m, strings.Split(path, "/")[1])
	return nil
}

func saveEmailInfoToDatabase(m map[string]string, user string) {
	data, _ := json.Marshal(m)
	// each user is an index
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+user+"/_doc", strings.NewReader(string(data)))
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

func fetchEmails(name string) string {
	data := `
	{
		"searchtype":"alldocuments",
    	"max_results": 999999999
	}
	`
	// each user is an index
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+name+"/_search", strings.NewReader(string(data)))
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
	strResponse, _ := ioutil.ReadAll(resp.Body)
	return string(strResponse)
}
