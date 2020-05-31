package csvreader

import (
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	goodPath := filepath.FromSlash("../wordlists/4-dice.csv")
	_, err := Read(goodPath)
	if err != nil {
		t.Error("Reading from an existing file should not return an error")
	}

	badPath := filepath.FromSlash("does/not/exist.csv")
	_, err = Read(badPath)
	if err == nil {
		t.Error("Reading from a non-existing file should return an error")
	}
}
