// +heroku goVersion go1.14
// +heroku install ./...

module github.com/heroku/tbalthazar-runtime-university

go 1.14

require (
	github.com/golang/protobuf v1.4.1
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.22.0
)
