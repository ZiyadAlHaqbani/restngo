package cache

import (
	"slices"
	"testing"
	"time"
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

func TestSweepTTL(t *testing.T) {

	_configs.SweepTriggerCount = 1
	_configs.TTL = Duration(time.Microsecond * 10)
	_configs.Strategy = "TTL"

	LoadFile("test_load.json")
	LoadFile("cache_config.json")

	LoadFile("test_load.json")
	LoadFile("test_load.json")
	LoadFile("test_load.json")

	t.Errorf("%v\n%+v", cache, metadata)
}
