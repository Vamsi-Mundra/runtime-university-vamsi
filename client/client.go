package client

import (
	"context"

	"github.com/heroku/tbalthazar-runtime-university/spec"
	"google.golang.org/grpc"
)

type RouteGuide struct {
	client spec.RouteGuideClient
}

func New(conn *grpc.ClientConn) *RouteGuide {
	return &RouteGuide{client: spec.NewRouteGuideClient(conn)}
}

func (rg *RouteGuide) GetFeatures(ctx context.Context, points []spec.Point) ([]spec.Feature, error) {
	// TODO: if the caller cancels the ctx, the GetFeature client will prolly
	// return an error and this method will return the error too. Should this
	// method also test for cxt.Done() and return early? It doesn't do anything
	// time consuming outside of calling GetFeature on the client, so prolly not
	// worth it.
	features := make([]spec.Feature, len(points))
	for i := range points {
		f, err := rg.client.GetFeature(ctx, &points[i])
		if err != nil {
			return nil, err
		}
		// TODO: not sure yet how to properly silcen this warning:
		// assignment copies lock value to features[i]:
		// github.com/heroku/tbalthazar-runtime-university/spec.Feature contains
		// google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutex
		features[i] = *f
	}
	return features, nil
}
