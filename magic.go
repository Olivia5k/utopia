package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var combinations = make(map[string]Element)
var elements = make(map[string]Element)

type Element struct {
	Id   string
	Name string `yaml:"name"`
}

// Combine tries to combine two elements and returns a match if possible
func (e *Element) Combine(other *Element) (combined Element) {
	combined = combinations[e.Id+"-"+other.Id]
	return
}

func (e *Element) String() string {
	return e.Name
}

func loadElementFiles(dir string) {
	fileList := []string{}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".yml") {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, file := range fileList {
		loadElement(file, elements)
	}
}

// loadElement loads the yaml file fn and stores the resulting Element{}
// inside the map argument
func loadElement(fn string, elements map[string]Element) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Extract the filename without the .yml extension
	id := strings.Split(path.Base(fn), ".")[0]

	elem := Element{Id: id}
	yaml.Unmarshal(data, &elem)
	elements[id] = elem
}

func main() {
	loadElementFiles("elements/")

	log.Printf("Elements: %s\n", elements)
	// fire := elements["fire"]
	// ice := elements["ice"]
	// log.Printf("Combining fire and ice gives: %s\n", fire.Combine(&ice))
}
