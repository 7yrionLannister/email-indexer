package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"os"
	"path/filepath"
)

var currentIndex string

func main() {
	rootPath := "maildir"
	rootDir, _ := os.ReadDir(rootPath)
	// iterate all users
	for i := 0; i < len(rootDir); i++ {
		currentIndex = rootDir[i].Name()
		// iterate all files and directories in userDir recursively and perform the given function on each item
		filepath.Walk(rootPath+"/"+currentIndex+"/", processEmailFile)
	}
}

func processEmailFile(path string, info os.FileInfo, err error) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	fileContent := string(file)
	m := make(map[string]string)
	lines := strings.Split(fileContent, "\n")
	startOfBody := 0
	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		if strings.Contains(currentLine, ":") {
			keyVal := strings.Split(currentLine, ":")
			m[keyVal[0]] = keyVal[1]
			startOfBody += len(currentLine)
		} else {
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
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+currentIndex+"/_doc", strings.NewReader(string(data)))
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
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
