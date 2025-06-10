package cache

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"slices"
	"time"
)

var _Policies = []string{"TTL", "LRU"}

// workaround to support json unmarshalling a duration type
type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		print(err)
		return err
	}

	dur, err := time.ParseDuration(s)
	if err != nil {
		print(err)
		return err
	}

	*d = Duration(dur)
	return nil
}

// can be overridden by the build options
var configPath string = "../../assets/configs/cache_config.json"

var _configs struct {
	TTL               Duration
	MaxEntries        uint
	MaxSizeBytes      uint
	Strategy          string
	SweepTriggerCount uint
}

// load the configs from cache dir
func loadConfigs() {

	abs_file_path, err := filepath.Abs(configPath)
	if err != nil {
		log.Panicf("couldn't get the absolute file path of: %q, %+v", configPath, abs_file_path)
	}

	config_stream, err := os.Open(abs_file_path)
	if err != nil {
		log.Panicf("couldn't load cache configs: %+v", err)
	}

	decoder := json.NewDecoder(config_stream)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&_configs)
	if err != nil {
		log.Panicf("couldn't decode the configs: %+v", err)
	}

	if !slices.Contains(_Policies, _configs.Strategy) {
		log.Panicf("unrecognized Strategy: %q, instead chose: %v", _configs.Strategy, _Policies)
	}

}
