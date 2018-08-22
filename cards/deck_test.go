package main

import (
	"os"
	"testing"
)

var testingFilename = "__testing_file_name"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("deck length should be 52, but got: %v", len(d))
	}

	if d[0] != "♠︎A" {
		t.Errorf("deck's first item should be ♠︎A, but get: %v", d[0])
	}

	if d[len(d)-1] != "♣︎K" {
		t.Errorf("deck's first item should be ♣︎K, but get: %v", d[len(d)-1])
	}
}

func TestSave(t *testing.T) {
	os.Remove(testingFilename)

	save(toString(newDeck()), testingFilename)

	if _, err := os.Stat("./" + testingFilename); err != nil {
		t.Errorf(testingFilename+"should be exist, but got: %v", err)
	}

	os.Remove(testingFilename)
}

func TestRead(t *testing.T) {
	os.Remove(testingFilename)

	save(toString(newDeck()), testingFilename)

	d := read(testingFilename)

	if len(d) != 52 {
		t.Errorf("deck length should be 52, but got: %v", len(d))
	}

	os.Remove(testingFilename)
}
