package main

import (
	"context"
	"flag"
	"log"

	"github.com/heroku/tbalthazar-runtime-university/client"
	"github.com/heroku/tbalthazar-runtime-university/spec"
	"github.com/joeshaw/envdecode"
	"google.golang.org/grpc"
)

type Config struct {
	ServerURL string `env:"SERVER_URL,required"`
}

func main() {
	var lat = flag.Int("lat", 407838351, "latitude of the point")
	var long = flag.Int("long", -746143763, "longitude of the point")
	flag.Parse()

	var cfg Config
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatal("Could not decode env variables: ", err)
	}

	points := []spec.Point{
		{Latitude: int32(*lat), Longitude: int32(*long)},
	}

	conn, err := grpc.Dial(cfg.ServerURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not create client connection: ", err)
	}
	defer conn.Close()
	log.Println("connection created")

	rg := client.New(conn)
	features, err := rg.GetFeatures(context.Background(), points)
	if err != nil {
		log.Fatal("couldn't not get features: ", err)
	}

	log.Println("Features:", features)
}
