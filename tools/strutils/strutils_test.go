package strutils

import (
	"fmt"
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

func TestTwoSet(t *testing.T) {
	slice1 := []string{"1", "2", "3", "3", "6", "8", "8"}
	slice2 := []string{"2", "3", "3", "5", "0"}
	un := Union(slice1, slice2)
	fmt.Println("slice1与slice2的并集为：", un)
	in := Intersect(slice1, slice2)
	fmt.Println("slice1与slice2的交集为：", in)
	di := Difference(slice1, slice2)
	fmt.Println("slice1与slice2的差集为：", di)
}
