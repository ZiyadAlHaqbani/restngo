package cache

import (
	"slices"
	"testing"
	"time"
)

func TestFetchFileBytes(t *testing.T) {

	test_load_contents, err := FetchFileBytes("../../assets/test_load.json")
	if err != nil {
		t.Errorf("expected successful first read, got %+v", err)
	}

	test_load_cache, err := FetchFileBytes("../../assets/test_load.json")
	if err != nil {
		t.Errorf("expected successful cache read, got %+v", err)
	}

	if !slices.Equal(test_load_cache, test_load_contents) {
		t.Errorf("expected true, got false")
	}

}

func TestUnloadFile(t *testing.T) {

	TestFetchFileBytes(t)

	UnloadFile("../../assets/test_load.json")

	_, err := GetFileBytesRaw("../../assets/test_load.json")
	if err == nil {
		t.Error("expected err, got nil")
	}

}

func TestSweepTTL(t *testing.T) {

	_configs.SweepTriggerCount = 1
	_configs.TTL = Duration(time.Millisecond * 10)
	_configs.Strategy = "TTL"

	//load files for the first time and generate metadata
	FetchFileBytes("../../assets/test_load.json")
	FetchFileBytes("../../assets/configs/cache_config.json")

	time.Sleep(time.Millisecond * 10)

	//activate sweep
	FetchFileBytes("../../assets/test_load.json")

	if len(cache) > 0 {
		t.Errorf("expected empty cache, got cache of size: %+v", metadata)
	}
}
