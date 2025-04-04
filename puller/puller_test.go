package puller

import (
	"testing"

	"github.com/graphql-go/compatibility-base/types"
)

func TestNew(t *testing.T) {
	puller := New()

	expected := &Puller{}

	if puller == nil {
		t.Fatalf("expected: %+v, got: nil", expected)
	}
}

func TestPullerPull(t *testing.T) {
	puller := New()

	expected := &PullResult{}

	params := &PullParams{
		Implementation: &types.Repository{
			Name:          "graphql-graphql-js",
			URL:           "https://github.com/graphql/graphql-js",
			ReferenceName: "v0.6.0",
			Dir:           "./repos/graphql-graphql-js/",
		},
	}
	result, err := puller.Pull(params)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if result == nil {
		t.Fatalf("expected: %+v, got: nil", expected)
	}
}
