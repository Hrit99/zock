package main

import (
	"testing"
)

func Test_Loadenv(t *testing.T) {
	result := Loadenv()

	if result != nil {
		t.Errorf("\"Loadenv()\" FAILED, expected -> <nil>, got -> %v", result)
	} else {
		t.Logf("\"Loadenv()\" PASSED, expected -> <nil>, got -> %v", result)
	}

}

func Benchmark_Loadenv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loadenv()
	}
}
