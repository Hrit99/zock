package database

import (
	"testing"

	"github.com/Hrit99/zock.git/config"
)

func Test_ConnectDb(t *testing.T) {
	_, result := ConnectDb(config.Uri)

	if result != nil {
		t.Errorf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", result)
	} else {
		t.Logf("\"ConnectDb()\" PASSED, expected -> <nil>, got -> %v", result)
	}

}

func Benchmark_ConnectDb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectDb(config.Uri)
	}
}
