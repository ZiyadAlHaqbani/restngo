
#   go to the root project dir
cd ..

#   build to avoid undeterministic delay before program start
go build .

#   start program in background
./htestp &
#save pid to kill later
GO_PID=$!

#   wait until pprof server starts on original process
sleep 5

#   save 5 seconds of the profiling session
curl -o cpu.prof "http://localhost:6060/debug/pprof/profile?seconds=10"

#   start pprof web view for the downloaded profile
go tool pprof -http=:8080 cpu.prof &

#kill process and end profile session
kill $GO_PID
