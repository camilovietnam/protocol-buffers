module main

go 1.17

replace go_gen/src => /Users/long/erase/protocol-buffers/6-golang-programming/src

require (
	github.com/golang/protobuf v1.5.2
	go_gen/src v0.0.0-00010101000000-000000000000
)

require google.golang.org/protobuf v1.27.1 // indirect
