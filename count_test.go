package gofalsesharing

import (
	"testing"
)

func TestCount(t *testing.T) {
	expected := 4
	actual := Count("uh2789y23fiubjwheb2783rgfwhefbwef", 'f')
	if actual != expected {
		t.Fatalf("actual=%d, expected=%d", actual, expected)
	}
}

func TestCountConcurrent(t *testing.T) {
	expected := 4
	actual := CountConcurrent("uh2789y23fiubjwheb2783rgfwhefbwef", 'f')
	if actual != expected {
		t.Fatalf("actual=%d, expected=%d", actual, expected)
	}
}
