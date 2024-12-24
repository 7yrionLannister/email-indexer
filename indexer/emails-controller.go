package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func processEmailFile(path string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		// it is a file, base case
		file, err := os.ReadFile(path)
		if err != nil {
			log.Println("Error reading file at '", path, "': ", err)
		}
		fileContent := string(file)
		m := make(map[string]string)
		lines := strings.Split(fileContent, "\n")
		startOfBody := 0
		lastKey := ""
		// iterate each line of the file
		for i, currentLine := range lines {
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
	} else {
		// it is a directory,recursive case
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i]
			processEmailFile(path + "/" + dir.Name())
		}
	}
}

func saveEmailInfoToDatabase(m map[string]string, user string) {
	data, _ := json.Marshal(m)
	// each user is an index
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+user+"/_doc", strings.NewReader(string(data)))
	checkError(&err)
	setHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	checkError(&err)
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
	checkError(&err)
	setHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	checkError(&err)
	defer resp.Body.Close()
	strResponse, _ := io.ReadAll(resp.Body)
	return string(strResponse)
}

func checkError(err *error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setHeaders(req *http.Request) {
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
}
