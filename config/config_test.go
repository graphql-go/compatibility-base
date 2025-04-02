package config

import "testing"

func TestNewConfig(t *testing.T) {
	cfg := New()

	expected := &Config{}

	if cfg == nil {
		t.Fatalf("expected: %+v, got: nil", expected)
	}
}
