package puller

import "testing"

func TestNew(t *testing.T) {
	puller := New()

	expected := &Puller{}

	if puller == nil {
		t.Fatalf("expected: %+v, got: nil", expected)
	}
}
