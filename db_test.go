package main

import (
	"testing"
)

func Test_ConnectDb(t *testing.T) {
	_, result := ConnectDb(uri)

	if result != nil {
		t.Errorf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", result)
	} else {
		t.Logf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", result)
	}

}
