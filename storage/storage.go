package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

// Model interface describes what functions are available to the Storage object
type Model interface {
	Init()
	DumpToFile()
	Store(string, string)
	Delete(string)
	Get(string) (string, error)
	Contains(string) bool
	loadFromFile()
	saveToFile()
}

// Storage struct is the way we're going to store/get all the required links and also store them inside a json file
// For reloading later after the restart. Not doing databases for now.
type Storage struct {
	links    map[string]string
	FilePath string
}

// Init does all the initialization when creating the storage object
func (s *Storage) Init() {
	var f *os.File
	if _, err := os.Stat(s.FilePath); errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(s.FilePath)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = f.WriteString("{}")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	s.links = make(map[string]string)
	s.loadFromFile()
}

// DumpToFile dumps json to file
func (s *Storage) DumpToFile() {
	s.saveToFile()
}

// Store method uses shortener APIs to create a short link and store the long link in a map
// as map[shortLink] = longLink, where later we can use the shortLink that we store to get the full link
func (s *Storage) Store(full string, short string) {
	s.links[short] = full
	s.saveToFile()
}

// Delete deletes a long link from the Storage
func (s *Storage) Delete(link string) {
	delete(s.links, link)
}

// Get returns you the long link that was stored using the StoreLink function
func (s *Storage) Get(shortLink string) (string, error) {
	link := s.links[shortLink]

	if len(link) > 0 {
		return link, nil
	} else {
		return "", errors.New("couldn't find the link inside the storage")
	}
}

// loadFromFile uses the FilePath variable inside the Storage struct to load a json file from the given path
// and then put all the data inside the map[string]string
func (s *Storage) loadFromFile() {
	data, err := ioutil.ReadFile(s.FilePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(data, &s.links)

	if err != nil {
		log.Fatal(err)
		return
	}
}

// saveToFile dumps all the map[string]string data to the FilePath variable inside the Storage struct in a JSON format.
// Format might change later as I assume JSON is not the fastest things possible
func (s *Storage) saveToFile() {
	data, err := json.Marshal(s.links)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile(s.FilePath, data, 0644)

	if err != nil {
		log.Fatal(err)
		return
	}
}

// Contains returns whether our storage contains a given item
func (s *Storage) Contains(link string) bool {
	if _, ok := s.links[link]; ok {
		return ok
	}
	return false
}
