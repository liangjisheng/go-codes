package strutils

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	t.Log(CompareVersion("1.4.0", "1.3.0"))
}

func TestDiff(t *testing.T) {
	testCase := [][]string{
		{"foo", "bar", "hello"},
		{"foo", "bar", "world"},
	}
	result := Diff(testCase[0], testCase[1])
	if len(result) != 1 || result[0] != "hello" {
		t.Fatalf("Diff failed")
	}
}
