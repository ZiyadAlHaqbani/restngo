package cache

import (
	"slices"
	"testing"
)

func TestFetchFileBytes(t *testing.T) {

	test_load_contents, err := FetchFileBytes("test_load.json")
	if err != nil {
		t.Errorf("expected successful first read, got %+v", err)
	}

	test_load_cache, err := FetchFileBytes("test_load.json")
	if err != nil {
		t.Errorf("expected successful cache read, got %+v", err)
	}

	if !slices.Equal(test_load_cache, test_load_contents) {
		t.Errorf("expected true, got false")
	}

}

func TestUnloadFile(t *testing.T) {

	TestFetchFileBytes(t)

	UnloadFile("test_load.json")

	_, err := GetFileBytesRaw("test_load.json")
	if err == nil {
		t.Error("expected err, got nil")
	}

}

func TestSweep(t *testing.T) {

}
