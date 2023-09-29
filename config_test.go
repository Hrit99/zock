package main

import (
	"testing"
)

func Test_Loadenv(t *testing.T) {
	result := Loadenv()

	if result != nil {
		t.Errorf("\"Loadenv()\" FAILED, expected -> nil, got -> %v", result)
	} else {
		t.Logf("\"Loadenv()\" FAILED, expected -> nil, got -> %v", result)
	}

}
