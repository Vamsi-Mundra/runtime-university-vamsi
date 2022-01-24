package client

import (
	"context"
	"errors"
	"testing"

	"github.com/heroku/tbalthazar-runtime-university/spec"
	"google.golang.org/grpc"
)

type fakeRouteGuideClient struct {
	feature *spec.Feature
	err     error
}

func (c *fakeRouteGuideClient) GetFeature(ctx context.Context, in *spec.Point, opts ...grpc.CallOption) (*spec.Feature, error) {
	if c.err != nil {
		return nil, c.err
	}
	return c.feature, nil
}

func (c *fakeRouteGuideClient) ListFeatures(ctx context.Context, in *spec.Rectangle, opts ...grpc.CallOption) (spec.RouteGuide_ListFeaturesClient, error) {
	return nil, nil
}
func (c *fakeRouteGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (spec.RouteGuide_RecordRouteClient, error) {
	return nil, nil
}
func (c *fakeRouteGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (spec.RouteGuide_RouteChatClient, error) {
	return nil, nil
}

func TestGetFeatures(t *testing.T) {
	tc := []struct {
		name       string
		points     []spec.Point
		nbFeatures int
		err        error
	}{
		{"single point", []spec.Point{{}}, 1, nil},
		{"multiple points", []spec.Point{{}, {}}, 2, nil},
		{"error", []spec.Point{{}}, 0, errors.New("something bad")},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			rg := RouteGuide{
				client: &fakeRouteGuideClient{
					feature: &spec.Feature{},
					err:     tt.err,
				},
			}
			features, err := rg.GetFeatures(context.Background(), tt.points)

			if got, want := len(features), tt.nbFeatures; got != want {
				t.Fatalf("nb features: got %v feature(s), wanted %v", got, want)
			}
			if err != tt.err {
				t.Fatalf("error: got %v, wanted %v", err, nil)
			}
		})
	}
}
