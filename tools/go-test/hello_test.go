package gotest

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{"111", args{
			"hello",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HelloWorld(tt.args.str)
		})
	}
}
