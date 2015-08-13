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

var elements = make(map[string]Element)

type Element struct {
	Id       string
	Name     string   `yaml:"name"`
	BasedOn  []string `yaml:"based_on"`
	combines map[string]string
}

// Combine tries to combine two elements and returns a match if possible
func (e *Element) Combine(other *Element) (combined Element) {
	combined = elements[e.combines[other.Id]]
	return
}

// SetCombines sets up combination relations on the other elements.
// As such, the relation definition is only set on the child elements
// (x + y = z, only z defines x and y as bases), but stored upon setup.
func (e *Element) SetCombines() {
	if e.BasedOn != nil {
		first := elements[e.BasedOn[0]]
		second := elements[e.BasedOn[1]]

		first.combines[second.Id] = e.Id
		second.combines[first.Id] = e.Id
	}
}

func (e *Element) String() string {
	return e.Name
}

func NewElement(id string) (element Element) {
	elem := Element{Id: id}
	elem.combines = make(map[string]string)
	return elem
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

	log.Println("Loading yaml files...")
	for _, file := range fileList {
		loadElement(file, elements)
	}

	log.Println("Setting combination mappings...")
	for _, elem := range elements {
		elem.SetCombines()
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

	elem := NewElement(id)

	yaml.Unmarshal(data, &elem)
	elements[id] = elem
}

func main() {
	loadElementFiles("elements/")

	fire := elements["fire"]
	water := elements["water"]
	log.Printf("Combining fire and water gives: %s\n", fire.Combine(&water).Name)
}
