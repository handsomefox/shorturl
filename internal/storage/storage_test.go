package storage

import (
	"errors"
	"os"
	"testing"
)

const filePath = "C:\\Go\\Saved\\test.json"

func TestInit(t *testing.T) {
	s := Storage{FilePath: filePath}
	s.Init()
	if s.links == nil {
		t.Fatal()
	}
}

func TestStore(t *testing.T) {
	s := Storage{FilePath: filePath}
	s.Init()
	s.Store("full.go", "short.go")

	if !s.Contains("short.go") {
		t.Fatal()
	}
}
func TestDelete(t *testing.T) {
	s := Storage{FilePath: filePath}
	s.Init()
	s.Store("full.go", "short.go")

	s.Delete("short.go")

	if s.Contains("short.go") {
		t.Fatal()
	}
}
func TestGet(t *testing.T) {
	s := Storage{FilePath: filePath}
	s.Init()
	s.Store("full.go", "short.go")
	data, err := s.Get("short.go")

	if err != nil {
		t.Fatal(err)
	}

	if data != "full.go" {
		t.Fatal()
	}
}

func TestContains(t *testing.T) {
	s := Storage{FilePath: filePath}
	s.Init()
	s.Store("full.go", "short.go")

	if !s.Contains("short.go") {
		t.Fatal()
	}
}

func TestDumpToFile(t *testing.T) {
	path := "C:\\Go\\Saved\\dump.json"
	s := Storage{FilePath: path}
	s.Init()
	s.Store("full.go", "short.go")

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
}
func TestLoadFromFile(t *testing.T) {

	s := Storage{FilePath: filePath}
	s.loadFromFile()

	if s.links == nil {
		t.Fatal()
	}
}
func TestSaveToFile(t *testing.T) {
	s := Storage{FilePath: filePath}
	s.links = make(map[string]string)

	s.links["zxc"] = "qwe"

	s.saveToFile()
	s.loadFromFile()

	if s.links["zxc"] != "qwe" {
		t.Fatal()
	}
}
