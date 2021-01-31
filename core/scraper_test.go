package core

import (
	"errors"
	"testing"
)

func TestNewDefaultScraper(t *testing.T) {
	// test success
	_, err := NewDefaultScraper("", t.TempDir())
	if err != nil {
		t.Error("expected success, failed with", err)
	}

	// expect error on bad path
	_, err = NewDefaultScraper("", "asdf")
	if !errors.Is(err, ErrNotExist) {
		t.Error("expected ErrNotExist. Got: ", err)
	}
}
