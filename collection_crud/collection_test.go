package main

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	var f *os.File
	f, _ = os.OpenFile("collection.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	Create(f, "testBird", "testusBirdus", "thisComputer", "uhhPixels", 42)
}

func TestRead(t *testing.T) {
	var f *os.File
	f, _ = os.OpenFile("collection.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	Read(f)
}

func TestUpdate(t *testing.T) {
	var f *os.File
	f, _ = os.OpenFile("collection.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	Update(f, "1", "NorthernMockingbird", "MimusPolyglottos", "13", "NorthAmerica", "Insectivore")
}

func TestDestroy(t *testing.T) {
	var f *os.File
	f, _ = os.OpenFile("collection.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	Destroy(f, "1")
}
