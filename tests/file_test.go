package tests

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {

	_, err := os.Stat("path")
	if err != nil {
		os.Mkdir("path", 0755)
	}
	f, err := os.Create(filepath.Join("path", "Amogus.txt"))
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	defer f.Close()
	log.Println(f.Name())
}
