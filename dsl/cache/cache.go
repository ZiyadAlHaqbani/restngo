package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

var cache = map[string][]byte{}

// we use absolute file path to remove the potential of redundant loading of files
// for example a user could load 'D:proj/json_load.json' and 'json_load.json', both containt the same
// file but the cache key will be different.
//
//	for both reading and writing to the cache.
func GetFileBytes(file_path string) ([]byte, error) {

	abs_file_path, err := filepath.Abs(file_path)
	if err != nil {
		return []byte{}, err
	}

	if cached_content, exists := cache[abs_file_path]; exists {
		return cached_content, nil
	}

	file_content, err := os.ReadFile(abs_file_path)
	if err != nil {
		return []byte{}, err
	}

	cache[abs_file_path] = file_content
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

	cache[abs_file_path] = file_content
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
		return nil
	}

	return fmt.Errorf("file path: %q doesn't exist in cache", abs_file_path)
}
