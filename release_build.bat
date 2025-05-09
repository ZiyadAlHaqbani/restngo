@echo off

::  this script is for building the release version of the dsl interpreter
::  it sets the cache config path to one adjacent to the binaries instead
::  of "dsl/cache/cache_config.json" found in the project files. 
go build -ldflags "-X 'dsl/cache/cache.configPath=cache_config.json'" -o app .