package myregex

import "testing"

func TestIPv4(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{
			ip:   "",
			want: false,
		},
		{
			ip:   "0",
			want: false,
		},
		{
			ip:   "0.",
			want: false,
		},
		{
			ip:   ".",
			want: false,
		},
		{
			ip:   "0.0",
			want: false,
		},
		{
			ip:   "0.0.",
			want: false,
		},
		{
			ip:   "0.0.0",
			want: false,
		},
		{
			ip:   "0.0.0.",
			want: false,
		},
		{
			ip:   "0.0.0.0",
			want: true,
		},
		{
			ip:   "0.0.0.0.",
			want: false,
		},
		{
			ip:   "0.0.0.0.0",
			want: false,
		},
		{
			ip:   "1.47.245.128",
			want: true,
		},
		{
			ip:   "a.47.245.128",
			want: false,
		},
		{
			ip:   "123.47.245.1289",
			want: false,
		},
	}

	for _, tt := range tests {
		got := RegexIPv4.MatchString(tt.ip)
		if got != tt.want {
			t.Errorf("ip %s want %v got %v", tt.ip, tt.want, got)
		}
	}

	for _, tt := range tests {
		got := Regex1IPv4.MatchString(tt.ip)
		if got != tt.want {
			t.Errorf("ip %s want %v got %v", tt.ip, tt.want, got)
		}
	}
}
