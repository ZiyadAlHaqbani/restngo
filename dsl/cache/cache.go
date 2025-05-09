package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var cache = map[string][]byte{}
var metadata = map[string]CacheMetadata{}
var accessCount uint = 0

func init() {
	loadConfigs()
}

// we use absolute file path to remove the potential of redundant loading of files
// for example a user could load 'D:proj/json_load.json' and 'json_load.json', both containt the same
// file but the cache key will be different.
//
//	for both reading and writing to the cache.
func FetchFileBytes(file_path string) ([]byte, error) {
	defer sweep()

	abs_file_path, err := filepath.Abs(file_path)
	if err != nil {
		return []byte{}, err
	}

	if cached_content, exists := cache[abs_file_path]; exists {
		accessCount++
		meta := metadata[abs_file_path]
		meta.AccessTime = time.Now()
		meta.TimesUsed++
		metadata[abs_file_path] = meta
		return cached_content, nil
	}

	file_content, err := os.ReadFile(abs_file_path)
	if err != nil {
		return []byte{}, err
	}

	accessCount++
	cache[abs_file_path] = file_content
	metadata[abs_file_path] = CacheMetadata{
		AccessTime: time.Now(),
		FirstStore: time.Now(),
		LastMod:    time.Now(),
		TimesUsed:  1,
	}
	return file_content, nil
}

// this function will always read from file, even if it exists in cache.
func LoadFile(file_path string) ([]byte, error) {

	abs_file_path, err := filepath.Abs(file_path)
	if err != nil {
		return []byte{}, err
	}

	file_content, err := os.ReadFile(abs_file_path)
	if err != nil {
		return []byte{}, err
	}

	accessCount++
	cache[abs_file_path] = file_content
	metadata[abs_file_path] = CacheMetadata{
		AccessTime: time.Now(),
		FirstStore: time.Now(),
		LastMod:    time.Now(),
		TimesUsed:  1,
	}
	return file_content, nil

}

// for removing a file from cache, if it doesn't exist return error
func UnloadFile(file_path string) error {

	abs_file_path, err := filepath.Abs(file_path)
	if err != nil {
		return err
	}

	if _, exists := cache[abs_file_path]; exists {
		delete(cache, abs_file_path)
		delete(metadata, abs_file_path)
		return nil
	}

	return fmt.Errorf("file path: %q doesn't exist in cache", abs_file_path)
}

// doesn't load files if not found in cache
func GetFileBytesRaw(file_path string) ([]byte, error) {

	abs_file_path, err := filepath.Abs(file_path)
	if err != nil {
		return []byte{}, err
	}

	if _, exists := cache[abs_file_path]; !exists {
		return []byte{}, fmt.Errorf("file path: %q doesn't exist in cache", abs_file_path)
	}

	accessCount++
	meta := metadata[abs_file_path]
	meta.AccessTime = time.Now()
	meta.TimesUsed++
	metadata[abs_file_path] = meta
	return cache[abs_file_path], nil
}

func sweep() {

	if accessCount%_configs.SweepTriggerCount != 0 {
		return
	}

	for key, meta := range metadata {
		if meta.isStale() {
			delete(metadata, key)
			delete(cache, key)
		}
	}

}
